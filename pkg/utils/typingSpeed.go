package utils

import (
	"strings"
	"time"
)

func CalculateTypingSpeed(sentence string, duration time.Duration) float64 {
	words := strings.Fields(sentence)
	wordCount := len(words)
	seconds := duration.Seconds()
	return float64(wordCount) / (seconds / 60.0)
}
