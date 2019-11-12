package service

import "github.com/go-martini/martini"

func NewServer(port string) {
	m := martini.Classic() //创建一个martini实例

	m.Get("/", func(params martini.Params) string { //接受对\的GET方法请求，第二个参数是对请求的处理方法
		return "Hello world!"
	})

	m.RunOnAddr(":" + port)
}

