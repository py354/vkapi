package vkapi

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

type WallUploadServer struct {
	UploadURL string `json:"upload_url"`
	AlbumID	int 	`json:"album_id"`
	UserID int		`json:"user_id"`
}

type WallUploadServerResponse struct {
	Response WallUploadServer `json:"response"`
}

func (c *UserClient) getWallUploadServer(groupID int) (uploadURL string, albumID, userID int) {
	r := c.Request("photos.getWallUploadServer", "group_id=" + strconv.Itoa(groupID))
	data := WallUploadServerResponse{}
	err := json.Unmarshal(r, &data)
	CheckError(err)
	return data.Response.UploadURL, data.Response.AlbumID, data.Response.UserID
}

func (c *UserClient) saveWallPhoto(groupID int, photo string, server int, hash string) (mediaID, ownerID int) {
	params := fmt.Sprintf("group_id=%d&photo=%s&server=%d&hash=%s", groupID, photo, server, hash)
	r := c.Request("photos.saveWallPhoto", params)
	resp := savedPhotoResponse{}
	err := json.Unmarshal(r, &resp)
	CheckError(err)
	return resp.Response[0].ID, resp.Response[0].OwnerID
}

func (c *UserClient) UploadPhotoToWall(reader io.Reader, groupID int) (mediaID, ownerID int) {
	url, _, _ := c.getWallUploadServer(groupID)
	server, photo, hash := UploadPhotoToServer(reader, url)
	return c.saveWallPhoto(groupID, photo, server, hash)
}

func (c *UserClient) UploadPhotoToWallFromPath(path string, groupID int) (mediaID, ownerID int) {
	file, err := os.Open(path)
	defer file.Close()
	CheckError(err)

	return c.UploadPhotoToWall(file, groupID)
}
