package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Client struct{
	client *http.Client
}

func NewClient() *Client{
	return &Client{
		&http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (s *Client) Ping(url string, id int){
	var status string
	if !strings.Contains(url, "https://"){
		url = "https://" + url
	}

	resp, err := s.client.Get(url)
	resp.Body.Close()
	if err != nil{
		url = "http://" + url
		resp, err := s.client.Get(url)
		resp.Body.Close()
		if err!= nil{
			status = "400"
			s.SendStatus(id, status)
			return
		}
	}
	s.SendStatus(id, fmt.Sprint(resp.StatusCode))
}


func (s *Client) SendStatus(id int, status string){

}