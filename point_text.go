package sprint

import "fmt"

// Define the Point structure
type Point struct {
	X    float32
	Y    float32
	Text string
}

func PointText(p Point) Point {
	formattedText := fmt.Sprintf("%s at (%.2f, %.2f)", p.Text, p.X, p.Y)

	return Point{
		X:    p.X,
		Y:    p.Y,
		Text: formattedText,
	}
}
