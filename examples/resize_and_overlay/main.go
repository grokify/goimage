package main

import (
	"fmt"
	"log"

	"github.com/grokify/mogo/errors/errorsutil"
	"github.com/grokify/mogo/image/imageutil"
)

// https://www.golangprograms.com/how-to-add-watermark-or-merge-two-image.html

func main() {
	fileTop := "_top.png"
	fileBkg := "_background.png"
	fileOut := "_output.png"

	imgTop, _, err := imageutil.ReadImageFile(fileTop)
	if err != nil {
		log.Fatal(errorsutil.Wrap(err, fileTop))
	}
	imgBkg, _, err := imageutil.ReadImageFile(fileBkg)
	if err != nil {
		log.Fatal(errorsutil.Wrap(err, fileBkg))
	}

	imgTop = imageutil.AddBackgroundWhite(imgTop)
	imgTop = imageutil.Resize(120, 0, imgTop, imageutil.ScalerBest())
	imgOut := imageutil.OverlayMore(imgBkg, imgTop, "middleleft", 0, 0)
	err = imageutil.Image{Image: imgOut}.WritePNGFile(fileOut)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DONE")
}
