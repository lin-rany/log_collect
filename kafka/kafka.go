package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var (
	producer sarama.SyncProducer
)

func Init(addr []string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	producer, err = sarama.NewSyncProducer(addr, config)
	if err != nil {
		logs.Error("初识化Kafka producer失败:", err)
		return
	}
	logs.Debug("初始化Kafka producer成功,地址为:", addr)
	return
}
