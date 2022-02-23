package WonderlandNumber

const (
	MIN_SIZE_NUMBER = 100000
	MAX_SIZE_NUMBER = 1000000
)

func generatewonderlandnumber(size int) int {
	var result int
	if size < MIN_SIZE_NUMBER && size >= MAX_SIZE_NUMBER {
		return 0
	}
	for i := MIN_SIZE_NUMBER; i <= size; i++ {
		check := wonderlandnumber(i)
		if check {
			result = i
		}
	}
	return result
}

func wonderlandnumber(number int) bool {
	var b bool
	var results int
	var storeresults []int
	store := storepresentDigits(number)
	for i := 1; i <= 6; i++ {
		results = number * i
		check := checkNumbersExist(results, store)
		if check {
			storeresults = append(storeresults, results)
		}
	}

	if len(storeresults) == len(store) {
		b = true
	} else if len(storeresults) < len(store) {
		b = false
	}

	return b
}

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
	results := make([]int, 0)
	digits := storepresentDigits(number)
	for i := 1; i <= len(digits); i++ {
		for j := 1; j <= len(store); j++ {
			if store[j-1] == digits[i-1] {
				results = append(results, store[j-1])
			}
		}
	}
	if len(results) == len(store) {
		b = true
	} else if results == nil {
		b = false
	}

	return b
}
