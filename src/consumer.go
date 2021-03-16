package main

import (
  "os"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"github.com/mevzosvlad/event-schemas"
)


var kafka_host =  os.Getenv("KAFKA_HOST")
var consumer_group = os.Getenv("KAFKA_CONSUMER_GROUP")
var kafka_topic =  os.Getenv("KAFKA_TOPIC")
var schema_registry_url = os.Getenv("SCHEMA_REGISTRY_URL")


func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafka_host,
		"group.id":          consumer_group,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	c.SubscribeTopics([]string{kafka_topic}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if msg.Headers != nil {
          event_type := ""
					for _ , header := range msg.Headers {
            if header.Key == "EventType" {
              event_type = string(header.Value)
						}
					}
					fmt.Printf("%% event_type: %v\n", event_type)
				}
		if err == nil {
			// &addressbook.Addressbook{} is pointer to a struct schema here : "github.com/mevzosvlad/event-schemas"
      payload := &addressbook.Addressbook{}
			// i would like to do something like this:
			//payload := &addressbook.event_type{}
			// means i will use diff schemas depend on event_type string
			err = proto.Unmarshal(msg.Value[5:],payload) // need to change if no magic byte applied
			payload_json, _ := json.Marshal(payload)
			fmt.Printf("Here is the message %s\n", string(payload_json))
		} else {
			fmt.Printf("Error consuming the message: %v (%v)\n", err, msg)
		}
	}

	c.Close()

}
