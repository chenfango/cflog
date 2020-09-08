package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
	"logagent/common"
	"logagent/tailfile"
	"time"
)

var (
	client *clientv3.Client
)

// 1. 初始化etcd连接
func Init(address []string) (err error) {
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   address,
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v", err)
		return
	}

	return
}

// 去etcd中根据给定的key获取配置文件

// 配置文件是一个切片, 切片中每一项是日志文件的路径和topic
func GetConf(key string) (collectEntryList []common.CollectEntry, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	gResp, err := client.Get(ctx, key)
	if err != nil {
		logrus.Error("get conf from etcd by key:%s failed, err:%v", key, err)
		return
	}

	if len(gResp.Kvs) == 0 {
		logrus.Warnf("get len:0 conf from etcd by key:%s failed", key)
		return
	}

	ret := gResp.Kvs[0]
	err = json.Unmarshal(ret.Value, &collectEntryList)

	fmt.Println(ret.Value)

	if err != nil {
		logrus.Error("json unmarshal fail, err:%v", err)
		return
	}

	return

}

func WatchConf(key string) {
	for {
		watchCh := client.Watch(context.Background(), key)
		for wresp := range watchCh {
			logrus.Info("get new conf from etcd!")
			for _, evt := range wresp.Events {
				fmt.Printf("type:%s key:%s value:%s\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
				var newConf []common.CollectEntry
				if evt.Type == clientv3.EventTypeDelete {
					// 如果是删除
					logrus.Warning("FBI warning:etcd delete the key!!!")
					tailfile.SendNewConf(newConf) // 没有任何接收就是阻塞的
					continue
				}
				err := json.Unmarshal(evt.Kv.Value, &newConf)
				fmt.Println("new confing====\n", newConf)
				tailfile.SendNewConf(newConf) // 没有任何接收就是阻塞的
				if err != nil {
					logrus.Errorf("json unmarshal new conf failed, err:%v", err)
					continue
				}

			}
		}

	}
}
