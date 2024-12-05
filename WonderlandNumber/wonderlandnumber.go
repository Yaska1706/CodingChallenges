package WonderlandNumber

const (
	MIN_SIZE_NUMBER = 100000
	MAX_SIZE_NUMBER = 1000000
)

func generateWonderlandNumber(size int) int {
	var result int
	if size < MIN_SIZE_NUMBER && size >= MAX_SIZE_NUMBER {
		return 0
	}
	for i := MIN_SIZE_NUMBER; i <= size; i++ {
		if wonderlandNumber(i) {
			result = i
		}
	}
	return result
}

func wonderlandNumber(number int) bool {
	var (
		b            bool
		results      int
		storeresults []int
	)

	store := storePresentDigits(number)
	for i := 1; i <= 6; i++ {
		results = number * i
		if checkNumbersExist(results, store) {
			storeresults = append(storeresults, results)
		}
	}

	if len(storeresults) == len(store) {
		return true
	} else if len(storeresults) < len(store) {
		return false
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

func storePresentDigits(number int) []int {
	n := countDigits(number)
	if n != 6 {
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
	digits := storePresentDigits(number)
	for i := 1; i <= len(digits); i++ {
		for j := 1; j <= len(store); j++ {
			if store[j-1] == digits[i-1] {
				results = append(results, store[j-1])
			}
		}
	}
	if len(results) == len(store) {
		return true
	} else if results == nil {
		return false
	}

	return b
}
