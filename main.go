package main

import (
	"fmt"
	"image/color"
	"image/png"
	"os"
)

func main() {
	infile, err := os.Open("dots2.png")
	if err != nil {
		fmt.Printf("Oh no %s\n", err)
	}

	defer infile.Close()

	src, err := png.Decode(infile)
	if err != nil {
		// replace this with real error handling
		fmt.Printf("Oh no %s\n", err)
	}

	countByColor := make(map[color.Color]int)

	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			// src.
			color := src.At(x, y)

			// r, g, b,
			_, _, _, a := color.RGBA()
			// if r>>8 == 255 && g>>8 == 255 && b>>8 == 255 {
			if a == 0 {
				continue
			}

			countByColor[color] = countByColor[color] + 1
		}
	}
	filteredCounts, total := filterSimilar(countByColor)

	for color, count := range filteredCounts {
		r, g, b, _ := color.RGBA()
		fmt.Printf("%d %d %d: %d%%\n", r>>8, g>>8, b>>8, (count * 100 / total))
	}
}

func filterSimilar(countByColor map[color.Color]int) (map[color.Color]int, int) {
	retMap := make(map[color.Color]int)
	total := 0
	for color, count := range countByColor {
		if count >= 2000 {
			retMap[color] = count
			total += count
		}
	}

	return retMap, total
}
