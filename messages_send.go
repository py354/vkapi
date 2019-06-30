package vkapi

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// messageTypes: sticker, text + kb + attachment
func (c *Client) SendSticker(peerID int, stickerID int) {
	params := fmt.Sprintf("peer_id=%d&sticker_id=%d", peerID, stickerID)
	c.Request("messages.send", params)
}

func (c *Client) sendMessage(dst string, message, keyboard, attachment string) []byte {
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

func (c *Client) SendMessage(peerID int, message, keyboard, attachment string) []byte {
	return c.sendMessage("peer_id="+strconv.Itoa(peerID), message, keyboard, attachment)
}

func (c *Client) Broadcast(userIDs, message, keyboard, attachment string) {
	c.sendMessage("userIDs="+userIDs, message, keyboard, attachment)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
