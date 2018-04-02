package api

import (
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/ivancduran/imgserver/src/models"
	"github.com/ivancduran/imgserver/src/utils"
	"github.com/labstack/echo"
	"github.com/mssola/user_agent"
)

func GetHandler(c echo.Context) error {
	imgUrl := c.QueryParam("i")
	width := c.QueryParam("w")
	height := c.QueryParam("h")
	key := c.QueryParam("k")
	trans := c.QueryParam("t")
	quality := c.QueryParam("q")
	webp := c.QueryParam("c")
	percent := c.QueryParam("p")
	profile := c.QueryParam("wp")
	face := c.QueryParam("f")
	// effect := c.QueryParam("e")

	Ua := c.Request().UserAgent()
	ua := user_agent.New(Ua)
	browser, _ := ua.Browser()
	widthN := 0
	heightN := 0
	qualityN := 90
	percentN := 100
	faceN := 0

	if profile != "" {
		switch profile {
		case "web":
			quality = "80"
			percent = "100"
		case "lossless":
			quality = "92"
			percent = "100"
		}
	}

	if width != "" {
		i, _ := strconv.Atoi(width)
		widthN = i
	}
	if height != "" {
		i, _ := strconv.Atoi(height)
		heightN = i
	}

	if quality != "" {
		i, _ := strconv.Atoi(quality)
		qualityN = i
	} else {
		quality = "90"
	}

	if widthN == 0 && heightN == 0 && trans != "face" {
		trans = "default"
	}

	if percent != "" {
		i, _ := strconv.Atoi(percent)
		percentN = i
	}

	if face != "" {
		i, _ := strconv.Atoi(face)
		faceN = i - 1
	}

	// // testing get vars
	// fmt.Println("imageurl: " + imgUrl)
	// fmt.Println("width: " + width)
	// fmt.Println("height: " + height)
	// fmt.Println("key: " + key)
	// fmt.Println("trans: " + trans)
	// fmt.Println("quality" + quality)
	// fmt.Println(qualityN)
	// fmt.Println("percent: " + percent)
	// fmt.Println("profile: " + profile)

	name := imgUrl + strconv.Itoa(widthN) + strconv.Itoa(heightN) + trans + strconv.Itoa(faceN) + strconv.Itoa(qualityN) + strconv.Itoa(percentN)

	var thumb *image.NRGBA
	ext := filepath.Ext(imgUrl)

	file := spath + key + "/flavors/" + name + ext

	fmt.Println(file)

	// if file not exist
	if _, err := os.Stat(file); os.IsNotExist(err) {

		fmt.Println("Processing Image...")

		source := spath + key + "/sources/" + imgUrl
		fmt.Println(source)

		img, err := imaging.Open(source)

		b := img.Bounds()
		imgWidth := b.Max.X
		imgHeight := b.Max.Y

		if percent != "" {
			pct, _ := utils.ParseFloatPercent(percent+"%", 64)

			x := float64(imgWidth)
			y := float64(imgHeight)

			widthN = int(x * pct)
			heightN = int(y * pct)

			imgWidth = widthN
			imgHeight = heightN

			if err != nil {
				panic(err)
			}
		}

		// dst := imaging.New(widthN, heightN, color.NRGBA{0, 0, 0, 0})

		switch trans {
		case "resize":
			thumb = imaging.Resize(img, widthN, heightN, imaging.Box)

		case "fill":
			thumb = imaging.Fill(img, widthN, heightN, imaging.Center, imaging.Box)

		case "fit":
			thumb = imaging.Fit(img, widthN, heightN, imaging.Box)

		case "face":
			nf := FaceCMD(source)

			if faceN <= len(nf) || faceN == 0 {
				cut := Cut(img, nf[faceN][0], nf[faceN][1], nf[faceN][2], nf[faceN][3])
				thumb = cut.(*image.NRGBA)
			}

		default:
			trans = "default"
			if widthN != 0 && heightN != 0 {
				thumb = imaging.Resize(img, widthN, heightN, imaging.Box)
			} else {
				thumb = img.(*image.NRGBA)
				// thumb = imaging.Resize(img, imgWidth, imgHeight, imaging.Box)
			}
		}

		// dst = imaging.Paste(dst, thumb, image.Pt(0, 0))
		err = imaging.Save(thumb, spath+key+"/flavors/"+name+ext)
		if err != nil {
			panic(err)
		}

		// buffer := new(bytes.Buffer)
		switch invformats[ext] {
		case "image/jpeg":
			cmd := "jpegoptim --strip-all --all-progressive -m " + quality + " " + spath + key + "/flavors/" + name + ext
			fmt.Println(cmd)
			out, _ := exec.Command(cmd).Output()

			fmt.Println(string(out))

		case "image/png":

			cmd := "optipng -o2 -strip all " + spath + key + "/flavors/" + name + ext
			out, _ := exec.Command("sh", "-c", cmd).Output()

			fmt.Println(string(out))

		}

		t := models.Transform{
			Response:  true,
			Code:      name,
			Url:       name,
			Format:    invformats[ext],
			Extension: ext,
			Width:     widthN,
			Height:    heightN,
			Face:      faceN,
			Transform: trans,
			Bucket:    key,
		}

		t.Save()

	}

	// validate webp convertion by get params
	if webp != "false" {
		// validate webp by browser
		if webp == "true" || browser == "Chrome" || browser == "Opera" {
			file = spath + key + "/flavors/" + name + ".webp"

			if _, err := os.Stat(file); os.IsNotExist(err) {
				fmt.Println("Processing Webp")

				dir := "cwebp"

				input := spath + key + "/flavors/" + name + ext
				output := spath + key + "/flavors/" + name + ".webp"
				com := "-o"

				cmds := exec.Command(dir, input, com, output)
				fmt.Println(dir)
				fmt.Println(input)
				fmt.Println(output)
				err := cmds.Run()
				cmds.Wait()

				if err != nil {
					fmt.Println(err)
				}

				t := models.Transform{
					Response:  true,
					Code:      name,
					Url:       name,
					Format:    invformats[ext],
					Extension: ".webp",
					Width:     widthN,
					Height:    heightN,
					Face:      faceN,
					Transform: trans,
					Bucket:    key,
				}

				t.Save()

			}
			ext = ".webp"

		}
	}

	imgs, err := ioutil.ReadFile(file)

	if err != nil {
		return err
	}

	// w.Header().Set("Vary", "Accept-Encoding")
	// w.Header().Set("Cache-Control", "max-age=2592000, public, must-revalidate, proxy-revalidate")
	// w.Header().Set("Expires", "Thu, 01 Dec 2040 16:00:00 GMT")
	// w.Header().Set("Last-Modified", "Tue, 15 Nov 1994 12:45:26 GMT")

	c.Response().Header().Set(echo.HeaderContentType, invformats[ext])
	c.Response().Write(imgs)

	return nil
}
