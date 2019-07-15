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
	workers *[]*Client
	workerIndex int
}

type ClientsPool struct {
	*Client
	clients []*Client
}

type ServiceClient struct {
	Client
}

type UserClient struct {
	Client
}

func NewPool(tokens []string) ClientsPool {
	if len(tokens) == 0 {
		panic("need at least one token")
	}

	pool := ClientsPool{
		Client:  nil,
		clients: make([]*Client, 0, 10),
	}

	for _, token := range tokens {
		c := NewClient(token)
		c.ActivatePool(&pool.clients)
		pool.clients = append(pool.clients, c)
	}

	pool.Client = pool.clients[0]
	return pool
}

func NewClient(token string) *Client {
	return &Client{token: token}
}

func NewUserClient(token string) *UserClient {
	return &UserClient{*NewClient(token)}
}

func NewServiceClient(token string) *ServiceClient {
	return &ServiceClient{*NewClient(token)}
}

func (c *Client) ActivatePool(workers *[]*Client) {
	c.workWithPool = true
	c.workers = workers
}

func (c *Client) DisablePool() {
	c.workWithPool = false
}

func (c *Client) request(method, params string) []byte {
	left := time.Since(c.lastRequest)
	if left < sleepTime {
		time.Sleep(sleepTime - left)
	}
	c.lastRequest = time.Now()

	rURL := getRequestUrl(method, c.token)
	fmt.Println(rURL)
	fmt.Printf("%#v\n\n", params)

	reader := strings.NewReader(params)
	r, err := http.Post(rURL, "application/x-www-form-urlencoded", reader)
	defer r.Body.Close()
	CheckError(err)

	binAnswer, err := ioutil.ReadAll(r.Body)
	if strings.Contains(string(binAnswer), "err") {
		fmt.Println(string(binAnswer))
	}
	return binAnswer
}

func (c *Client) Request(method, params string) []byte {
	if !c.workWithPool {
		return c.request(method, params)
	} else {
		c.workerIndex += 1
		if c.workerIndex == len(*c.workers) {
			c.workerIndex = 0
		}
		return (*c.workers)[c.workerIndex].request(method, params)
	}
}
