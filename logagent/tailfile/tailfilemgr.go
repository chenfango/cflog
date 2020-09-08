package tailfile

import (
	"github.com/sirupsen/logrus"
	"logagent/common"
)

// 管理tailTask

type tailTaskMgr struct {
	tailTaskMap      map[string]*tailTask       // 所有的任务
	collectEntryList []common.CollectEntry      // 所有的配置项
	confChan         chan []common.CollectEntry // 等待新配置的通道

}

var (
	ttMgr *tailTaskMgr
)

// main函数调用
func Init(allConf []common.CollectEntry) (err error) {

	// allConf 里面存了若干个日志的收集项
	// 针对每一个日志收集项, 创建一个对应的tailObj对象

	ttMgr = &tailTaskMgr{
		tailTaskMap:      make(map[string]*tailTask, 20),
		collectEntryList: allConf,
		confChan:         make(chan []common.CollectEntry, 1), // 做一个阻塞channel
	}

	for _, conf := range allConf {
		tailtask := newTailTask(conf.Path, conf.Topic) // 创建一个日志收集的任务
		err = tailtask.Init()

		if err != nil {
			logrus.Error("create tailObj for path:%s , err:%v", conf.Path, err)
			continue
		}
		logrus.Info("create a tail task for path: %s success", conf.Path)
		//
		ttMgr.tailTaskMap[tailtask.path] = tailtask // 把创建的任务保存起来，方便后续管理
		// 去收集日志,后台goroutine

		go tailtask.run()

	}

	// 派一个小弟等着新配置来
	go ttMgr.watch() // 在后台等新的配置

	return
}

func (t *tailTaskMgr) watch() {
	for {

		// 派一个小弟等新的配置
		newConf := <-t.confChan // 取到值说明新的配置来了
		logrus.Infof("new file is going!")

		// 新的配置来后, 管理一下我们之前启动的那些tailTask

		logrus.Infof("get new conf from etcd, conf:%v, start manager tailTask...", newConf)

		for _, conf := range newConf {
			// 1. 原来已经存在的任务不用动
			if t.isExist(conf) {
				continue
			}

			// 2. 原来没有的新创建一个tailTask任务
			tailtask := newTailTask(conf.Path, conf.Topic)
			err := tailtask.Init() // 去打开日志文件准备读
			if err != nil {
				logrus.Errorf("create a tail task for path:%s failed", conf.Path)
			}
			logrus.Infof("create a tail task for path:%s success", conf.Path)
			t.tailTaskMap[tailtask.path] = tailtask

			// 起一个后台去收集日志
			go tailtask.run()

		}

		// 3. 原来有的,现在没有的需要停掉
		// 找出tailTaskMap中存在, 但newConf不存在的那些tailTask, 把他们都停掉
		for key, task := range t.tailTaskMap {
			var found bool
			for _, conf := range newConf {
				if key == conf.Path {
					found = true
					break
				}

			}

			if !found {
				// 这个task 要停掉了
				logrus.Infof("the task collect path:%s need to stop", task.path)
				delete(t.tailTaskMap, key) // 从管理类删掉
				task.cancel()

			}
		}
	}

}

func (t *tailTaskMgr) isExist(conf common.CollectEntry) bool {
	_, ok := t.tailTaskMap[conf.Path]
	return ok
}

func SendNewConf(newConf []common.CollectEntry) {
	ttMgr.confChan <- newConf
	logrus.Infof("sendNewConf success", ttMgr.confChan)
}
