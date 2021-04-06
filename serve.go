package main

import (
	"belajar_golang/Config"
	"belajar_golang/Routes"
)

func main() {
	r := Routes.SetupRoutes()

	Config.ConnectDB() // new

	r.Run()
}
