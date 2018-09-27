package api

import iris "github.com/kataras/iris"

func Index() func(ctx iris.Context) {
	return func(ctx iris.Context) {
		// Bind: {{.message}} with Message
		ctx.ViewData("message", "WIM Training Landing API test ttttt!")
		// Render template file: ./views/index.html
		ctx.View("index.html")
	}
}

func ParamFunction() func(ctx iris.Context) {
	return func(ctx iris.Context) {
		userID, _ := ctx.Params().GetInt64("id")
		ctx.Writef("User ID: %d", userID)
	}
}
