package image

import (
	"github.com/anthonynsimon/bild/transform"
	"github.com/disintegration/gift"
	"github.com/disintegration/imaging"
	log "github.com/sirupsen/logrus"
	"image"
	"image/color"
	"testing"
)

func Benchmark_Imaging_Rotate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = imaging.Rotate(DefaultImage, 30, color.Black)
	}
}

func Test_Imaging_Rotate(b *testing.T) {
	dst := imaging.Rotate(DefaultImage, 30, color.Transparent)
	// Save the resulting image as JPEG.
	err := imaging.Save(dst, "pic/Imaging_Rotate.jpg")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}

func Benchmark_Bild_Rotate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = transform.Rotate(DefaultImage, 45, &transform.RotationOptions{ResizeBounds: true})
	}
}

func Benchmark_Gift_Rotate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		g := gift.New(gift.Rotate(30, color.Transparent, gift.CubicInterpolation))
		dst := image.NewNRGBA(g.Bounds(DefaultImage.Bounds()))
		g.Draw(dst, DefaultImage)
	}
}
