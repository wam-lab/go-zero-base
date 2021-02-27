package svc

import (
	"github.com/mojocn/base64Captcha"
	"github.com/yguilai/timetable-micro/services/captcha/rpc/internal/config"
	"github.com/yguilai/timetable-micro/services/captcha/rpc/internal/store"
	"image/color"
)

type ServiceContext struct {
	c       config.Config
	Captcha *base64Captcha.Captcha
	Store   *store.RedisStore
}

func NewServiceContext(c config.Config) *ServiceContext {
	d := newDriver(c.Captcha)
	s := store.NewRedisStore(c.Redis.NewRedis(), c.Captcha.Expire)
	return &ServiceContext{
		c:       c,
		Store:   s,
		Captcha: base64Captcha.NewCaptcha(d, s),
	}
}

func newDriver(conf config.CaptchaConfig) (driver base64Captcha.Driver) {
	bg := newDriverBg()
	fs := []string{"wqy-microhei.ttc"}
	switch conf.Driver {
	case "string":
		driver = newDriverString(conf, bg, fs)
	case "math":
		driver = newDriverMath(conf, bg, fs)
	case "number":
		fallthrough
	default:
		driver = newDriverDigit(conf)
	}
	return
}

func newDriverBg() *color.RGBA {
	return &color.RGBA{
		R: 0,
		G: 0,
		B: 0,
		A: 125,
	}
}

func newDriverString(conf config.CaptchaConfig, bg *color.RGBA, fs []string) base64Captcha.Driver {
	return (&base64Captcha.DriverString{
		Height:          conf.Height,
		Width:           conf.Width,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          conf.Length,
		Source:          "1234567890qwertyuiopasdfghjklzxcvbnm",
		BgColor:         bg,
		Fonts:           fs,
	}).ConvertFonts()
}

func newDriverMath(conf config.CaptchaConfig, bg *color.RGBA, fs []string) base64Captcha.Driver {
	return (&base64Captcha.DriverMath{
		Height:          conf.Height,
		Width:           conf.Width,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		BgColor:         bg,
		Fonts:           fs,
	}).ConvertFonts()
}

func newDriverDigit(conf config.CaptchaConfig) base64Captcha.Driver {
	return &base64Captcha.DriverDigit{
		Height:   conf.Height,
		Width:    conf.Width,
		Length:   conf.Length,
		MaxSkew:  0,
		DotCount: 0,
	}
}
