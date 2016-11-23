package api

import (
	"fmt"
	"io"

	"github.com/kataras/iris"
)

func CutterHandler(ctx *iris.Context) {
	// vars := mux.Vars(req)
	// imgparam := vars["imgparam"]
	imgUrl := ctx.URLParam("img")

	rImg, err := client.Get(imgUrl)
	defer rImg.Body.Close()

	imgFormat := rImg.Header.Get("Content-Type")
	if err != nil {
		// fmt.Fprintf(w, "Error %d", err)

		return
	}

	// w.Header().Set("Content-Length", fmt.Sprint(rImg.ContentLength))
	// w.Header().Set("Content-Type", imgFormat)

	ctx.Response.Header.Set("Content-Length", fmt.Sprint(rImg.ContentLength))
	ctx.SetContentType(imgFormat)

	if _, err = io.Copy(ctx.Response.BodyWriter(), rImg.Body); err != nil {
		// fmt.Fprintf(w, "Error %d", err)
		panic(err)
	}

}
