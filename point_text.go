package sprint

import "fmt"

// Define the Point structure
type Point struct {
	X    float32
	Y    float32
	Text string
}

func PointText(p Point) Point {
	xFormatted := fmt.Sprintf("%.2f", p.X)
	yFormatted := fmt.Sprintf("%.2f", p.Y)

	formattedText := fmt.Sprintf("%s at (%s, %s)", p.Text, xFormatted, yFormatted)

	return Point{
		X:    p.X,
		Y:    p.Y,
		Text: formattedText,
	}
}
