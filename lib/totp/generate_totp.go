package totp

import (
	"github.com/xlzd/gotp"
)

func (c *defaultClient) GenerateTOTP(provision string) (string, string) {
	secret := gotp.RandomSecret(16)
	totp := gotp.NewDefaultTOTP(secret)

	uri := totp.ProvisioningUri(provision, c.name)

	return secret, uri
}
