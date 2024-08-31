package CaptchaUtil

import (
	"net/http"

	captchas "github.com/dchest/captcha"
)

type Captcha struct {
	Length int
	Width  int
	Height int
}

var DefaultCaptcha = &Captcha{
	Length: captchas.DefaultLen,
	Width:  captchas.StdWidth,
	Height: captchas.StdHeight,
}

// 获取图片二维码地址
func (c *Captcha) GetCaptchaId() string {
	return captchas.NewLen(c.Length)
}

// 实现接口handler
func (c *Captcha) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	captchas.Server(c.Width, c.Height).ServeHTTP(w, r)
}

func (c *Captcha) VerifyString(id string, input string) bool {
	return captchas.VerifyString(id, input)
}
