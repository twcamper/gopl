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
	"net/url"
	"strconv"
)

func main() {
	http.HandleFunc("/lissajous", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	imgCfg := configure(getForm(r))
	lissajous(w, imgCfg)
}

func lissajous(out io.Writer, cfg *imageConfig) {
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

type imageConfig struct {
	XOscillatorCycles float64
	AngularResolution float64
	CanvasSize        int
	FrameCount        int
	Delay             int
}

func configure(form *url.Values) *imageConfig {
	cfg := imageConfig{
		XOscillatorCycles: 5,
		AngularResolution: 0.001,
		CanvasSize:        100,
		FrameCount:        64,
		Delay:             8}

	for k, v := range *form {
		switch k {
		case "XOscillatorCycles":
			setFloat(v, &cfg.XOscillatorCycles)
		case "AngularResolution":
			setFloat(v, &cfg.AngularResolution)
		case "CanvasSize":
			setInt(v, &cfg.CanvasSize)
		case "FrameCount":
			setInt(v, &cfg.FrameCount)
		case "Delay":
			setInt(v, &cfg.Delay)
		default:
			log.Printf("Incorrect query parameter '%s'", k)
		}
	}
	return &cfg
}

func getForm(r *http.Request) *url.Values {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	return &r.Form
}

func setInt(slice []string, field *int) {
	s := slice[0]
	i, error := strconv.Atoi(s)
	if error == nil {
		*field = i
	} else {
		log.Printf("setInt(): %v\n", error)
	}
}

func setFloat(slice []string, field *float64) {
	s := slice[0]
	f, error := strconv.ParseFloat(s, 64)
	if error == nil {
		*field = f
	} else {
		log.Printf("setFloat(): %v\n", error)
	}
}
