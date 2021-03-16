FROM golang

MAINTAINER Valdimir Mevzos

# install dependencies
RUN go get -u github.com/riferrei/srclient
RUN go get -u github.com/google/uuid
RUN go get -u gopkg.in/confluentinc/confluent-kafka-go.v1/kafka

# env
ENV KAFKA_HOST kafka.kafka-ca1
ENV KAFKA_CONSUMER_GROUP serde
ENV KAFKA_PORT 9092
ENV KAFKA_TOPIC tracking
ENV SCHEMA_REGISTRY_URL http://schema-registry-cp-schema-registry.kafka:8081

# copy app
ADD . /app
WORKDIR /app

# build
RUN go build -o build/consumer src/*.go

ENTRYPOINT ["/app/docker-entrypoint.sh"]
