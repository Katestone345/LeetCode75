package challenges_test

import (
	"github.com/leetcode75/challenges"
	"reflect"
	"testing"
)

func Test_KidsWithCandies(t *testing.T) {

	var candies1 []int
	var expected1 []bool

	expected1 = append(expected1, true, true, true, false, true)
	candies1 = append(candies1, 2, 3, 5, 1, 3)
	result := challenges.KidsWithCandies(candies1, 3)

	if reflect.DeepEqual(result, expected1) == false {
		t.Fatalf(`result = %v; want [true,true,true,false,true]`, result)
	}

	var candies2 []int
	var expected2 []bool

	expected2 = append(expected2, true, false, false, false, false)
	candies2 = append(candies2, 4, 2, 1, 1, 2)
	result = challenges.KidsWithCandies(candies2, 1)

	if reflect.DeepEqual(result, expected2) == false {
		t.Fatalf(`result = %v; want [true,false,false,false,false]`, result)
	}

	var candies3 []int
	var expected3 []bool

	expected3 = append(expected3, true, false, true)
	candies3 = append(candies3, 12, 1, 12)
	result = challenges.KidsWithCandies(candies3, 10)

	if reflect.DeepEqual(result, expected3) == false {
		t.Fatalf(`result = %v; want [true, false, true]`, result)
	}
}
