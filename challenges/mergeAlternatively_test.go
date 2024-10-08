package challenges_test

import (
	"github.com/leetcode75/challenges"
	"testing"
)

func Test_MergeAlternately(t *testing.T) {
	var result string

	result = challenges.MergeAlternately("abc", "pqr")
	if result != "apbqcr" {
		t.Fatalf(`result = %s; want "apbqcr"`, result)
	}

	result = challenges.MergeAlternately("ab", "pqrs")
	if result != "apbqrs" {
		t.Fatalf(`result = %s; want "apbqrs", result`, result)
	}

	result = challenges.MergeAlternately("abcd", "pq")
	if result != "apbqcd" {
		t.Fatalf(`result = %s; want "apbqcd", result`, result)
	}
}
