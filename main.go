package main

import (
	"go-blog-api/app"
)

func main() {
	app:= &app.App{}
	app.Init()
	app.Run(":5000")
}