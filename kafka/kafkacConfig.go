package kafka

import (
	"github.com/Shopify/sarama"
)

func GetProducer() (sarama.SyncProducer, error) {
	brokersUrl := []string{"192.168.130.109:30092"}
	produceConfig := sarama.NewConfig()
	produceConfig.Producer.Return.Successes = true
	produceConfig.Producer.RequiredAcks = sarama.WaitForAll
	produceConfig.Producer.Retry.Max = 5
	// NewSyncProducer creates a new SyncProducer using the given broker addresses and configuration.
	producer, err := sarama.NewSyncProducer(brokersUrl, produceConfig)
	if err != nil {
		return nil, err
	}
	return producer, err

}

func GetConsumer() (sarama.Consumer, error) {
	brokersUrl := []string{"192.168.130.109:30092"}
	consumerConfig := sarama.NewConfig()
	consumerConfig.Consumer.Return.Errors = true
	// NewConsumer creates a new consumer using the given broker addresses and configuration
	consumer, err := sarama.NewConsumer(brokersUrl, consumerConfig)
	if err != nil {
		return nil, err
	}
	return consumer, err
}
