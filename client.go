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

type Request struct {
	MethodName string
	params string
}

func getRequestUrl(method, token string) string {
	return fmt.Sprintf(requestURL, method, token)
}

type Client struct {
	token       string
	lastRequest time.Time

	workWithPool bool
	requestsChan chan Request
}

type ServiceClient struct {
	Client
}

func NewClient(token string) *Client {
	return &Client{token: token}
}

func NewServiceClient(token string) *ServiceClient {
	return &ServiceClient{*NewClient(token)}
}

func (c *Client) ActivatePool(requestsChan chan Request) {
	c.requestsChan = requestsChan
	c.workWithPool = true
}

func (c *Client) DisablePool() {
	c.workWithPool = false
}

func (c *Client) ListenQuery() {
	for req := range c.requestsChan {
		c.Request(req.MethodName, req.params)
	}
}

func (c *Client) request(method, params string) []byte {
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

func (c *Client) Request(method, params string) []byte {
	if !c.workWithPool {
		return c.request(method, params)
	} else {
		c.requestsChan <- Request{
			MethodName: method,
			params:     params,
		}
		return []byte{}
	}
}
