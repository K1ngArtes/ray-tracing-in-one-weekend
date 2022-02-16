package main

import (
	"fmt"
	"log"
	"os"

	"github.com/K1ngArtes/ray-tracing-in-one-weekend/geom"
	"github.com/K1ngArtes/ray-tracing-in-one-weekend/trace"
)

// const (
// 	imageWidth  = 256
// 	imageHeight = 256
// )

func main() {

	// Image
	aspectRatio := 16.0 / 9.0
	imageWidth := 1000
	imageHeight := int(float64(imageWidth) / aspectRatio)

	// Camera
	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	origin := geom.Vec3{0, 0, 0}
	horizontal := geom.Vec3{viewportWidth, 0, 0}
	vertical := geom.Vec3{0, viewportHeight, 0}
	// origin - horizontal/2 - vertical/2 - vec3(0, 0, focal_length);
	lowerLeftCorner := origin.Minus(horizontal.Div(2)).Minus(vertical.Div(2)).Minus(geom.Vec3{0, 0, focalLength})

	l := log.New(os.Stderr, "", 0)

	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	f := os.Stdout
	defer f.Close()

	// Picture is read row by row
	for row := imageHeight - 1; row >= 0; row-- {
		// l.Printf("\rScanlines remaining: %d", row)
		for col := 0; col < imageWidth; col++ {
			u := float64(col) / (float64(imageWidth) - 1)
			v := float64(row) / (float64(imageHeight) - 1)

			// ray r(origin, lower_left_corner + u*horizontal + v*vertical - origin);
			ray := trace.Ray{origin, lowerLeftCorner.Plus(horizontal.Scaled(u)).Plus(vertical.Scaled(v)).Minus(origin)}

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
	sphere := trace.NewSphere(geom.Vec3{0, 0, -1}, 0.5)
	var hit trace.Hit
	isHit := sphere.Hit(r, 0, 100, &hit)

	if isHit {
		normal := hit.Normal
		return trace.Color{normal.X()+1, normal.Y()+1, normal.Z()+1}.Scaled(0.5)
	}
	unitDirection := geom.Vec3(r.Dir).Unit()

	hit.T = 0.5 * unitDirection.Y() + 0.5;

	blue := trace.Color{0.5, 0.7, 1.0}

	// linear interpolation
	// blendedValue=(1−𝑡)⋅startValue+𝑡⋅endValue
	// (1.0-t)*color(1.0, 1.0, 1.0) + t*color(0.5, 0.7, 1.0);
	return trace.White.Scaled(1.0 - hit.T).Plus(blue.Scaled(hit.T))
}