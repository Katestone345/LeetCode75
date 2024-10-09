package challenges

import "fmt"

func CanPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 1 {
		if flowerbed[0] == 0 {
			n--
			return n <= 0
		}
	} else if len(flowerbed) == 2 {
		if flowerbed[0] == 0 && flowerbed[1] == 0 {
			n--
			return n <= 0
		}
	}

	if flowerbed[0] == 0 && flowerbed[1] == 0 {
		n--
		flowerbed[0] = 1
	}

	for i := 1; i < len(flowerbed)-1; i++ {
		if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 && flowerbed[i] != 1 {
			flowerbed[i] = 1
			n--
		}
	}

	if flowerbed[len(flowerbed)-1] == 0 && flowerbed[len(flowerbed)-2] == 0 {
		n--
	}
	fmt.Println(n)
	return n <= 0
}
