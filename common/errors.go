package common

import "github.com/pkg/errors"

// captcha
var (
	ErrCaptchaVerifyFailed = errors.New("Captcha verify answer failed")
)

// token
var (
	ErrInvalidToken = errors.New("Invalid Token")
)