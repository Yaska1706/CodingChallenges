package FizzBuzz

import "fmt"

func fizzBuzz(number int) []string {
	if number < 0 {
		return []string{}
	}
	results := make([]string, 0)
	for i := 1; i <= number; i++ {
		if i%3 == 0 && i%5 == 0 {
			results = append(results, "FizzBuzz")
		} else if i%3 == 0 {
			results = append(results, "Fizz")
		} else if i%5 == 0 {
			results = append(results, "Buzz")
		} else {
			results = append(results, fmt.Sprintln(i))
		}
	}
	return results
}
