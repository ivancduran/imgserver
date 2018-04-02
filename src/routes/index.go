package routes

import (
	"github.com/ivancduran/imgserver/src/api"
	"github.com/labstack/echo"
)

func ApiRoutes(app *echo.Group) {
	// api := app.Group("/private")
	PublicAPI(app)
}

func PublicAPI(app *echo.Group) {
	app.GET("/http", api.CutterHandler)
	app.GET("/up", api.UpHandler)
	app.GET("/get", api.GetHandler)
	app.GET("/install", api.InstallHandler)
}
