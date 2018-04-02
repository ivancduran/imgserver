package api

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ivancduran/imgserver/src/models"
	"github.com/ivancduran/imgserver/src/utils"
	"github.com/labstack/echo"
)

func UpHandler(c echo.Context) error {
	// vars := mux.Vars(req)
	// imgparam := vars["imgparam"]
	imgUrl := c.QueryParam("i")
	key := c.QueryParam("k")
	rImg, err := client.Get(imgUrl)
	defer rImg.Body.Close()

	imgFormat := rImg.Header.Get("Content-Type")
	if err != nil {
		errs := models.Response{Response: false}
		return c.JSON(http.StatusOK, errs)
	}

	ext := formats[imgFormat]

	rcode := utils.RandString(16)

	path := spath + key + "/"
	source := path + "sources/"
	flavors := path + "flavors/"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)

		if _, err := os.Stat(source); os.IsNotExist(err) {
			os.Mkdir(source, os.ModePerm)
		}

		if _, err := os.Stat(flavors); os.IsNotExist(err) {
			os.Mkdir(flavors, os.ModePerm)
		}

	}

	path = source

	file, err := os.Create(path + rcode + "." + ext)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(file, rImg.Body)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	res := models.Response{true, rcode, imgFormat, ext}

	u := models.Upload{
		Response:  true,
		Url:       imgUrl,
		Code:      rcode,
		Format:    imgFormat,
		Extension: ext,
	}
	u.Save()

	return c.JSON(http.StatusOK, res)

}
