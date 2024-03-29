package main

import (
	"fmt"
	"image/jpeg"
	"log"

	"github.com/grokify/mogo/image/imageutil"
)

func main() {
	urlsmatrix := [][]string{
		{"https://raw.githubusercontent.com/grokify/goimage/master/read_testdata/gopher_appengine_color.jpg"},
		{
			"https://raw.githubusercontent.com/grokify/goimage/master/read_testdata/gopher_color.jpg",
			"https://raw.githubusercontent.com/grokify/goimage/master/read_testdata/gopher_color.jpg"},
	}

	outfile := "_merged.jpg"

	matrix, err := imageutil.MatrixRead(urlsmatrix)
	if err != nil {
		log.Fatal(err)
	}

	i2 := imageutil.Image{Image: matrix.Merge(true, true)}
	err = i2.WriteJPEGFile(
		outfile,
		&imageutil.JPEGEncodeOptions{
			Options: &jpeg.Options{Quality: imageutil.JPEGQualityMax},
		})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("WROTE [%v]\n", outfile)

	fmt.Println("DONE")
}
