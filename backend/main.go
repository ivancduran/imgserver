package main

import (
	"github.com/ivancduran/imgserver/backend/api"
	"github.com/ivancduran/imgserver/backend/db"
	"github.com/ivancduran/imgserver/backend/libs"
	"github.com/ivancduran/imgserver/backend/routes"

	"github.com/iris-contrib/middleware/logger"
	"github.com/iris-contrib/middleware/recovery"
	"github.com/kataras/go-template/html"
	"github.com/kataras/iris"
)

var cfg = libs.GetConf()

func main() {

	if cfg.Devel {
		iris.Config.IsDevelopment = true
		iris.Config.Gzip = false
		// set the global middlewares
		// iris.Use(logger.New(iris.Logger))
		iris.Use(logger.New())
	} else {
		iris.Config.IsDevelopment = false
		iris.Config.Gzip = true
		iris.Config.DisableBanner = true
		// set Revocery ON panic
		iris.Use(recovery.New())
	}

	// set the template engine
	iris.UseTemplate(html.New(html.Config{Layout: "layout.html"})).Directory("../frontend/templates", ".html")
	// set the favicon
	iris.Favicon("../frontend/public/images/favicon.ico")

	// set static folder(s)
	iris.Static("/public", "../frontend/public", 1)

	// set the global middlewares
	iris.Use(logger.New())

	// set the custom errors
	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.Render("errors/404.html", iris.Map{"Title": iris.StatusText(iris.StatusNotFound)})
	})

	iris.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
		ctx.Render("errors/500.html", nil, iris.RenderOptions{"layout": iris.NoLayout})
	})

	// register the routes & the public API
	registerRoutes()
	registerAPI()

	// start the server
	iris.Listen(cfg.Host)
}

func registerRoutes() {
	iris.Get("/", routes.Home)
}

func registerAPI() {

	iris.Get("/http", api.CutterHandler)
	iris.Get("/up", api.UpHandler)
	iris.Get("/get", api.GetHandler)
	iris.Get("/install", api.InstallHandler)

	u := api.AuthAPI{}
	iris.Post("/v1/auth/register", u.Register)
	iris.Post("/v1/auth/login", u.Login)
	iris.Post("/v1/auth/logout", u.Logout)
}

func DbMain() {
	// Database Main Conexion
	Db := db.New()
	// Index keys
	keys := []string{"email"}
	Db.Index("users", keys)

}
