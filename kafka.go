package main

import (
	"log"
	"os"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/kelseyhightower/envconfig"
)

type KafkaConfig struct {
	Brokers []string `required:"true"`
	Topic   string   `required:"true"`

	TLS          bool
	SASLUser     string `envconfig:"SASL_USER"`
	SASLPassword string `envconfig:"SASL_PASSWORD"`

	ConsumerGroup string `envconfig:"CONSUMER_GROUP"`

	Debug bool `envconfig:"DEBUG"`
}

func ParseKafkaConfigFromEnv() (KafkaConfig, error) {
	var kafkaConfig KafkaConfig
	err := envconfig.Process("KAFKA", &kafkaConfig)
	if err != nil {
		return KafkaConfig{}, err
	}
	return kafkaConfig, err
}

func CreateKafkaClient(kafkaConfig KafkaConfig) (*cluster.Client, error) {
	config := cluster.NewConfig()
	config.Config.Net.TLS.Enable = kafkaConfig.TLS
	if kafkaConfig.SASLUser != "" {
		config.Config.Net.SASL.Enable = true
		config.Config.Net.SASL.User = kafkaConfig.SASLUser
		config.Config.Net.SASL.Password = kafkaConfig.SASLPassword
	}
	if kafkaConfig.Debug {
		sarama.Logger = log.New(os.Stdout, "[sarama-debug] ", log.LstdFlags)
	}
	config.Consumer.Return.Errors = true

	client, err := cluster.NewClient(kafkaConfig.Brokers, config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func CreateKafkaConsumer(kafkaconfig KafkaConfig) (*cluster.Consumer, error) {
	client, err := CreateKafkaClient(kafkaconfig)
	if err != nil {
		return nil, err
	}
	consumer, err := cluster.NewConsumerFromClient(client, kafkaconfig.ConsumerGroup, []string{kafkaconfig.Topic})
	if err != nil {
		return nil, err
	}
	return consumer, nil
}
