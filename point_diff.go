package sprint

type Point struct {
	X    float32
	Y    float32
	Text string
}

func PointDiff(p1, p2 Point) Point {
	dist1 := p1.X*p1.X + p1.Y*p1.Y
	dist2 := p2.X*p2.X + p2.Y*p2.Y

	if dist1 > dist2 {
		return p1
	} else if dist2 > dist1 {
		return p2
	} else {
		return p1
	}
}
