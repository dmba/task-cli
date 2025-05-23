package utils

type FilterFunc[T any] func(T) bool

func Filter[T any](s []T, predicate FilterFunc[T]) []T {
	result := make([]T, 0, len(s))
	for _, v := range s {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}
