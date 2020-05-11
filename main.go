package main

import "github.com/kataras/iris"

func main() {
	app := iris.Default()

	// Method:   GET
	// Resource: http://localhost:3000/
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Hello world!</h1>")
	})

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:3000/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	// Method:   GET
	// Resource: http://localhost:3000/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello iris web framework."})
	})

	// http://localhost:3000
	// http://localhost:3000/ping
	// http://localhost:3000/hello
	app.Run(iris.Addr(":3000"))
}