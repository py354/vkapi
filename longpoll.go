package vkapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Longpoll struct {
	*Client
	Key    string
	Server string
	TS     string
}

type Attachment struct {
	Type string `json:"type"`
}

type Message struct {
	FromID      int
	PeerID      int
	Text        string
	Payload     string
	ChatID      int
	Attachments []Attachment
}

type getLongPollServerData struct {
	Key    string `json:"key"`
	Server string `json:"server"`
	Ts     string `json:"ts"`
}

type LongpollEvent struct {
	FromID      int        `json:"from_id"`
	Out         int          `json:"out"`
	PeerID      int        `json:"peer_id"`
	Text        string       `json:"text"`
	Payload     string       `json:"Payload"`
	Attachments []Attachment `json:"attachments"`
}

type longpollUpdate struct {
	LongpollEvent `json:"object"`
	Type          string `json:"type"`
	GroupId       int  `json:"group_id"`
}

type longpollResponse struct {
	Ts      string           `json:"ts"`
	Updates []longpollUpdate `json:"updates"`
	Failed  int              `json:"failed"`
}

type getLongpollServerResponse struct {
	Response getLongPollServerData `json:"response"`
}

func NewLongpoll(token string) *Longpoll {
	return &Longpoll{Client: NewClient(token)}
}

func (lp *Longpoll) initVKParams() {
	jsonR := lp.Request("groups.getLongPollServer", "group_id="+strconv.Itoa(lp.GetGroupID()))
	response := getLongpollServerResponse{}
	err := json.Unmarshal(jsonR, &response)
	CheckError(err)

	lp.Key = response.Response.Key
	lp.Server = response.Response.Server
	lp.TS = response.Response.Ts
}

func (lp *Longpoll) getEvents() longpollResponse {
	url := fmt.Sprintf("%s?act=a_check&key=%s&ts=%s&wait=25", lp.Server, lp.Key, lp.TS)
	r, err := http.Get(url)
	CheckError(err)
	defer r.Body.Close()

	answer, err := ioutil.ReadAll(r.Body)
	CheckError(err)

	response := longpollResponse{}
	err = json.Unmarshal(answer, &response)
	CheckError(err)
	return response
}

// support only "message_new" event from users
func (lp *Longpoll) Listen(inputMessages chan<- *Message) {
	lp.initVKParams()

	for {
		response := lp.getEvents()
		if response.Failed != 0 || response.Ts == "" {
			lp.initVKParams()
			continue
		}

		lp.TS = response.Ts
		for _, event := range response.Updates {
			if event.Type != "message_new" || event.Out == 1 || event.FromID < 0 {
				continue
			}

			message := Message{
				FromID:      event.FromID,
				PeerID:      event.PeerID,
				Text:        event.Text,
				Payload:     event.Payload,
				Attachments: event.Attachments,
			}

			if event.FromID == event.PeerID {
				message.ChatID = 0
			} else {
				message.ChatID = event.PeerID - 2000000000
			}

			inputMessages <- &message
		}
	}
}
