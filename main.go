package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"wxcloudrun-golang/service"
)

func main() {
	app := iris.New()
	app.Get("/hello", func(context iris.Context) {
		context.WriteString("hello")
	})
	app.Get("/code2id", service.Code2Id)
	app.Get("/login", service.Login)

	app.Listen(fmt.Sprintf("0.0.0.0:%d", 80))
}
