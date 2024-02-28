package kzline

import (
	"encoding/json"
	"fmt"
	"github.com/kuroshibaz/lib/utils"
	"log"
)

type PushMessageRequest struct {
	Message string `json:"message" form:"message"`
	//ImageThumbnail       string `json:"imageThumbnail" form:"imageThumbnail"`
	//ImageFullSize        string `json:"imageFullsize" form:"imageFullsize"`
	//ImageFile            string `json:"imageFile" form:"imageFile"`
	//StickerPackageId     int  `json:"stickerPackageId" form:"stickerPackageId"`
	//StickerId            int  `json:"stickerId" form:"stickerId"`
	NotificationDisabled bool `json:"notificationDisabled" form:"notificationDisabled"`
}

type PushMessageResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (c *defaultClient) PushMessage(request PushMessageRequest) (*PushMessageResponse, error) {
	url := fmt.Sprintf("%s/api/notify", c.LineApi)

	formData := utils.FillMapStruct(request)

	log.Print(formData)

	response, err := c.restyClient.R().
		SetHeaders(map[string]string{
			"Authorization": c.getHeaderToken(),
			"Content-Type":  "application/x-www-form-urlencoded",
		}).
		SetFormData(formData).
		Post(url)
	if err != nil {
		return nil, err
	}

	result := PushMessageResponse{}

	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
