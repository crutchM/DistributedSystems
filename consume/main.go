package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/crutchm/consumer/internal"
)

type Link struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}

func main() {
	consumer := internal.NewConsumer()
	client := internal.NewClient()
	log.Printf("Consumer ready, PID: %d", os.Getpid())
	stop := make(chan struct{})
	go func() {
		for msg := range consumer.MessageChannel {
			log.Printf("Making request by url: %s", msg.Body)
			url := &Link{}
			err := json.Unmarshal(msg.Body, url)
			if err != nil {
				log.Print("invalid message")
			}
			client.Ping(url.Url, url.Id)
		}
	}()
	<-stop

}
