package clientmodels

type LoginData struct {
	MobileNumber int    `json:"mobile_number"`
	CountryCode  int    `json:"country_code"`
	Password     string `json:"password"`
}

type LoginRequest struct {
	LoginData
}

type LoginResponse struct {
	Authentication `json:"authentication"`
	LoginUser      `json:"user"`
}

type Authentication struct {
	AccessToken        string `json:"access_token"`
	AccessTokenExpire  int64  `json:"access_token_expire"`
	RefreshToken       string `json:"refresh_token"`
	RefreshTokenExpire int64  `json:"refresh_token_expire"`
	Domain             string `json:"domain"`
}

type LoginUser struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	MobileNumber string `json:"mobile_number"`
}
