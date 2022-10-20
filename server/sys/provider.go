package server

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type Provider struct {
	AmqpChannel *amqp.Channel
	Queue       string
}

type SendStruct struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}

func NewProvider() *Provider {
	provider := &Provider{}
	provider.getCon()
	return provider
}

func (s *Provider) Publish(id int, url string) {
	link := SendStruct{
		Id:  id,
		Url: url,
	}
	body, err := json.Marshal(link)
	if err != nil {
		log.Fatal("error encoding JSON")
	}

	err = s.AmqpChannel.Publish("", s.Queue, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	if err != nil {
		log.Fatal("error publishing message reason: ", err)
	}

	log.Printf("url: %s was sended", link.Url)
}

func (s *Provider) getCon() {
	conn, err := amqp.Dial("amqp://myuser:22334455@rabbit:5672/")
	if err != nil {
		log.Fatal("Can't connect to AMQP reason:", err)
	}

	amqpChannel, err := conn.Channel()
	if err != nil {
		log.Fatal("Can't create a amqpChannel")
	}

	s.AmqpChannel = amqpChannel
	queue, err := s.AmqpChannel.QueueDeclare("links", true, false, false, false, nil)
	if err != nil {
		log.Fatal("Could not declare `links` queue")
	}
	s.Queue = queue.Name

}
