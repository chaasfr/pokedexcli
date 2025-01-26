package main

import "strings"

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	trimLowerText := strings.Fields(lowerText)
	return trimLowerText
}