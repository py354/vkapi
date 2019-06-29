package vkapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const apiVersion = "5.100"

const requestURL string = "https://api.vk.com/method/%s?access_token=%s&v=" + apiVersion

const sleepTime = time.Second / 20

func getRequestUrl(method, token string) string {
	return fmt.Sprintf(requestURL, method, token)
}

type client struct {
	token       string
	lastRequest time.Time
}

type serviceClient struct {
	client
}

func Client(token string) *client {
	return &client{token: token}
}

func ServiceClient(token string) *client {
	return &client{token: token}
}

func (c *client) Request(method, params string) []byte {
	left := time.Since(c.lastRequest)
	if left < sleepTime {
		time.Sleep(sleepTime - left)
	}

	rURL := getRequestUrl(method, c.token)

	reader := strings.NewReader(params)
	r, err := http.Post(rURL, "application/x-www-form-urlencoded", reader)
	defer r.Body.Close()
	CheckError(err)

	binAnswer, err := ioutil.ReadAll(r.Body)
	return binAnswer
}
