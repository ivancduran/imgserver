package api

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func InstallHandler(c echo.Context) error {

	source := spath

	if _, err := os.Stat(source); os.IsNotExist(err) {
		os.Mkdir(source, os.ModePerm)
	}

	return c.String(http.StatusOK, "install")
}
