# Kafka to Mattermost incoming webhook

Consume messages from a Kafka topic and post them to a Mattermost channel, by making use of [Mattermost incoming webhook](https://docs.mattermost.com/developer/webhooks-incoming.html).

## Usage

```
docker run -d \
    -e KAFKA_BROKERS=${KAFKA_BROKERS} \
    -e KAFKA_TLS=${KAFKA_TLS} \
    -e KAFKA_SASL_USER=${KAFKA_SASL_USER} \
    -e KAFKA_SASL_PASSWORD=${KAFKA_SASL_PASSWORD} \
    -e KAFKA_CONSUMER_GROUP=${KAFKA_CONSUMER_GROUP} \
    -e KAFKA_TOPIC=${KAFKA_TOPIC} \
    -e MATTERMOST_URL=${MATTERMOST_URL} \
    -e INCOMING_HOOK_KEY=${INCOMING_HOOK_KEY} \
    sebgl/kafka-to-mattermost
```

## Kafka message format

See [Mattermost incoming webhook documentation](https://docs.mattermost.com/developer/webhooks-incoming.html).
Example:
```
{
    "text": "my message :tada:",
    "channel": "optionnal",
    "username": "optionnal"
}
```