package main

import (
	"github.com/ig-rocker/jukebox/initializers"
	"github.com/ig-rocker/jukebox/routes"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.SyncDB()

}
func main() {

	//starting the server
	routes.Run()

}
