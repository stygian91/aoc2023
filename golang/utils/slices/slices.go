package slices

import "fmt"

func Any[A any](fn func(int, A) bool, list []A) bool {
	for i, v := range list {
		if fn(i, v) {
			return true
		}
	}

	return false
}

func All[A any](fn func(int, A) bool, list []A) bool {
	for i, v := range list {
		if !fn(i, v) {
			return false
		}
	}

	return true
}

func Chunk[A any](size int, list []A) ([][]A, error) {
	if size < 1 {
		return nil, fmt.Errorf("Size must be a positive int.")
	}

	result := [][]A{}
	chunk := []A{}

	for i, v := range list {
		chunk = append(chunk, v)
		if i%size == 1 || i == len(list)-1 {
			result = append(result, chunk)
			chunk = []A{}
		}
	}

	return result, nil
}

func Map[I any, O any](fn func(int, I) O, list []I) []O {
	result := []O{}

	for key, value := range list {
		result = append(result, fn(key, value))
	}

	return result
}

func Filter[A any](fn func(int, A) bool, list []A) []A {
	result := []A{}

	for key, value := range list {
		if fn(key, value) {
			result = append(result, value)
		}
	}

	return result
}
