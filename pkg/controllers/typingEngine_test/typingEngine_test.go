package typingEngine_test

import (
	"github.com/stretchr/testify/assert"
	"main.go/typingEngine"
	"testing"
)

func TestCompareAndHighlightInput(t *testing.T) {
	input := "Hello"
	prompt := "Helzo"

	highlighted, accuracy := typingEngine.CompareAndHighlightInput(input, prompt)

	expectedHighlighted := "<style for correct>H</style><style for correct>e</style><style for correct>l</style><style for incorrect>l</style><style for incorrect>o</style>"
	expectedAccuracy := 60.0

	assert.Equal(t, expectedHighlighted, highlighted)
	assert.Equal(t, expectedAccuracy, accuracy)
}
