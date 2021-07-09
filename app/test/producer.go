package test

import (
	"fmt"
	"github.com/Shopify/sarama"
	"stduentProject/kafka"
)

func sendToKafka(topic string,message []byte) error {
	producer,err := kafka.GetProducer()
	defer producer.Close()
	if err!=nil {
		panic("创建kafka生产者失败, error=" + err.Error())
	}
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	return nil
}