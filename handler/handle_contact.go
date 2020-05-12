package handler

import "github.com/kataras/iris/v12"

// GetContact function obtain HTML string
func GetContact(ctx iris.Context) {
	ctx.HTML("<h1> Hello from /contact </h1>")
}
