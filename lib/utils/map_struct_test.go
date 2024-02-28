package utils

import "testing"

type MapTest struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Single bool   `json:"single"`
}

type MapPushMessageRequest struct {
	Message              string `json:"message" form:"message"`
	ImageThumbnail       string `json:"imageThumbnail" form:"imageThumbnail"`
	ImageFullSize        string `json:"imageFullsize" form:"imageFullsize"`
	ImageFile            string `json:"imageFile" form:"imageFile"`
	StickerPackageId     int    `json:"stickerPackageId" form:"stickerPackageId"`
	StickerId            int    `json:"stickerId" form:"stickerId"`
	NotificationDisabled bool   `json:"notificationDisabled" form:"notificationDisabled"`
}

func Test_FillMapStruct(t *testing.T) {
	res := FillMapStruct(MapTest{
		Name:   "Test",
		Age:    10,
		Single: false,
	})

	res1 := FillMapStruct(MapPushMessageRequest{
		Message:              "Test",
		ImageThumbnail:       "",
		ImageFullSize:        "",
		ImageFile:            "",
		StickerPackageId:     0,
		StickerId:            0,
		NotificationDisabled: false,
	})
	t.Log(res)
	t.Log(res1)
}
