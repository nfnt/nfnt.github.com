/*
Copyright (c) 2012, Jan Schlicht <jan.schlicht@gmail.com>

Permission to use, copy, modify, and/or distribute this software for any purpose
with or without fee is hereby granted, provided that the above copyright notice
and this permission notice appear in all copies.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND
FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS
OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER
TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF
THIS SOFTWARE.
*/

package main

import (
	"github.com/nfnt/resize"
	"image"
	_ "image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"runtime"
)

var images = map[string]string{
	"rings":    "http://nfnt.github.com/img/rings_lg_orig.png",
	"IMG_3694": "http://nfnt.github.com/img/IMG_3694_720.jpg",
}

var filters = map[string]resize.InterpolationFunction{
	"NearestNeighbor":   resize.NearestNeighbor,
	"Bilinear":          resize.Bilinear,
	"Bicubic":           resize.Bicubic,
	"MitchellNetravali": resize.MitchellNetravali,
	"Lanczos2":          resize.Lanczos2,
	"Lanczos3":          resize.Lanczos3,
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// Create resized sample images
func main() {
	for name, uri := range images {
		resp, err := http.Get(uri)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		img, _, err := image.Decode(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		resizeWithAllFilters(name, img)
	}
}

func resizeWithAllFilters(imgName string, img image.Image) {
	for name, filter := range filters {
		m := resize.Resize(300, 0, img, filter)

		fileName := imgName + "_300_" + name + ".png"

		writeImageToFile(fileName, m)
	}
}

func writeImageToFile(fileName string, img image.Image) error {
	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()

	png.Encode(out, img)

	return nil
}
