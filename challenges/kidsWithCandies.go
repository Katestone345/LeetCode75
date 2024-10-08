package challenges

import "errors"

func checkCandies(candies []int, extraCandies int) error {
	if len(candies) < 2 || len(candies) > 100 {
		return errors.New("insufficient number of kids")
	}
	for i := 0; i < len(candies); i++ {
		if candies[i] < 1 || candies[i] > 100 {
			return errors.New("insufficient number of candies per kid")
		}
	}
	if extraCandies < 1 || extraCandies > 50 {
		return errors.New("insufficient number of extraCandies")
	}
	return nil
}

func getTotalCandies(candies []int, extraCandies int) ([]int, error) {
	newCandies := make([]int, len(candies))
	for i := 0; i < len(candies); i++ {
		newCandies[i] = candies[i] + extraCandies
	}

	return newCandies, nil
}

func getWinner(newCandies, oldCandies []int) ([]bool, error) {

	win := 0
	var result []bool

	for i := 0; i < len(newCandies); i++ {
		for j := 0; j < len(oldCandies); j++ {
			if newCandies[i] >= oldCandies[j] {
				win += 1
			}
		}

		if win == len(newCandies) {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
		win = 0

	}
	return result, nil
}

func KidsWithCandies(candies []int, extraCandies int) []bool {

	// n = candy.length // amount of kids
	// 2 <= n <= 100
	// 1 <= candies[i] <= 100
	// 1 <= extraCandies <= 50

	err := checkCandies(candies, extraCandies)
	if err != nil {
		return []bool{false, false}
	}

	newCandies, err := getTotalCandies(candies, extraCandies)
	if err != nil {
		return []bool{false, false}
	}

	result, err := getWinner(newCandies, candies)
	if err != nil {
		return []bool{false, false}
	}

	return result
}
