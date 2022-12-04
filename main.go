package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"key-management-service/biz/conf"
	"os"
	"sync"
)

// 优雅退出的hook注册
func shutDownHook(h *server.Hertz) {
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		fmt.Println("shutdown timeout hook")
		<-ctx.Done()
		fmt.Println("exit timeout!")
	})

	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		fmt.Println("shutdown hook")
	})
}

func service(ip string, port string) {
	h := server.Default(
		server.WithHostPorts(ip + ":" + port),
	)
	shutDownHook(h)

	// 这个 register 不太合理，会把所有的路由都分别注册到每一个端口，后续改造
	register(h)
	h.Spin()
}

var wg sync.WaitGroup

func main() {
	// 启动多端口服务
	wg.Add(3)

	// 从配置文件读取配置
	conf.Init()

	// 通过 nginx 反向代理暴露到公网的密钥管理服务
	go func() {
		defer wg.Done()
		// 设置host
		service(os.Getenv("HOST"), "8000")
	}()

	// 通过 nginx 反向代理暴露到公网的加解密工具服务
	go func() {
		defer wg.Done()
		// 设置host
		service(os.Getenv("HOST"), "9001")
	}()

	// 不暴露到公网的本地加解密服务 agent
	go func() {
		defer wg.Done()
		// 设置host
		service(os.Getenv("HOST"), "13001")
	}()

	wg.Wait()
}
