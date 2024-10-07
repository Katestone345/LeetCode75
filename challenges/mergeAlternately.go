package challenges

import (
	"errors"
	"fmt"
)

// Helper function to verify that neither word 1 or word 2 is not over 100 or under 1
func verifyWordLength(lenWord1, lenWord2 int) error {

	if lenWord1 == 0 || lenWord2 == 0 {
		return errors.New(fmt.Sprintf("length of word1: %d \n length of word2: %d", lenWord1, lenWord2))
	}
	if lenWord1 > 100 || lenWord2 > 100 {
		return errors.New(fmt.Sprintf("length of word1: %d \n length of word2: %d", lenWord1, lenWord2))
	}

	return nil
}

func MergeAlternately(word1 string, word2 string) string {
	err := verifyWordLength(len(word1), len(word2))
	if err != nil {
		return err.Error()
	}

	var result string

	// If word 1 is larger than word 2
	if len(word1) > len(word2) {
		for i := 0; i < len(word1); i++ {
			if i > len(word2) || i == len(word2) {
				result += string(word1[i])
			}
			if i < len(word2) {
				result += string(word1[i])
				result += string(word2[i])
			}
		}
	}

	// If word 2 is larger than word 1
	if len(word2) > len(word1) {
		for i := 0; i < len(word2); i++ {
			if i > len(word1) || i == len(word1) {
				result += string(word2[i])
			}
			if i < len(word1) {
				result += string(word1[i])
				result += string(word2[i])
			}
		}
	}

	// Words are the same length
	if len(word1) == len(word2) {
		for i := 0; i < len(word1); i++ {
			result += string(word1[i])
			result += string(word2[i])
		}
	}

	return result
}
