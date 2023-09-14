package usermodels

type User struct {
	Id           int64  `json:"id" mapstructure:"id"`
	MobileNumber string `json:"mobile_number" mapstructure:"mobile_number"`
	Password     string `json:"password" mapstructure:"password_encrypted"`
	Name         string `json:"name" mapstructure:"name"`
	Active       bool   `json:"active" mapstructure:"active"`
	IsTFA        bool   `json:"is_tfa" mapstructure:"is_tfa"`
	TFACode      string `json:"tfa_code" mapstructure:"tfa_code"`
}
