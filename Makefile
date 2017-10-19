dev:
	go run *.go

build-go:
	go build

build-docker:
	docker build -t sebgl/kafka-to-mattermost .

push-docker:
	docker push sebgl/kafka-to-mattermost

run:
	docker run --rm -ti --net=host \
		-e KAFKA_BROKERS=${KAFKA_BROKERS} \
		-e KAFKA_TLS=${KAFKA_TLS} \
		-e KAFKA_SASL_USER=${KAFKA_SASL_USER} \
		-e KAFKA_SASL_PASSWORD=${KAFKA_SASL_PASSWORD} \
		-e KAFKA_CONSUMER_GROUP=${KAFKA_CONSUMER_GROUP} \
		-e KAFKA_TOPIC=${KAFKA_TOPIC} \
		-e MATTERMOST_URL=${MATTERMOST_URL} \
		-e INCOMING_HOOK_KEY=${INCOMING_HOOK_KEY} \
		sebgl/kafka-to-mattermost