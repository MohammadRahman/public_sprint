package sprint

func Season(monthName string) string {

	monthToSeason := map[string]string{
		"jan": "winter", "feb": "winter", "dec": "winter",
		"mar": "spring", "apr": "spring", "may": "spring",
		"jun": "summer", "jul": "summer", "aug": "summer",
		"sep": "autumn", "oct": "autumn", "nov": "autumn",
	}

	// Check if the normalizedMonth exists in the map
	if season, ok := monthToSeason[monthName]; ok {
		return season
	}

	// If the month is not found, return invalid input with the month name
	return "invalid input: " + monthName
}
