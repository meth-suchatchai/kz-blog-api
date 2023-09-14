package clientmodels

type SuccessResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func CreateSuccessResponse(data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Code:    0,
		Message: "",
		Data:    data,
	}
}
