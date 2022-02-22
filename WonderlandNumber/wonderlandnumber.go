package WonderlandNumber

func countDigits(number int) int {
	if number < 10 {
		return 1
	} else {
		return 1 + countDigits(number/10)
	}
}

func storepresentDigits(number int) []int {
	n := countDigits(number)
	if n < 6 || n > 6 {
		return []int{}
	}
	store := make([]int, 6)
	for i := 1; i <= n; i++ {
		digit := number % 10
		store[i-1] = digit
		number = number / 10
	}
	return store
}

func checkNumbersExist(number int, store []int) bool {
	var b bool

	return b
}
