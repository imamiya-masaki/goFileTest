package main

import (
	"gocv.io/x/gocv"
	"image"
	"os"
)

func main () {
	// link := "./asset/graph"
	p , err := os.Open("./asset/graph.png")
	// fmt.Println(link, p)
	// fmt.Println(j, err)
	if err != nil {
		panic(err)
	}
	defer p.Close()
	im := gocv.IMRead("./asset/graph.png",1)
	var im_gray gocv.Mat
	var im_blur gocv.Mat
	gocv.CvtColor(im, &im_gray, gocv.ColorBGRToGray)
	gocv.GaussianBlur(im_gray, &im_blur, image.Point{11,11}, 0, 0, gocv.BorderDefault)
	// output, err := os.Create("./output/graph.png")
	// defer output.Close()
	gocv.IMWrite("./output/graph2.png", im_blur)
}