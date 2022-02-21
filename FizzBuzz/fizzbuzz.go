package FizzBuzz

import "fmt"

func fizzBuzz(number int) []string {
	if number < 0 {
		return []string{}
	}
	results := make([]string, number)
	for i := 1; i <= number; i++ {
		if i%3 == 0 && i%5 == 0 {
			results[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			results[i-1] = "Fizz"
		} else if i%5 == 0 {
			results[i-1] = "Buzz"
		} else {
			results[i-1] = fmt.Sprintln(i)
		}
	}
	return results
}
