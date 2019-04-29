package ignore_test

import (
	"github.com/akyoto/ignore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIgnoreParentheses(t *testing.T) {
	test(t, "a () b", 4)
	test(t, "a (ignored) b", 4)
	test(t, "a (ignored () ignored) b", 4)

	test(t, "a [] b", 4)
	test(t, "a [ignored] b", 4)
	test(t, "a [ignored [] ignored] b", 4)

	test(t, "a {} b", 4)
	test(t, "a {ignored} b", 4)
	test(t, "a {ignored [} ignored} b", 4)
}

func TestIgnoreStrings(t *testing.T) {
	test(t, "a \"ignored\" b", 4)
	test(t, "a \"ignored(ignored\" b", 4)
	test(t, "a \"ignored[(]ignored\" b", 4)
}

func TestIgnoreEscapeSequences(t *testing.T) {
	test(t, "a \\n\\n b", 4+2)
}

func TestIgnoreCharacterString(t *testing.T) {
	test(t, "a 'c' b", 4)
	test(t, "a 'cde' b", 4)
}

func test(t *testing.T, s string, expectedCount int) {
	reader := ignore.Reader{}
	count := 0

	for _, r := range s {
		if reader.CanIgnore(r) {
			continue
		}

		count++
	}

	assert.Equal(t, expectedCount, count)
}
