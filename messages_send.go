package vkapi

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// messageTypes: sticker, text + kb + attachment
func (c *client) SendSticker(peerID int, stickerID int) {
	params := fmt.Sprintf("peer_id=%d&sticker_id=%d", peerID, stickerID)
	c.Request("messages.send", params)
}

func (c *client) sendMessage(dst string, message, keyboard, attachment string) []byte {
	params := dst + "&random_id=&"

	if message != "" {
		params += "message=" + message + "&"
	}

	if keyboard != "" {
		params += "keyboard=" + keyboard + "&"
	}

	if attachment != "" {
		params += "attachment=" + attachment + "&"
	}
	return c.Request("messages.send", params)
}

func (c *client) SendMessage(peerID int, message, keyboard, attachment string) []byte {
	return c.sendMessage("peer_id="+strconv.Itoa(peerID), message, keyboard, attachment)
}

func (c *client) Broadcast(userIDs, message, keyboard, attachment string) {
	c.sendMessage("userIDs="+userIDs, message, keyboard, attachment)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
