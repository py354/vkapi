package vkapi

import "fmt"

type RepostProfile struct {
	ID int `json:"id"`
}

type ItemsResp struct {
	Profiles []RepostProfile `json:"profiles"`
}

type GetRepostsResponse struct {
	Response ItemsResp `json:"response"`
}

func (c *serviceClient) GetRepostsUsers(groupID, postID int) []int {
	userIDs := make([]int, 0, 1000)

	temp := "owner_id=-%d&post_id=%d&count=1000&offset=%d"
	offset := 0
	for {
		jsonR := c.Request("wall.getReposts", fmt.Sprintf(temp, groupID, postID, offset))
		response := GetRepostsResponse{}
		err := json.Unmarshal(jsonR, &response)
		CheckError(err)

		if len(response.Response.Profiles) == 0 {
			break
		}

		for _, profile := range response.Response.Profiles {
			userIDs = append(userIDs, profile.ID)
		}
		offset += 1000
	}

	return userIDs
}
