package main

import (
	"github.com/streadway/amqp"
	"log"
)

type Consumer struct{
	Channel *amqp.Channel
	MessageChannel <- chan amqp.Delivery
}

func NewConsumer() *Consumer  {
	consumer := &Consumer{}
	consumer.connect()
	consumer.subscribe()
	return consumer
}

func (s *Consumer) connect(){
	con, err := amqp.Dial("amqp://user:user@rabbit:5672")
	if err != nil{
		log.Fatal("cannot connect to rabbit")
		con.Close()
	}
	s.Channel, err = con.Channel()
	if err != nil{
		log.Fatal("cant create channel")
		s.Channel.Close()
	}
}

func (s *Consumer) subscribe() {
	queue, err := s.Channel.QueueDeclare("links", true, false, false,false, nil)
	if err != nil{
		log.Fatal("couldn't declare queue")
	}

	err = s.Channel.Qos(1, 0, false)
	if err != nil{
		log.Fatal("could'n configure QoS")
	}
	messageChannel, err := s.Channel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
		)
	if err != nil{
		log.Fatal("could'n register channel")
	}
	s.MessageChannel = messageChannel

}