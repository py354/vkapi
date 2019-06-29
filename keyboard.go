package vkapi

import "log"

const (
	KbRed   = "negative"
	KbGreen = "positive"
	KbWhite = "default"
	KbBlue  = "primary"
)

type Payload struct {
	Command string `json:"button"`
}

type Action struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
	Label   string `json:"label"`
}

type Button struct {
	Action `json:"action"`
	Color  string `json:"color"`
}

func (bi *Action) SetPayload(pl string) {
	j, err := json.Marshal(Payload{pl})
	if err != nil {
		panic(err)
	}

	bi.Payload = string(j)
}

func (msg *Message) GetPayload() string {
	if msg.Payload == "" {
		return ""
	}

	data := Payload{}
	err := json.Unmarshal([]byte(msg.Payload), &data)
	if err != nil {
		return ""
	}
	return data.Command
}

func MakeButton(name, payload, color string) Button {
	action := Action{
		Type:  "text",
		Label: name,
	}
	action.SetPayload(payload)

	return Button{
		Action: action,
		Color:  color,
	}
}

type keyboard struct {
	OneTime     bool       `json:"one_time"`
	ButtonsGrid [][]Button `json:"buttons"`
	Cache       string     `json:"-"`
}

func Keyboard(oneTime bool, buttonsGrid [][]Button) *keyboard {
	return &keyboard{OneTime: oneTime, ButtonsGrid: buttonsGrid}
}

func (kb *keyboard) String() string {
	if kb.Cache == "" {
		data, err := json.Marshal(kb)
		if err != nil {
			log.Fatalln(data)
		}
		kb.Cache = string(data)
	}
	return kb.Cache
}
