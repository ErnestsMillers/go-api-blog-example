package main

import (
	"os"

	"go-api-blog-example/server"
	"go-api-blog-example/utils"
)

func main() {
	app := server.New()
	utils.InfoLog.Println("starting the server on port", os.Getenv("API_PORT"))
	if err := app.Listen(os.Getenv("API_PORT")); err != nil {
		utils.ErrorLog.Panic(err.Error())
	}
}
