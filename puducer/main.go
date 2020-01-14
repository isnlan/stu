package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	brokers := []string{"192.168.99.100:30201", "192.168.99.100:30221", "192.168.99.100:30231", "192.168.99.100:30211"}
	producer, err := newProducer(brokers)
	if err != nil {
		panic(err)
	}

	myTopic := func(bytes []byte) {
		partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{
			Topic: "my_topic",
			Key:   sarama.ByteEncoder([]byte("my_topic")),
			Value: sarama.ByteEncoder(bytes),
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("myTopic partition：%d, offset:%d\n", partition, offset)
	}

	otherTopic := func(bytes []byte) {
		partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{
			Topic: "other_topic",
			Key:   sarama.ByteEncoder([]byte("other_topic")),
			Value: sarama.ByteEncoder(bytes),
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("otherTopic partition：%d, offset:%d\n", partition, offset)
	}

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			myTopic([]byte(fmt.Sprintf("my topic:%d", i)))
		} else {
			otherTopic([]byte(fmt.Sprintf("other topic:%d", i)))
		}
	}

}
func newProducer(brokers []string) (sarama.SyncProducer, error) {
	// For the data collector, we are looking for strong consistency semantics.
	// Because we don't change the flush settings, sarama will try to produce messages
	// as fast as possible to keep latency low.
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
	config.Producer.Return.Successes = true

	return sarama.NewSyncProducer(brokers, config)
}
