package image

import (
	"github.com/disintegration/imaging"
	log "github.com/sirupsen/logrus"
	"image"
	"image/png"
	"os"
)

var (
	DefaultImage image.Image = newImage()
)

func newImage() image.Image {
	src, err := imaging.Open("pic/img.png")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	return src
}

func loadImage(filename string) image.Image {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("os.Open failed: %v", err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatalf("image.Decode failed: %v", err)
	}
	return img
}

func saveImage(filename string, img image.Image) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("os.Create failed: %v", err)
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		log.Fatalf("png.Encode failed: %v", err)
	}
}
