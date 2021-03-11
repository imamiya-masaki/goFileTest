package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image/png"
	"os"
)

func main () {
	// link := "./asset/graph"
	// fmt.Println(link, p)
	// fmt.Println(j, err)
	im := gocv.IMRead("./asset/graph.png",1)
	defer im.Close()
	var im_gray gocv.Mat
	// var im_blur gocv.Mat
	gocv.CvtColor(im, &im_gray, gocv.ColorBGRToGray)
	// gocv.GaussianBlur(im_gray, &im_blur, image.Point{11,11}, 0, 0, gocv.BorderDefault)
	// output, err := os.Create("./output/graph.png")
	// defer output.Close()
	image, err := im_gray.ToImage()
	if err != nil {
		fmt.Println(err)
		return
	}
	file, errOs := os.Create("./output/graph2.png")
	if errOs != nil {
		fmt.Println(err)
		return
	}
	if err := png.Encode(file, image); err != nil {
		fmt.Println(err)
		return
	}
}