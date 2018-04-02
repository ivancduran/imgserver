package api

import (
	"fmt"
	"image"

	"github.com/oliamb/cutter"
)

func Cut(img image.Image, ws, hs, we, he int) image.Image {

	// infile, err := os.Open("./sample.jpg")
	// if err != nil {
	// 	// replace this with real error handling
	// 	panic(err)
	// }
	// defer infile.Close()

	// // Decode will figure out what type of image is in the file on its own.
	// // We just have to be sure all the image packages we want are imported.
	// src, _, err := image.Decode(infile)
	// if err != nil {
	// 	// replace this with real error handling
	// 	panic(err)
	// }

	fmt.Println("start cutting")

	croppedImg, err := cutter.Crop(img, cutter.Config{
		Width:  we,
		Height: he,
		Anchor: image.Point{ws, hs},
	})

	if err != nil {
		panic(err.Error())
	}

	return croppedImg

	// outfilename := spath + "/result.jpeg"
	// outfile, err := os.Create(outfilename)

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer outfile.Close()
	// jpeg.Encode(outfile, croppedImg, nil)

}
