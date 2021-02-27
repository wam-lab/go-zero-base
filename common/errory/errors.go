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

// email
var (
	ErrInvalidVerifyCode = NewCustom(40331, "Invalid email verify code")
)

// user
var (
	ErrUserNotFound      = NewCustom(40411, "User not found")
	ErrUserPwdIncorrect  = NewCustom(40341, "Password is incorrect")
	ErrUserCodeIncorrect = NewCustom(40341, "Email verify code is incorrect")
)
