//Server3 responds to requests with lissajous
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
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {

	cycles, res, size, nframes, delay := 8, 0.01, 100, 64, 8

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	if val, err := strconv.Atoi(r.FormValue("cycles")); err == nil {
		cycles = val
	}
	if val, err := strconv.ParseFloat(r.FormValue("res"), 32); err == nil {
		res = val
	}
	if val, err := strconv.Atoi(r.FormValue("size")); err == nil {
		size = val
	}
	if val, err := strconv.Atoi(r.FormValue("nframes")); err == nil {
		nframes = val
	}
	if val, err := strconv.Atoi(r.FormValue("delay")); err == nil {
		delay = val
	}

	lissajous(w, cycles, res, size, nframes, delay)

}

var palette = []color.Color{color.Black, color.RGBA{0, 255, 0, 1}, color.RGBA{255, 0, 0, 1}, color.RGBA{0, 0, 255, 1}}

func lissajous(out io.Writer, cycles int, res float64, size int, nframes int, delay int) {

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles*2)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), uint8(i%4))
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
