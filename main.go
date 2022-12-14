package main

import (
	"Qianshou/src/global"
	"Qianshou/src/routing/router"
	"Qianshou/src/setting"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy

func init() {
	log.SetFlags(log.Ltime | log.Llongfile)
}

// @title        chat
// @version      1.0
// @description  在线聊天系统

// @license.name  raja,chong
// @license.url

// @host      chat.humraja.xyz
// @BasePath  /

// @securityDefinitions.basic  BasicAuth
func main() {
	setting.AllInit()
	if global.PbSettings.Server.RunMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := router.NewRouter() // 注册路由
	s := &http.Server{
		Addr:           global.PbSettings.Server.Address,
		Handler:        r,
		MaxHeaderBytes: 1 << 20,
	}
	global.Logger.Info("Server started!")
	fmt.Println("AppName:", global.PbSettings.App.Name, "Version:", global.PbSettings.App.Version, "Address:", global.PbSettings.Server.Address, "RunMode:", global.PbSettings.Server.RunMode)
	errChan := make(chan error, 1)
	defer close(errChan)
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			errChan <- err
		}
	}()
	// 优雅退出
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case err := <-errChan:
		global.Logger.Error(err.Error())
	case <-quit:
		global.Logger.Info("ShutDown Server...")
		// 给几秒完成剩余任务
		ctx, cancel := context.WithTimeout(context.Background(), global.PbSettings.Server.DefaultContextTimeout)
		defer cancel()
		if err := s.Shutdown(ctx); err != nil { // 优雅退出
			if !errors.Is(err, context.DeadlineExceeded) {
				global.Logger.Error("Server forced to ShutDown,Err:" + err.Error())
			}
		}
	}
	global.Logger.Info("Server exited!")
}
