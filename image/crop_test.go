package image

import (
	"github.com/anthonynsimon/bild/transform"
	"github.com/disintegration/gift"
	"github.com/disintegration/imaging"
	"image"
	"testing"
)

func Benchmark_Imaging_Crop(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = imaging.Crop(DefaultImage, image.Rect(70, 70, 210, 210))
	}
}

func Benchmark_Bild_Crop(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = transform.Crop(DefaultImage, image.Rect(70, 70, 210, 210))
	}
}

func Benchmark_Gift_Crop(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		g := gift.New(gift.Crop(image.Rect(70, 70, 210, 210)))
		dst := image.NewNRGBA(g.Bounds(DefaultImage.Bounds()))
		g.Draw(dst, DefaultImage)
	}
}
