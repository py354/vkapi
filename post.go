package vkapi

import (
	"strconv"
)

type PostData struct {
	PostID int	`json:"post_id"`
}

type PostDataResponse struct {
	Response PostData `json:"response"`
}

func (c *UserClient) Post(ownerID int, message, attachments string) int {
	params := "owner_id=" + strconv.Itoa(ownerID) + "&"

	if message != "" {
		params += "message=" + message + "&"
	}

	if attachments != "" {
		params += "attachments=" + attachments + "&"
	}

	response := c.Request("wall.post", params)
	data := PostDataResponse{}
	err := json.Unmarshal(response, &data)
	CheckError(err)
	return data.Response.PostID
}
