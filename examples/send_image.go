package main

import (
	"fmt"
	"log"
	"os"
	"vkapi"
)

func main() {
	// token from your group (need permission to send messages)
	token := "eb9dc863b36e36683cd2a18f976705171011e0b9c9952f33121997940384e7f3047d63b877f881f132289"

	// your user_id
	peerID := 222691811

	// client initialization
	c := vkapi.NewClient(token)

	// send message with text "hello"
	response := c.SendMessage(peerID, "hello", "", "")
	log.Println(string(response))

	// open image
	file, err := os.Open("examples/photo.png")
	vkapi.CheckError(err)

	// download it to vk server
	ownerID, mediaID := c.UploadPhotoToMessages(file, peerID)

	// send this image with message
	response = c.SendMessage(peerID, "", "", fmt.Sprintf("photo%d_%d", ownerID, mediaID))
	log.Println(string(response))
}