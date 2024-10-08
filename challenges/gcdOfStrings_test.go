package challenges_test

import (
	"github.com/leetcode75/challenges"
	"testing"
)

func Test_GCDOfStrings(t *testing.T) {
	var result string

	result = challenges.GCDOfStrings("TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX")
	if result != "TAUXX" {
		t.Fatalf(`result = %s; want "TAUXX"`, result)
	}

	result = challenges.GCDOfStrings("ABCABC", "ABC")
	if result != "ABC" {
		t.Fatalf(`result = %s; want "ABC"`, result)
	}

	result = challenges.GCDOfStrings("ABABAB", "ABAB")
	if result != "AB" {
		t.Fatalf(`result = %s; want "AB", result`, result)
	}

	result = challenges.GCDOfStrings("LEET", "CODE")
	if result != "" {
		t.Fatalf(`result = %s; want "", result`, result)
	}

	result = challenges.GCDOfStrings("ABCDEF", "ABC")
	if result != "" {
		t.Fatalf(`result = %s; want "", result`, result)
	}
}
