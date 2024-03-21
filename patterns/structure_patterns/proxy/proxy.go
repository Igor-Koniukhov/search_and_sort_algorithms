package main

import "fmt"

type Image interface {
	DisplayImage()
}

type RealImage struct {
	filename string
}

func (r *RealImage) DisplayImage() {
	fmt.Println("Displaying", r.filename)
}

func NewRealImage(filename string) *RealImage {
	fmt.Println("Loading image", filename)
	return &RealImage{filename: filename}
}

type ProxyImage struct {
	realImage *RealImage
	filename  string
}

func (p *ProxyImage) DisplayImage() {
	if p.realImage == nil {
		p.realImage = NewRealImage(p.filename)
	}
	p.realImage.DisplayImage()
}

func NewProxyImage(filename string) *ProxyImage {
	return &ProxyImage{filename: filename}
}

func main() {
	image1 := NewProxyImage("photo1.jpg")
	image2 := NewProxyImage("photo2.jpg")

	// Image will be loaded and displayed
	image1.DisplayImage()

	// Image will not be loaded again, but displayed directly
	image1.DisplayImage()

	// Image will be loaded and displayed
	image2.DisplayImage()
}
