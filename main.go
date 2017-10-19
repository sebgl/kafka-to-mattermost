package main

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/Sirupsen/logrus"
)

func main() {
	mattermostPoster, err := ParseMattermostConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	kafkaConfig, err := ParseKafkaConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	consumer, err := CreateKafkaConsumer(kafkaConfig)
	if err != nil {
		log.Fatal(err)
	}

	// catch stop signals to properly close Kafka consumer
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	log.Info("Starting Kafka to Mattermost")

	for {
		select {

		case err := <-consumer.Errors():
			log.WithError(err).Error("Fail to consume Kafka message")

		case msg := <-consumer.Messages():
			err := mattermostPoster.PostMessage(msg.Value)
			if err != nil {
				log.WithError(err).Error("Fail to post message")
			}
			consumer.MarkOffset(msg, "")

		case <-shutdown:
			consumer.Close()
			return
		}
	}
}
