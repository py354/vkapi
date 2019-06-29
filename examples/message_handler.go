package main

import (
	"log"
	"vkapi"
)

func main() {
	// token from your group (need permission to send messages)
	token := "cd9407e6d0945b647c874fe553bf32e0021cb6c2c5366ee28e963c0b334fac867113f875eb347ff08fa92"

	// client initialization
	c := vkapi.Client(token)

	// make channel for input messages
	inputMessages := make(chan *vkapi.Message)

	// longpoll initialization
	lp := vkapi.Longpoll(token)

	// start listening
	go lp.Listen(inputMessages)

	// create keyboard
	kb := vkapi.Keyboard(false, [][]vkapi.Button{
		{
			vkapi.MakeButton("Button name", "payload1", vkapi.KbGreen),
		},
	})

	// handle input messages from channel
	for msg := range inputMessages {

		// check if button was pressed
		if msg.GetPayload() == "payload1" {
			r := c.SendMessage(msg.PeerID, "button was pressed", kb.String(), "")
			log.Println(string(r))

		} else {
			// send message with keyboard
			r := c.SendMessage(msg.PeerID, "hello", kb.String(), "")
			log.Println(string(r))
		}
	}
}
