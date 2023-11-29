package main

func FindOddNumber(input []int) int {
	if len(input) == 1 {
		return input[0]
	}

	odds := 0
	for _, v := range input {
		odds = odds ^ v
	}
	return odds
}
