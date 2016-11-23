package routes

import "github.com/kataras/iris"

func Home(ctx *iris.Context) {
	ctx.Text(iris.StatusOK, "Home!")
}
