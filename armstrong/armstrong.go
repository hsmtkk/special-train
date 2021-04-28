package armstrong

import (
	"log"
	"math"
	"strconv"
)

func IsArmstrong(n int) bool {
	s := 0
	p := len(strconv.Itoa(n))
	for _, d := range NumberToDigit(n) {
		s += int(math.Pow(float64(d), float64(p)))
	}
	return n == s
}

func NumberToDigit(n int) []int {
	result := []int{}
	for _, ch := range strconv.Itoa(n) {
		d, err := strconv.Atoi(string(ch))
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, d)
	}
	return result
}
