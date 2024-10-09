package challenges_test

import (
	"github.com/leetcode75/challenges"
	"testing"
)

func Test_CanPlaceFlowers(t *testing.T) {
	var flowerbed []int

	flowerbed = append(flowerbed, 1, 0, 0, 0, 1)
	result := challenges.CanPlaceFlowers(flowerbed, 1)

	if result != true {
		t.Fatalf(`result = %v; want %v`, result, true)
	}

	var flowerbed2 []int

	flowerbed2 = append(flowerbed2, 1, 0, 0, 0, 0, 1)
	result = challenges.CanPlaceFlowers(flowerbed2, 2)

	if result != false {
		t.Fatalf(`result = %v; want %v`, result, false)
	}

	var flowerbed3 []int

	flowerbed3 = append(flowerbed3, 1, 0, 0, 0, 1)
	result = challenges.CanPlaceFlowers(flowerbed3, 2)

	if result != false {
		t.Fatalf(`result = %v; want %v`, result, false)
	}

	var flowerbed4 []int

	flowerbed4 = append(flowerbed4, 0, 0, 1, 0, 1)
	result = challenges.CanPlaceFlowers(flowerbed4, 1)

	if result != true {
		t.Fatalf(`result = %v; want %v`, result, true)
	}

	var flowerbed5 []int

	flowerbed5 = append(flowerbed5, 0)
	result = challenges.CanPlaceFlowers(flowerbed5, 1)

	if result != true {
		t.Fatalf(`result = %v; want %v`, result, true)
	}

	var flowerbed6 []int

	flowerbed6 = append(flowerbed6, 0, 1, 0)
	result = challenges.CanPlaceFlowers(flowerbed6, 1)

	if result != false {
		t.Fatalf(`result = %v; want %v`, result, false)
	}
}
