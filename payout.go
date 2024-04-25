package sprint

func Payout(amount int, denominations []int) []int {
	for i := 0; i < len(denominations); i++ {
		for j := i + 1; j < len(denominations); j++ {
			if denominations[i] < denominations[j] {
				denominations[i], denominations[j] = denominations[j], denominations[i]
			}
		}
	}
	result := []int{}
	for _, denom := range denominations {
		for amount >= denom {
			result = append(result, denom)
			amount -= denom
		}
	}
	if amount == 0 {
		return result
	}
	return []int{}
}
