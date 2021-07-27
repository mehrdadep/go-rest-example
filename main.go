package main

import (
	"github.com/mehrdadep/go-rest-example/api/view"
)

func main() {
	// Set mode to release
	//gin.SetMode(gin.ReleaseMode)
	view.StartServer()
}
