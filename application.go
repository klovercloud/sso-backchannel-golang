package main

import (
	"backChannel/config"
	"backChannel/router"
	"backChannel/server"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Starting Application")
	err := config.InitEnVars()
	if err != nil {
		fmt.Println("[ERROR]: ", err.Error())
	}
	srv := server.New()
	router.Routes(srv)
	srv.Logger.Fatal(srv.Start(":" + strconv.Itoa(config.ServerPort)))

}
