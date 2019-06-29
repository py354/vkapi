package vkapi

import "strconv"

type getUserNameData struct {
	FirstName string `json:"first_name"`
}

type getUserNameResponse struct {
	Response []getUserNameData `json:"response"`
}

func (c *client) GetUserName(userID int) string {
	jsonR := c.Request("users.get", "user_ids="+strconv.Itoa(userID))
	response := getUserNameResponse{}
	err := json.Unmarshal(jsonR, &response)
	CheckError(err)
	return response.Response[0].FirstName
}
