package vkapi

import url2 "net/url"

type getShortLinkData struct {
	ShortURL string `json:"short_url"`
}

type getShortLinkResponse struct {
	Response getShortLinkData `json:"response"`
}

func (c *Client) GetShortLink(url string) string {
	jsonR := c.Request("utils.getShortLink", "url="+url2.QueryEscape(url))
	response := getShortLinkResponse{}
	err := json.Unmarshal(jsonR, &response)
	CheckError(err)
	return response.Response.ShortURL
}
