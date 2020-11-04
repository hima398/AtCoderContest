package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblemD(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		assert.Equal(t, "Yes", ProblemD("1234"))
	})
	t.Run("Example 2", func(t *testing.T) {
		assert.Equal(t, "No", ProblemD("1333"))
	})
	t.Run("Example 3", func(t *testing.T) {
		assert.Equal(t, "Yes", ProblemD("8"))
	})

}

func TestDeleteS(t *testing.T) {
	t.Run("", func(t *testing.T) {
		assert.Equal(t, "ac", DeleteS("abc", 1))
	})
}

func TestPremuteS(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		fmt.Println(PermuteS("1"))
	})
	t.Run("2", func(t *testing.T) {
		fmt.Println(PermuteS("12"))
	})
	t.Run("3", func(t *testing.T) {
		fmt.Println(PermuteS("123"))
	})
	t.Run("4", func(t *testing.T) {
		fmt.Println(PermuteS("1234"))
	})
}
