package sprint

func IntVsFloat(i int, f float32) string {
	if float32(i) > f {
		return "Integer"
	} else if f > float32(i) {
		return "Float"
	} else {
		return "Same"
	}
}
