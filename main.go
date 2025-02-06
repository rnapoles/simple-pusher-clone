package main

import (
	"simple-pusher-clone/server"
	"simple-pusher-clone/utils"
)

func main() {
	utils.InitEnv()
	server.StartServer()
}
