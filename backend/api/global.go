package api

import "net/http"

var formats = map[string]string{
	"image/jpeg": "jpg",
	"image/png":  "png",
	"image/gif":  "gif",
	"image/webp": "webp",
}

var invformats = map[string]string{
	".jpg":  "image/jpeg",
	".png":  "image/png",
	".gif":  "image/gif",
	".webp": "image/webp",
}

var client = http.Client{}

var spath = "../storage/"
