package api

import (
	"io"
	"log"
	"os"

	"github.com/ivancduran/imgserver/backend/models"
	"github.com/ivancduran/imgserver/backend/utils"
	"github.com/kataras/iris"
)

func UpHandler(ctx *iris.Context) {
	// vars := mux.Vars(req)
	// imgparam := vars["imgparam"]
	imgUrl := ctx.URLParam("i")
	key := ctx.URLParam("k")
	rImg, err := client.Get(imgUrl)
	defer rImg.Body.Close()

	imgFormat := rImg.Header.Get("Content-Type")
	if err != nil {
		errs := models.Response{Response: false}
		ctx.JSON(iris.StatusOK, errs)
		panic(err)
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

	ctx.JSON(iris.StatusOK, res)

}
