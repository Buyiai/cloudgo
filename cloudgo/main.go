package main

import (
	"os"

	"github.com/github-user/cloudgo/service"
	flag "github.com/spf13/pflag"
)

const (
	PORT string = "8080" //设置默认端口为8080
)

func main() {
	//如果没有监听到端口，则设为默认端口
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	//允许用户可以通过-p来设置端口
	pPort := flag.StringP("port", "p", "PORT", "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	service.NewServer(port) //新建服务
	//server.Run(":" + port) //服务启动
}
