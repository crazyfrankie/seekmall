package main

import (
	"github.com/crazyfrankie/seekmall/app/user/ioc"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	app := ioc.InitApp()

}
