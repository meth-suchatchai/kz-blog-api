package usermodels

type ProfileRequest struct{}

type ProfileResponse struct {
	UserProfile
}

type UserProfile struct {
	Id               int64  `json:"id"`
	Name             string `json:"name"`
	MobileNumber     string `json:"mobile_number"`
	TwoFactorEnabled bool   `json:"two_factor_enabled"`
}
