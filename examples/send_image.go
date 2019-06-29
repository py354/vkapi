package main

import (
	"easyimage/core"
	"log"
	"os"
	"vkapi"
)

func main() {
	// token from your group (need permission to send messages)
	token := "cd9407e6d0945b647c874fe553bf32e0021cb6c2c5366ee28e963c0b334fac867113f875eb347ff08fa92"

	// your user_id
	peerID := 222691811

	// client initialization
	c := vkapi.Client(token)

	// send message with text "hello"
	response := c.SendMessage(peerID, "hello", "", "")
	log.Println(string(response))

	// open image
	file, err := os.Open("path_to_image.png")
	core.CheckError(err)

	// download it to vk server
	attachment := c.UploadPhoto(file, peerID)

	// send this image with message
	response = c.SendMessage(peerID, "", "", attachment)
	log.Println(string(response))
}