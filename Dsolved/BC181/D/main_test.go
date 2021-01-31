package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblemD(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		assert.Equal(t, true, ProblemD("1234"))
	})
	t.Run("Example 2", func(t *testing.T) {
		assert.Equal(t, false, ProblemD("1333"))
	})
	t.Run("Example 3", func(t *testing.T) {
		assert.Equal(t, true, ProblemD("8"))
	})
	t.Run("Example 4", func(t *testing.T) {
		assert.Equal(t, true, ProblemD("16"))
	})
	t.Run("Example 5", func(t *testing.T) {
		assert.Equal(t, true, ProblemD("61"))
	})

}
