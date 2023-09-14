package totp

type Client interface {
	GenerateTOTP(provision string) (string, string)
	VerifyAccount(secret string) bool
}
