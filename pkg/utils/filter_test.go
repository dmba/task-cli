package utils

import (
	"reflect"
	"testing"
)

func TestFiltersElementsBasedOnPredicate(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	predicate := func(n int) bool { return n%2 == 0 }
	expected := []int{2, 4}

	result := Filter(input, predicate)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestReturnsEmptySliceWhenInputIsEmpty(t *testing.T) {
	input := []int{}
	predicate := func(n int) bool { return n%2 == 0 }
	expected := []int{}

	result := Filter(input, predicate)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestReturnsEmptySliceWhenNoElementsMatchPredicate(t *testing.T) {
	input := []int{1, 3, 5}
	predicate := func(n int) bool { return n%2 == 0 }
	expected := []int{}

	result := Filter(input, predicate)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestHandlesSliceWithSingleElementMatchingPredicate(t *testing.T) {
	input := []int{2}
	predicate := func(n int) bool { return n%2 == 0 }
	expected := []int{2}

	result := Filter(input, predicate)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestHandlesSliceWithSingleElementNotMatchingPredicate(t *testing.T) {
	input := []int{1}
	predicate := func(n int) bool { return n%2 == 0 }
	expected := []int{}

	result := Filter(input, predicate)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
