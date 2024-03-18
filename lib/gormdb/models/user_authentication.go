package dbmodels

import (
	"gorm.io/gorm"
	"time"
)

type UserAuthentication struct {
	gorm.Model
	UserId             uint      `gorm:"column:user_id;type:bigint;uniqueIndex"`
	MobileNumber       string    `gorm:"column:mobile_number;type:varchar(10);uniqueIndex"`
	CountryCode        string    `gorm:"column:country_code;type:varchar(5);not null"`
	AccessToken        string    `gorm:"column:access_token;type:text;not null"`
	AccessTokenExpire  time.Time `gorm:"column:access_token_expire;type:timestamp;not null"`
	RefreshToken       string    `gorm:"column:refresh_token;type:text;not null"`
	RefreshTokenExpire time.Time `gorm:"column:refresh_token_expire;type:timestamp;not null"`
}

type UpdateUserAuthentication struct {
	UserId             uint       `gorm:"column:user_id;type:bigint;uniqueIndex"`
	MobileNumber       string     `gorm:"column:mobile_number;type:varchar(10);uniqueIndex"`
	CountryCode        string     `gorm:"column:country_code;type:varchar(5);not null"`
	AccessToken        *string    `gorm:"column:access_token;type:text;not null"`
	AccessTokenExpire  *time.Time `gorm:"column:access_token_expire;type:timestamp;not null"`
	RefreshToken       *string    `gorm:"column:refresh_token;type:text;not null"`
	RefreshTokenExpire *time.Time `gorm:"column:refresh_token_expire;type:timestamp;not null"`
}
