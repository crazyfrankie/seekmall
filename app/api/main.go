package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/crazyfrankie/seekmall/app/api/config"
	"github.com/crazyfrankie/seekmall/app/api/ioc"
)

func main() {
	engine := ioc.InitGin()

	server := &http.Server{
		Addr:    config.GetConf().Server.Addr,
		Handler: engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("Server is running: %s\n", config.GetConf().Server.Addr)
	// 创建通道监听信号
	quit := make(chan os.Signal, 1)

	// 监听信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 阻塞直到收到信号
	<-quit
	log.Println("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅地关闭服务器
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server forced shutting down:%s\n", err.Error())
	}

	log.Println("Server exited gracefully")
}
