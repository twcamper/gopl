package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/lissajous", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	imageConfig := configure()
	lissajous(w, imageConfig)
}

func lissajous(out io.Writer, cfg *ImageConfig) {

	const (
		backgroundIndex = 0
		foregroundIndex = 1
	)

	size := cfg.CanvasSize
	backgroundColor := color.Black
	foregroundColor := color.RGBA{0x00, 0xa0, 0x00, 0xff}
	palette := []color.Color{backgroundColor, foregroundColor}
	yFrequency := rand.Float64() * 3.0 // relative frequency of y oscillator
	animation := gif.GIF{LoopCount: cfg.FrameCount}
	phaseDifference := 0.0

	for i := 0; i < cfg.FrameCount; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cfg.XOscillatorCycles*2*math.Pi; t += cfg.AngularResolution {
			x := calculateColorIndexCoordinate(math.Sin(t), size)
			y := calculateColorIndexCoordinate(math.Sin(t*yFrequency+phaseDifference), size)
			img.SetColorIndex(x, y, foregroundIndex)
		}

		phaseDifference += 0.1
		animation.Delay = append(animation.Delay, cfg.Delay)
		animation.Image = append(animation.Image, img)
	}
	gif.EncodeAll(out, &animation) // NOTE: ignoring encoding errors
}

func calculateColorIndexCoordinate(rawCoordinate float64, size int) int {
	offset := rawCoordinate*float64(size) + 0.5
	return size + int(offset)
}

type ImageConfig struct {
	XOscillatorCycles float64
	AngularResolution float64
	CanvasSize        int
	FrameCount        int
	Delay             int
}

func configure() *ImageConfig {
	return &ImageConfig{
		XOscillatorCycles: 5,
		AngularResolution: 0.001,
		CanvasSize:        350,
		FrameCount:        64,
		Delay:             8}
}
