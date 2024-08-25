package main

func getLast[T any](s []T) T {
	var zeroV T
	if len(s) > 0 {
		return s[len(s)-1]
	}
	return zeroV
}
