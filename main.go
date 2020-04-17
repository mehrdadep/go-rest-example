package main

import (
	"go-rest-example/api/view"
)

func main() {
	// Set mode to release
	//gin.SetMode(gin.ReleaseMode)
	view.StartServer()
}
