package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
