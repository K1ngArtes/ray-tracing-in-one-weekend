package main

import (
	"fmt"
	"log"
	"os"

	"github.com/K1ngArtes/ray-tracing-in-one-weekend/math"
	"github.com/K1ngArtes/ray-tracing-in-one-weekend/trace"
)

// const (
// 	imageWidth  = 256
// 	imageHeight = 256
// )

func main() {

	// Image
	aspectRatio := 16.0 / 9.0
	imageWidth := 400
	imageHeight := int(float64(imageWidth) / aspectRatio)

	// Camera
	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	origin := math.Vec3{0, 0, 0}
	horizontal := math.Vec3{viewportWidth, 0, 0}
	vertical := math.Vec3{0, viewportHeight, 0}
	lowerLeftCorner := origin.Minus(horizontal.Div(2)).Minus(vertical.Div(2)).Minus(math.Vec3{0, 0, focalLength})

	l := log.New(os.Stderr, "", 0)

	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	f := os.Stdout
	defer f.Close()

	// Picture is read row by row
	for row := imageHeight - 1; row >= 0; row-- {
		l.Printf("\rScanlines remaining: %d", row)
		for col := 0; col < imageWidth; col++ {
			u := float64(col) / (float64(imageWidth) - 1)
			v := float64(row) / (float64(imageHeight) - 1)

			ray := trace.Ray{origin, lowerLeftCorner.Plus(horizontal.Scaled(u)).Plus(vertical.Scaled(v)).Minus(origin)}

			// color := trace.Color{
			// 	float64(col) / (imageWidth - 1),
			// 	float64(row) / (imageHeight - 1),
			// 	0.25,
			// }

			color := rayColor(ray)

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

func rayColor(r trace.Ray) trace.Color {
	unitDirection := math.Vec3(r.Dir).Unit()

	t := 0.5 * (unitDirection.Y() + 1.0)

	return trace.White.Scaled(1.0 - t).Plus(trace.Color{0.5, 0.7, 1.0}.Scaled(t))
}
