package errory

// captcha
var (
	ErrCaptchaVerifyFailed = NewCustom(40311, "Captcha verify answer failed")
)

// token
var (
	ErrInvalidToken            = NewCustom(40321, "Invalid token")
	ErrUnexpectedSigningMethod = NewCustom(40322, "Unexpected signing method")
)