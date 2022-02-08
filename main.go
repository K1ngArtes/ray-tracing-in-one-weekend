package main

import (
	"fmt"
	"log"
	"os"

	"github.com/K1ngArtes/ray-tracing-in-one-weekend/trace"
)

const (
	imageWidth  = 256
	imageHeight = 256
)

func main() {
	l := log.New(os.Stderr, "", 0)

	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	f := os.Stdout
	defer f.Close()

	// Picture is read row by row
	for row := imageHeight - 1; row >= 0; row-- {
		l.Printf("\rScanlines remaining: %d", row)
		for col := 0; col < imageWidth; col++ {
			color := trace.Color{
				float64(col) / (imageWidth - 1),
				float64(row) / (imageHeight - 1),
				0.25,
			}

			writeColor(f, color)
		}
	}
	l.Print("Done!")
}

func writeColor(out *os.File, color trace.Color) {
	ir := int(255.999 * color.R())
	ig := int(255.999 * color.G())
	ib := int(255.999 * color.B())

	s := fmt.Sprintf("%d %d %d\n", ir, ig, ib)

	if _, err := out.WriteString(s); err != nil {
		panic(err)
	}
}
