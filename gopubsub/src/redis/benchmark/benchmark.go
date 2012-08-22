package main

import (
	"fmt"
	"redis"
	"os"
	"strconv"
)

func main() {
	var client, pub redis.Client
	messages := make(chan redis.Message)
	sub := make(chan string, 1)
	uid := os.Args[1]

	fmt.Println("\n\nuser " + uid + " Publish to his channel:")
	pub.Publish("channel" + uid , []byte("I'm go!"))
	for chan_num := 0; chan_num <= 500; chan_num++{
		chan_number := strconv.Itoa(chan_num)
		sub <- "channel" + chan_number
		if chan_number != uid  {
			go client.Subscribe(sub, nil, nil, nil, messages)
		}
	}

	for {
		msg := <-messages
		fmt.Println("user ", uid, "received from:", msg.Channel, " message:", string(msg.Message))
	}
}
