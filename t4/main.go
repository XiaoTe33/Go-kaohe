package main

import "fmt"

var eggMap = map[int]bool{}

func main() {
	var eggSlice []int
	n := 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		t := 0
		k := 0
		fmt.Scan(&t)
		fmt.Scan(&k)
		x := 0
		for j := 0; j < k; j++ {
			fmt.Scan(&x)
			eggMap[x] = true
		}
		eggSlice = append(eggSlice, len(eggMap))
	}
	for i, _ := range eggSlice {
		fmt.Println(eggSlice[i])
	}
}
