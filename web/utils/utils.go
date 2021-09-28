package utils

import (
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-micro/v2"
)

// 初始化micro
func InitMicro() micro.Service {
	// 初始化客户端
	consulReg := consul.NewRegistry()
	return micro.NewService(
		micro.Registry(consulReg),
	)
}
