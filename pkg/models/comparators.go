package models

import "cmp"

func ByCompleteness(a, b Task) int {
	return cmp.Or(
		ByStatus(b, a),
		ByCreatedAt(b, a),
	)
}

func ByStatus(a, b Task) int {
	order := map[Status]int{
		ToDo:       0,
		InProgress: 1,
		Done:       2,
	}

	return order[b.Status] - order[a.Status]
}

func ByCreatedAt(a, b Task) int {
	if a.CreatedAt.Before(b.CreatedAt) {
		return -1
	}
	if a.CreatedAt.After(b.CreatedAt) {
		return 1
	}
	return 0
}
