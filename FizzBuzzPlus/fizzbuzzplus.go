package FizzBuzzPlus

import "fmt"

func fizzbuzzplus(number int) []string {
	if number < 0 {
		return []string{}
	}
	results := make([]string, number)
	for i := 1; i <= number; i++ {
		if i%3 == 0 && checkNumber(i, 3) {
			results[i-1] = "Fizz"
		} else if i%5 == 0 && checkNumber(i, 5) {
			results[i-1] = "Buzz"
		} else {
			results[i-1] = fmt.Sprint(i)
		}
	}
	return results
}

func checkNumber(number, n int) bool {
	var b bool
	for number > 0 {
		digit := number % 10
		number = number / 10
		if digit == n {
			b = true
			break
		}
	}
	return b
}
