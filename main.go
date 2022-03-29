package main

import (
	"log"
	"math"
)

func main() {
	c := 1130.0

	h := 12.0
	w := 23.0
	l := 17.5

	f := c / 2 * math.Sqrt(math.Pow(1, 2)/math.Pow(l, 2)+math.Pow(0, 2)/math.Pow(w, 2)+math.Pow(0, 2)/math.Pow(h, 2))
	log.Println(f)
}
