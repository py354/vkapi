package main

import (
	"fmt"
	"vkapi"
)

func main() {
	token := "<token>"
	groupID := 183962820
	c := vkapi.NewUserClient(token)

	mediaID, ownerID := c.UploadPhotoToWallFromPath("examples/photo.png", groupID)
	c.Post(-groupID, "test", fmt.Sprintf("photo%d_%d", ownerID, mediaID))
}
