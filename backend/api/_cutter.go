package api

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/oliamb/cutter"
)

func Cut(, w string, h string) {

	infile, err := os.Open("./sample.jpg")
	if err != nil {
		// replace this with real error handling
		panic(err)
	}
	defer infile.Close()

	// Decode will figure out what type of image is in the file on its own.
	// We just have to be sure all the image packages we want are imported.
	src, _, err := image.Decode(infile)
	if err != nil {
		// replace this with real error handling
		panic(err)
	}

	croppedImg, err := cutter.Crop(src, cutter.Config{
		Width:  66,
		Height: 66,
		Anchor: image.Point{81, 22},
	})

	outfilename := "result.jpeg"
	outfile, err := os.Create(outfilename)
	if err != nil {
		panic(err.Error())
	}
	defer outfile.Close()
	jpeg.Encode(outfile, croppedImg, nil)

}
