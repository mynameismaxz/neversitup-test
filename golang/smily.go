package main

import "regexp"

func CountSmilyFace(text []string) int {
	if len(text) == 0 {
		return 0
	}
	pattern := "[;:][~-]?[)D]"
	count := 0
	for _, v := range text {
		if match, _ := regexp.MatchString(pattern, v); match {
			count++
		}
	}
	return count
}
