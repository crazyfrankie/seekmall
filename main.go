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

	"github.com/joho/godotenv"

	"github.com/crazyfrankie/seekmall/config"
	"github.com/crazyfrankie/seekmall/ioc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	app := ioc.InitApp()

	server := &http.Server{
		Addr:    config.GetConf().Server.Addr,
		Handler: app.Server,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("Server is running %s\n", config.GetConf().Server.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server forces shutting down: %s\n", err)
	}

	log.Println("Server exited gracefully")
}
