package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"getCaptcha/handler"
	"github.com/micro/go-plugins/registry/consul/v2"

	getCaptcha "getCaptcha/proto/getCaptcha"
)

func main() {
	//初始化consul
	consulReg := consul.NewRegistry()

	// New Service
	service := micro.NewService(
		micro.Address("192.168.1.8:12345"), //防止随机生成端口
		micro.Name("go.micro.service.getCaptcha"),
		micro.Registry(consulReg), //添加注册
		micro.Version("latest"),
	)


	// Register Handler
	getCaptcha.RegisterGetCaptchaHandler(service.Server(), new(handler.GetCaptcha))


	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
