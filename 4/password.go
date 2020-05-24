package password

func validPassword(password int, checkers ...func([]int) bool) bool {
	numSlice := intToSlice(password)
	for _,check := range checkers {
		if !check(numSlice) {
			return false
		}
	}
	return true
}

func checkLength(password []int) bool {
	return len(password) == 6
}

func checkNotDecrementing(password []int) bool {
	min := password[0]
	for _, digit := range password {
		if digit < min {
			return false
		}
		min = digit
	} 
	return true
}

func checkAdjacentPair(password []int) bool {
	var lastDigit int
	for _, digit := range password {
		if digit == lastDigit {
			return true
		}
		lastDigit = digit
	}
	return false
}

func checkHasDouble(password []int) bool {
	var counts []int
	lastDigit := password[0]
	counter := 1

	for _, digit := range password[1:] {
		
		if digit == lastDigit {
			counter++
		} else {
			counts = append(counts, counter)
			counter = 1
		}
		lastDigit = digit
	}
	counts = append(counts, counter)

	for _, count := range counts {
		if count == 2 {
			return true
		}
	}

	return false
}

// Part 1 checks 
func getKeys(min int, max int) []int {
	keys := []int{}
	
	for i := min; i <= max; i++ {
		if validPassword(i, checkLength, checkAdjacentPair, checkNotDecrementing) {
			keys = append(keys, i)
		}
	}
	return keys
}

// Part 2 checks 
func getKeys2(min int, max int) []int {
	keys := []int{}
	
	for i := min; i <= max; i++ {
		if validPassword(i, checkLength, checkAdjacentPair, checkNotDecrementing, checkHasDouble) {
			keys = append(keys, i)
		}
	}
	return keys
}


func intToSlice(num int) []int {
	res := []int{}
	for num / 10 > 0 {
		res = append([]int{num % 10}, res...)
		num /= 10
	}
	res = append([]int{num % 10}, res...)
	return res
}