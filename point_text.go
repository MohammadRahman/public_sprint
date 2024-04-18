package sprint

import "fmt"

type Point struct {
	X    float32
	Y    float32
	Text string
}

func trimFloat(f float32) string {
	intPart := int(f)
	fracPart := f - float32(intPart)

	intPartStr := fmt.Sprintf("%d", intPart)

	fracPartStr := ""
	for i := 0; i < 6; i++ {
		fracPart *= 10
		digit := int(fracPart) % 10
		fracPartStr += fmt.Sprintf("%d", digit)
		if int(fracPart) == 0 {
			break
		}
	}

	result := intPartStr
	if len(fracPartStr) > 0 {
		result += "." + fracPartStr
	}

	return result
}

func PointText(p Point) Point {
	xFormatted := trimFloat(p.X)
	yFormatted := trimFloat(p.Y)

	formattedText := fmt.Sprintf("Text at (%s, %s)", xFormatted, yFormatted)

	return Point{
		X:    p.X,
		Y:    p.Y,
		Text: formattedText,
	}
}
