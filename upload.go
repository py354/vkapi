package vkapi

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

type uploadServer struct {
	Url     string `json:"upload_url"`
	AlbumID int    `json:"album_id"`
	GroupID int    `json:"group_id"`
}

type uploadPhotoWallResponse struct {
	Server int `json:"server"`
	Photo  string `json:"photo"`
	Hash   string `json:"hash"`
}

type uploadServerResponse struct {
	Response uploadServer `json:"response"`
}

type savedPhoto struct {
	ID int	`json:"id"`
	OwnerID int 	`json:"owner_id"`
}

type savedPhotoResponse struct {
	Response []savedPhoto `json:"response"`
}

func (c *client) getMessagesUploadServer(peerID int) (url string, albumID, groupID int) {
	r := c.Request("photos.getMessagesUploadServer", "peer_id="+strconv.Itoa(peerID))
	response := uploadServerResponse{}
	err := json.Unmarshal(r, &response)
	CheckError(err)
	return response.Response.Url, response.Response.AlbumID, response.Response.GroupID
}

func (c *client) saveMessagesPhoto(photo, hash string, server int) (int, int) {
	params := fmt.Sprintf("photo=%s&server=%d&hash=%s", photo, server, hash)
	r := c.Request("photos.saveMessagesPhoto", params)
	resp := savedPhotoResponse{}
	err := json.Unmarshal(r, &resp)
	CheckError(err)
	return resp.Response[0].ID, resp.Response[0].OwnerID
}

func (c *client) UploadPhoto(reader io.Reader, peerID int) string {
	url, _, _ := c.getMessagesUploadServer(peerID)

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("photo", "test.png")
	CheckError(err)

	_, err = io.Copy(fileWriter, reader)
	CheckError(err)

	contentType := bodyWriter.FormDataContentType()
	err = bodyWriter.Close()
	CheckError(err)

	resp, err := http.Post(url, contentType, bodyBuf)
	CheckError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	CheckError(err)

	var uploaded uploadPhotoWallResponse
	err = json.Unmarshal(body, &uploaded)
	CheckError(err)

	mediaID, ownerID := c.saveMessagesPhoto(uploaded.Photo, uploaded.Hash, uploaded.Server)
	return fmt.Sprintf("photo%d_%d", ownerID, mediaID)
}

func (c *client) UploadPhotoFromPath(path string) string {
	file, err := os.Open("/home/danis/projects/go/src/vkapi/examples/зож.jpg")
	defer file.Close()
	CheckError(err)

	return c.UploadPhoto(file, 222691811)
}
