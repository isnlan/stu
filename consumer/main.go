package main

import (
	"fmt"
	KafkaCluster "github.com/bsm/sarama-cluster"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	brokers := []string{"127.0.0.1:9092", "127.0.0.1:9192", "127.0.0.1:9292"}
	topics := []string{"my_topic", "other_topic"}
	consumerId := "1"

	config := KafkaCluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Group.Return.Notifications = true

	//第二个参数是groupId
	consumer, err := KafkaCluster.NewConsumer(brokers, "consumer-group2", topics, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// 接收错误
	go func() {
		for err := range consumer.Errors() {
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	// 打印一些rebalance的信息
	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("Rebalanced: %+v\n", ntf)
		}
	}()

	// 消费消息
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				fmt.Printf("consumer_id:%s, topoc:%s, partition:%d, offset:%d, key:%s, value:%s\n",
					consumerId, msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				consumer.MarkOffset(msg, "") // 提交offset
			}
		case <-signals:
			return
		}
	}
}
