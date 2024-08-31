package CaptchaUtil

import (
	"encoding/json"
	"net/http"
	"testing"
)

func ResponseReturn(code int, message string, data interface{}) []byte {
	result := map[string]interface{}{"code": code, "message": message, "data": data}
	v, _ := json.Marshal(result)
	return v
}

func InitRouter(t *testing.T) {
	// 获取验证码ID
	http.HandleFunc("/captchaId", func(w http.ResponseWriter, r *http.Request) {
		w.Write(ResponseReturn(200, "获取验证码ID完成", DefaultCaptcha.GetCaptchaId()))
	})

	// 图片或语音验证码刷新地址
	// 图片地址 /captcha/{captchaId}.png
	// 语音地址 /captcha/{captchaId}.wav?lang=zh
	http.Handle("/captcha/", DefaultCaptcha)

	// 登录验证
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		param := map[string]string{}
		json.NewDecoder(r.Body).Decode(&param)

		verifyCaptcha := DefaultCaptcha.VerifyString(param["captchaId"], param["captcha"])

		if !verifyCaptcha {
			w.Write(ResponseReturn(201, "验证码验证失败", nil))
			return
		}

		if param["username"] == "admin" && param["password"] == "123456" {
			w.Write(ResponseReturn(200, "登录完成", nil))
		} else {
			w.Write(ResponseReturn(201, "账号密码错误", nil))
		}

	})

	// 静态文件地址
	fileHandler := http.FileServer(http.Dir("./"))
	http.Handle("/", fileHandler)
}

func TestCaptcha(t *testing.T) {
	InitRouter(t)
	http.ListenAndServe(":8000", nil)
}
