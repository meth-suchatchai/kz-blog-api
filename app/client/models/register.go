package clientmodels

type RegisterRequest struct {
	RegisterData
}

type RegisterResponse struct {
	OTPReferenceNumber string `json:"otp_reference_number"`
}

type RegisterData struct {
	MobileNumber int    `json:"mobile_number"`
	CountryCode  int    `json:"country_code"`
	Password     string `json:"password"`
	Name         string `json:"name"`
}

type RegisterOTPUser struct {
	Id                 int64  `json:"id"`
	Name               string `json:"name"`
	MobileNumber       int    `json:"mobile_number"`
	MessageId          string `json:"message_id"`
	OTPReferenceNumber string `json:"otp_reference_number"`
}
