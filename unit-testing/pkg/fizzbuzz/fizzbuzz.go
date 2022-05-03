package fizzbuzz

import "strconv"

// FizzBuzz performs a FizzBuzz operation over a range of integers
//
// Given a range of integers:
// - Return "Fizz" if the integer is divisible by the `fizzAt` value.
// - Return "Buzz" if the integer is divisible by the `buzzAt` value.
// - Return "FizzBuzz" if the integer is divisible by both the `fizzAt` and
//   `buzzAt` values.
// - Return the original number if is is not divisible by either the `fizzAt` or
//   the `buzzAt` values.
func FizzBuzz(total, fizzAt, buzzAt int64) []string {
	result := []string{}

	for i := int64(1); i <= total; i++ {
		if !(i%fizzAt == 0) && !(i%buzzAt == 0) {
			result = append(result, strconv.FormatInt(i, 10))
			continue
		}

		str := ""
		if i%fizzAt == 0 {
			str = "Fizz"
		}

		if i%buzzAt == 0 {
			str += "Buzz"
		}

		result = append(result, str)
	}

	return result
}
