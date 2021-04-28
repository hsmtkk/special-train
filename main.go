package main

import (
	"fmt"

	"github.com/hsmtkk/special-train/armstrong"
)

func main() {
	for i := 1; i < 1000; i++ {
		if armstrong.IsArmstrong(i) {
			fmt.Println(i)
		}
	}
}
