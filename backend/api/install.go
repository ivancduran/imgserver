package api

import (
	"os"

	"github.com/kataras/iris"
)

func InstallHandler(ctx *iris.Context) {

	source := spath

	if _, err := os.Stat(source); os.IsNotExist(err) {
		os.Mkdir(source, os.ModePerm)
	}

	ctx.Text(iris.StatusOK, "Install completed!")

}
