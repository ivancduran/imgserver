package api

import (
	"fmt"
	"io"

	"github.com/labstack/echo"
)

func CutterHandler(c echo.Context) error {
	// vars := mux.Vars(req)
	// imgparam := vars["imgparam"]
	imgUrl := c.Param("img")
	rImg, err := client.Get(imgUrl)
	defer rImg.Body.Close()

	imgFormat := rImg.Header.Get("Content-Type")
	if err != nil {
		return err
	}

	// w.Header().Set("Content-Length", fmt.Sprint(rImg.ContentLength))
	// w.Header().Set("Content-Type", imgFormat)

	c.Response().Header().Set(echo.HeaderContentLength, fmt.Sprint(rImg.ContentLength))
	c.Response().Header().Set(echo.HeaderContentType, fmt.Sprint(imgFormat))

	if _, err = io.Copy(c.Response().Writer, rImg.Body); err != nil {
		// fmt.Fprintf(w, "Error %d", err)
		panic(err)
	}

	return nil

}
