package arraytools

// EqualInt - Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func EqualInt(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// StartsWithInt - Checks if an array starts equal to another array
func StartsWithInt(a, starts []int) bool {
	if len(a) < len(starts) {
		return false
	}
	return EqualInt(a[0:len(starts)], starts)
}

// ReverseInt - Reverse the int
func ReverseInt(arr []int) []int {
	var result []int
	for i := len(arr) - 1; i >= 0; i-- {
		result = append(result, arr[i])
	}
	return result
}

func LastElementInt(arr []int) int {
	return arr[len(arr)-1]
}

func ElementInListInt(element int, list []int) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == element {
			return true
		}
	}
	return false
}

func FindIndexOfElementInt(element int, list []int) int {
	for i := 0; i < len(list); i++ {
		if list[i] == element {
			return i
		}
	}
	return -1
}
