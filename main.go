package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"sync"

	"github.com/dchest/captcha"
)

// var (
// 	Router = http.DefaultServeMux
// )

type GoContext struct {
	R *http.Request
	W http.ResponseWriter

	s map[string]interface{}
	m *sync.Mutex
}

func NewGoContext(w http.ResponseWriter, r *http.Request) *GoContext {
	return &GoContext{
		R: r,
		W: w,
		s: map[string]interface{}{},
		m: &sync.Mutex{},
	}
}

func (g *GoContext) Get(key string) interface{} {
	if v, ok := g.s[key]; ok {
		return v
	}
	return nil
}

func (g *GoContext) Set(key string, value interface{}) {
	g.m.Lock()
	g.s[key] = value
	g.m.Unlock()
}

func main() {

	http.HandleFunc("/captcha", func(w http.ResponseWriter, r *http.Request) {
		captchaId := captcha.NewLen(5)
		fmt.Println(r.URL.Host, r.URL.Path)
		fmt.Printf("%#v", r.URL.RequestURI())
		v, _ := json.Marshal(map[string]interface{}{"code": 200, "message": "获取验证码", "data": captchaId})
		w.Write(v)
	})

	http.HandleFunc("/img/", func(w http.ResponseWriter, r *http.Request) {
		captcha.Server(captcha.StdWidth, captcha.StdHeight).ServeHTTP(w, r)
	})

	srv := http.Server{
		Addr:    ":8000",
		Handler: nil,
	}
	log.Fatal(srv.ListenAndServe())
}

func get() {
	url := "https://github.com/dchest/captcha/12222.png"
	dir, file := path.Split(url)
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	print(dir, id)
	print(path.Base(dir))

}
