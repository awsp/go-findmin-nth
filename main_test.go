package main

import (
	"testing"
)

func assert(t *testing.T, expected int, actual int) {
	if actual != expected {
		t.Errorf("%s() failed. Expecting: %d but got %d", t.Name(), expected, actual)
	}
}

func TestMin(t *testing.T) {
	var input = []int{
		40, 42, 33, 45,
	}

	ans := Min(input, 2)
	expected := 40

	assert(t, expected, ans)
}

func TestMinMNotExists(t *testing.T) {
	var input = []int{
		3, 2, 9, 16, 11, 45,
	}
	ans := Min(input, 10)
	expected := -1

	assert(t, expected, ans)
}

func TestMinMNegative(t *testing.T) {
	var input = []int{
		3, 2, 9, 16, 11, 45,
	}
	ans := Min(input, -1)
	expected := -1

	assert(t, expected, ans)
}

func TestFindBoundary(t *testing.T) {
	var input = []int{
		3, 2, 9, 16, 11, 45,
	}

	max, min := FindBoundary(input)
	expectedMax, expectedMin := 45, 2

	assert(t, expectedMax, max)
	assert(t, expectedMin, min)
}
