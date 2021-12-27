package main

import "fmt"

var gold = 38

func main() {
	// for break 语句的使用 + label
	var sl = [][]int{
		{1, 34, 26, 35, 78},
		{3, 45, 13, 24, 99},
		{101, 13, 38, 7, 127},
		{54, 27, 40, 83, 81},
	}

outerloop:
	for i := 0; i < len(sl); i++ {
		for j := 0; j < len(sl[i]); j++ {
			if sl[i][j] == gold {
				fmt.Printf("found gold at [%d, %d]\n", i, j)
				break outerloop
			}
		}
	}
}
