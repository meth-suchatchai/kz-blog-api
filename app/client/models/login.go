package clientmodels

import "time"

type LoginData struct {
	MobileNumber string `json:"mobile_number" required:"true"`
	CountryCode  string `json:"country_code" required:"true"`
	Password     string `json:"password" required:"true"`
}

type LoginRequest struct {
	LoginData
}

type LoginResponse struct {
	Authentication `json:"authentication"`
	LoginUser      `json:"user"`
}

type Authentication struct {
	AccessToken        string    `json:"access_token"`
	AccessTokenExpire  time.Time `json:"access_token_expire"`
	RefreshToken       string    `json:"refresh_token"`
	RefreshTokenExpire time.Time `json:"refresh_token_expire"`
	Domain             string    `json:"domain"`
}

type LoginUser struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	MobileNumber string `json:"mobile_number"`
}
