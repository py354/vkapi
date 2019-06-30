package main

import "vkapi"

func main() {
	// our tokens
	tokens := []string{
		"6f8b5966d05b39b00a3ce5e45b249e7da495184e3b60cd59d9c536fe0a8aa8f01d5149b1e26d97dca1044",
		"90de7af44840900267f001ac0d94e116235f84359c18ad5f5a843519561437c8bac5e7e403539870b1bce",
		"cd9407e6d0945b647c874fe553bf32e0021cb6c2c5366ee28e963c0b334fac867113f875eb347ff08fa92",
	}

	// register clients
	clients := make([]*vkapi.Client, 0, 3)
	for _, token := range tokens {
		clients = append(clients, vkapi.NewClient(token))
	}

	// add client to pool
	for _, client := range clients {
		client.ActivatePool(&clients)
	}

	// some random client from our pool
	poolClient := clients[0]

	// done!
	// if one client can handle 20 messages per second,
	// our pool from 3 clients can handle 60 messages per second,
	// and if minimum delay between requests was 50 ms, now it ~16 ms

	// init longpoll and make messages handler
	inputMessages := make(chan *vkapi.Message, 100)
	lp := vkapi.NewLongpoll(tokens[0])
	go lp.Listen(inputMessages)

	for msg := range inputMessages {
		poolClient.SendMessage(msg.PeerID, "hello", "", "")
	}
}
