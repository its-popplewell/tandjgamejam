package main

func FindFunc[T any](slice []T, match func(T) bool) int {
	for i, v := range slice {
		if match(v) {
			return i
		}
	}
	return -1
}

func DeleteFunc[T any](slice *[]T, match func(T) bool) {
	i := FindFunc(*slice, match)
	if i == -1 {
		return
	}
	*slice = append((*slice)[:i], (*slice)[i+1:]...)
}
