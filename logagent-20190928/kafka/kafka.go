package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

var (
	client sarama.SyncProducer
	MsgChan chan *sarama.ProducerMessage
)


// 初始化全局的Kafka Client
func Init(address []string, chanSize int64)(err error){
	//

	// 0. 生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll  // ACK
	config.Producer.Partitioner = sarama.NewRandomPartitioner  //分区
	config.Producer.Return.Successes = true  // 确认

	// 1. 连接kafka
	client, err = sarama.NewSyncProducer(address,config)
	if err!=nil{
		logrus.Error("kafka:producer closed , err", err)
		return
	}

	// 初始化msgChan
	MsgChan = make(chan *sarama.ProducerMessage, chanSize)

	logrus.Info("init kafka success")

	// 起一个后台的goroutine 从MsgChan中读取数据
	go sendMsg()

	return


	//defer Client.Close()


}


func sendMsg(){
	for {
		select {
		case msg := <- MsgChan:
			pid,offset,err := client.SendMessage(msg)
			if err!=nil{
				logrus.Warning(" send msg failed, err", err)
			}

			logrus.Infof("send msg to kafka success pid:%v, offset:%v", pid,offset)

		}
	}
}