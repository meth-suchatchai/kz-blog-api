package clientmodels

type VerifyOTPRequest struct {
	VerifyOTPData
}

type VerifyOTPResponse struct {
}

type VerifyOTPData struct {
	OTPReferenceNumber string `json:"otp_reference_number"`
	OTPCode            string `json:"otp_code"`
}
