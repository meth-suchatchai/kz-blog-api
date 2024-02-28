package kzline

type LineNotification interface {
	PushMessage(request PushMessageRequest) (*PushMessageResponse, error)
	GetApiStatus() (*StatusResponse, error)
}
