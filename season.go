package sprint

func Season(monthName string) string {

	switch monthName {
	case "dec", "jan", "feb":
		return "winter"
	case "mar", "apr", "may":
		return "spring"
	case "jun", "jul", "aug":
		return "summer"
	case "sep", "oct", "nov":
		return "autumn"
	default:
		return "Invalid input: " + monthName
	}
}
