package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxDifference([]int32{2, 3, 10, 2, 4, 8, 1}))
}

func maxDifference(px []int32) int32 {
	// Write your code here
	// 2, 3, 10, 2, 4, 8, 1
	var terkecil = px[0]
	var maxNumber = px[1] - px[0]
	for i := 2; i < len(px); i++ {
		fmt.Println(i, terkecil, maxNumber)
		if px[i-1] < terkecil {
			terkecil = px[i-1]
		}
		// fmt.Println("selisih", px[i]-terkecil)
		if px[i]-terkecil > maxNumber {
			maxNumber = px[i] - terkecil
		}

	}
	if maxNumber == 0 {
		return -1
	} else {
		return maxNumber
	}
}
