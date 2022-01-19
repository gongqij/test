package image

import (
	"github.com/anthonynsimon/bild/transform"
	"github.com/disintegration/gift"
	"github.com/disintegration/imaging"
	"gopkg.in/gographics/imagick.v2/imagick"
	"image"
	"testing"
)

func Benchmark_Imaging_Resize(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = imaging.Resize(DefaultImage, 800, 800, imaging.Lanczos)
	}
}

func Benchmark_Bild_Resize(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = transform.Resize(DefaultImage, 800, 800, transform.Lanczos)
	}
}

func Benchmark_Gift_Resize(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		g := gift.New(gift.Resize(800, 800, gift.LanczosResampling))
		dst := image.NewNRGBA(g.Bounds(DefaultImage.Bounds()))
		g.Draw(dst, DefaultImage)
	}
}

func Benchmark_imagick_Resize(b *testing.B) {
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	_ = mw.ReadImage("pic/img.png")
	// Get original logo size
	width := mw.GetImageWidth()
	height := mw.GetImageHeight()

	// Calculate half the size
	hWidth := uint(width / 2)
	hHeight := uint(height / 2)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		imagick.Initialize()
		// Resize the image using the Lanczos filter
		// The blur factor is a float, where > 1 is blurry, < 1 is sharp
		_ = mw.ResizeImage(hWidth, hHeight, imagick.FILTER_LANCZOS, 1)
	}
}
