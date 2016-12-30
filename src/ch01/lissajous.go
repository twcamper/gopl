package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var backgroundColor = color.Black
var foregroundColor = color.RGBA{0x00, 0xa0, 0x00, 0xff}
var palette = []color.Color{backgroundColor, foregroundColor}

const (
	backgroundIndex = 0
	foregroundIndex = 1
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles     = 5     // number of complete x oscillator revolutions
		res        = 0.001 // angular resolution
		size       = 350   // image canvas covers [-size..+size]
		frameCount = 64    // animation frames
		delay      = 8     // delay between frames in 10ms units
	)

	yFrequency := rand.Float64() * 3.0 // relative frequency of y oscillator
	animation := gif.GIF{LoopCount: frameCount}
	phaseDifference := 0.0

	for i := 0; i < frameCount; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*yFrequency + phaseDifference)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), foregroundIndex)
		}

		phaseDifference += 0.1
		animation.Delay = append(animation.Delay, delay)
		animation.Image = append(animation.Image, img)
	}
	gif.EncodeAll(out, &animation) // NOTE: ignoring encoding errors
}
