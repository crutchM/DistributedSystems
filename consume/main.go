package main

import (
	"encoding/json"
	"log"
	"os"
)

type Link struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}

func main() {
	consumer := NewConsumer()
	client := NewClient()
	log.Printf("Consumer ready, PID: %d", os.Getpid())

	for msg := range consumer.MessageChannel {
		log.Printf("Making request by url: %s", msg.Body)
		url := &Link{}
		err := json.Unmarshal(msg.Body, url)
		if err != nil {
			log.Print("invalid message")
		}
		client.Ping(url.Url, url.Id)
	}

}
