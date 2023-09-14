package totp

import (
	"github.com/xlzd/gotp"
	"time"
)

func (c *defaultClient) VerifyAccount(secret string) bool {
	totp := gotp.NewDefaultTOTP(secret)
	otpValue := totp.Now()

	isVerified := totp.Verify(otpValue, time.Now().Unix())

	return isVerified
}
