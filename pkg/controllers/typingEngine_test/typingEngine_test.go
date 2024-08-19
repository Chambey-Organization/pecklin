package typingEngine_test

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/stretchr/testify/assert"
	"main.go/typingEngine"
	"testing"
)

func TestCompareAndHighlightInput(t *testing.T) {
	input := "Hello"
	prompt := "Helzo"

	correctStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	incorrectStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("1"))

	highlighted, accuracy := typingEngine.CompareAndHighlightInput(input, prompt)

	expectedHighlighted := fmt.Sprintf(correctStyle.Render("Hel") + incorrectStyle.Render("l") + correctStyle.Render("o"))
	expectedAccuracy := 80.0

	assert.Equal(t, expectedHighlighted, highlighted)
	assert.Equal(t, expectedAccuracy, accuracy)
}
