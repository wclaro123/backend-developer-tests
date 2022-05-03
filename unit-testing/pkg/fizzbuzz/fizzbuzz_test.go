package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("should return empty array if total is 0", func(t *testing.T) {
		expected := make([]string, 0)
		result := FizzBuzz(0, 2, 3)
		assert.EqualValues(t, expected, result)
	})

	t.Run("should return only numbers if don't are multiple neither Fizz or Buzz values", func(t *testing.T) {
		expected := []string{"1"}
		result := FizzBuzz(1, 2, 3)
		assert.EqualValues(t, expected, result)
	})

	t.Run("should return Fizz if the number is multiple of Fizz value", func(t *testing.T) {
		expected := []string{"1", "Fizz"}
		result := FizzBuzz(2, 2, 3)
		assert.EqualValues(t, expected, result)
	})

	t.Run("should return Buzz if the number is multiple of Buzz value", func(t *testing.T) {
		expected := []string{"1", "Fizz", "Buzz"}
		result := FizzBuzz(3, 2, 3)
		assert.EqualValues(t, expected, result)
	})

}
