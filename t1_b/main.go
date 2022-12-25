package main

import "fmt"

// main 在第二个defer前就已经return了，所以defer没有到达
// 那也就没法输出，注释掉ta
func main() {
	var a = true
	defer func() {
		fmt.Println("1")
	}()

	if a {
		fmt.Println("2")
		//return
	}

	defer func() {
		fmt.Println("3")
	}()
}
