package sprint

import "fmt"

// Define the Point structure
type Point struct {
	X    float32
	Y    float32
	Text string
}

func PointText(p Point) Point {
	// Format the X and Y coordinates to two decimal places
	xFormatted := fmt.Sprintf("%.2f", p.X)
	yFormatted := fmt.Sprintf("%.2f", p.Y)

	// Format the Text field with coordinates (X, Y)
	formattedText := fmt.Sprintf("Text at (%s, %s)", xFormatted, yFormatted)

	// Create and return a new Point with the updated Text field
	return Point{
		X:    p.X,
		Y:    p.Y,
		Text: formattedText,
	}
}
