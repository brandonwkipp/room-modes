package main

import (
	"log"
	"math"
)

type Mode struct {
	Dimension  string
	Frequency  float64
	Type       string
	Wavelength float64
}

// Speed of Sound (ft/s)
var c = 1125.33

func CalculateModeFrequency(d1, d2, d3, p, q, r float64) float64 {
	return c / 2 * math.Sqrt(math.Pow(p/d1, 2)+math.Pow(q/d2, 2)+math.Pow(r/d3, 2))
}

func CalculateWavelength(f float64) float64 {
	return c / f
}

func main() {
	modeLimit := 9

	axial := make([]Mode, (modeLimit * 3))

	h := 12.0
	w := 23.0
	l := 17.5

	for i := 0; i < modeLimit; i++ {
		fL := CalculateModeFrequency(l, w, h, float64(i+1), 0, 0)
		axial[i*3] = Mode{Frequency: fL, Dimension: "L", Type: "Axial", Wavelength: c / fL}

		fW := CalculateModeFrequency(l, w, h, 0, float64(i+1), 0)
		axial[i*3+1] = Mode{Frequency: fW, Dimension: "W", Type: "Axial", Wavelength: c / fW}

		fH := CalculateModeFrequency(l, w, h, 0, 0, float64(i+1))
		axial[i*3+2] = Mode{Frequency: fH, Dimension: "H", Type: "Axial", Wavelength: c / fH}
	}

	for i := range axial {
		log.Println(axial[i])
	}
}
