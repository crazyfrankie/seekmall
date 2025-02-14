package main

import (
	"github.com/joho/godotenv"

	"github.com/crazyfrankie/seekmall/app/user/ioc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	server := ioc.InitServer()

	err = server.Serve()
	if err != nil {
		panic(err)
	}
}
