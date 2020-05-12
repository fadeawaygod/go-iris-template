package main

import (
	"go-iris-template/handler"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()
	//app.Use(myMiddleware)

	// general
	app.Handle("GET", "/contact", handler.GetContact)
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello iris web framework."})
	})
	app.Get("/group/{id:uint64}", func(ctx iris.Context) {
		userID, _ := ctx.Params().GetUint64("id")
		ctx.Writef("group ID: %d", userID)
	})

	// sever-side rendering
	app.RegisterView(iris.HTML("./views", ".html"))
	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", "Hello world!")
		ctx.View("hello.html")
	})

	//Grouping Routes
	app.PartyFunc("/users", func(users iris.Party) {
		users.Use(func(ctx iris.Context) {
			ctx.WriteString("users_praty_prifix")
		})
		// http://localhost:8080/users/42/profile
		users.Get("/{id:uint64}/profile", func(ctx iris.Context) {
			ctx.WriteString("profile")
		})
		// http://localhost:8080/users/messages/1
		users.Get("/messages/{id:uint64}", func(ctx iris.Context) {
			ctx.WriteString("messages")
		})
	})

	app.Run(iris.Addr(":3000"))
}
