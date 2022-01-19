package image

import (
	"fmt"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"log"
	"testing"
)

func TestImaging(t *testing.T) {
	// Open a test image.
	src, err := imaging.Open("pic/img.png")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	// Crop the original image to 300x300px size using the center anchor.
	src = imaging.CropAnchor(src, 300, 300, imaging.Center)

	// Resize the cropped image to width = 200px preserving the aspect ratio.
	src = imaging.Resize(src, 200, 0, imaging.Lanczos)

	// Create a blurred version of the image.
	img1 := imaging.Blur(src, 5)

	// Create a grayscale version of the image with higher contrast and sharpness.
	img2 := imaging.Grayscale(src)
	img2 = imaging.AdjustContrast(img2, 20)
	img2 = imaging.Sharpen(img2, 2)

	// Create an inverted version of the image.
	img3 := imaging.Invert(src)

	// Create an embossed version of the image using a convolution filter.
	img4 := imaging.Convolve3x3(
		src,
		[9]float64{
			-1, -1, 0,
			-1, 1, 1,
			0, 1, 1,
		},
		nil,
	)

	// Create a new image and paste the four produced images into it.
	dst := imaging.New(400, 400, color.NRGBA{})
	dst = imaging.Paste(dst, img1, image.Pt(0, 0))
	dst = imaging.Paste(dst, img2, image.Pt(0, 200))
	dst = imaging.Paste(dst, img3, image.Pt(200, 0))
	dst = imaging.Paste(dst, img4, image.Pt(200, 200))

	// Save the resulting image as JPEG.
	err = imaging.Save(dst, "pic/imaging.jpg")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}

func TestBild(t *testing.T) {
	img, err := imgio.Open("pic/img.png")
	if err != nil {
		fmt.Println(err)
		return
	}

	//inverted := effect.Invert(img)
	resized := transform.Resize(img, 800, 800, transform.Lanczos)
	rotated := transform.Rotate(resized, 45, nil)

	if err := imgio.Save("pic/bild.jpg", rotated, imgio.JPEGEncoder(100)); err != nil {
		fmt.Println(err)
		return
	}
}
