package main

func Josephus(number, start int) int {
	if number == 1 {
		return 1
	}
	return (Josephus(number-1, start)+start-1)%number + 1
}
