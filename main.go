package main

import "fmt"

const (
	imageWidth = 256
	imageHeight = 256
)

func main() {
	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	// Picture is read row by row
	for row := imageHeight-1; row >= 0; row-- {
		for col := 0; col < imageWidth; col++ {
			r := float64(col) / (imageWidth-1)
			g := float64(row) / (imageHeight-1)
			b := 0.25
			
			ir := int(255.999 * r)
			ig := int(255.999 * g)
			ib := int(255.999 * b)

			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}