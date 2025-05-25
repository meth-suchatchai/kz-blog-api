package usermodels

type UserOTPEnabledRequest struct {
	Enabled bool `json:"enabled" validate:"required"`
}

type UserOTPEnabledResponse struct {
	URI string `json:"uri"`
}
