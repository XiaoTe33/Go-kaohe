package main

import "fmt"

func main() {
	for i := 0; i < 1000000; i++ {
		go JudgeNum(i)
	}
}
func JudgeNum(num int) {
	temp := num
	level := GetLevel(temp)
	t := 0
	for {
		if temp < 10 {
			t += GetChengfang(temp, level)
			break
		}
		t += GetChengfang(temp%10, level)
		temp /= 10
	}
	if t == num {
		fmt.Println(num)
	}

}

func GetLevel(num int) int {

	if num > 99999 {
		return 6
	}
	if num > 9999 {
		return 5
	}
	if num > 999 {
		return 4
	}
	if num > 99 {
		return 3
	}
	if num > 9 {
		return 2
	}
	return 1
}

func GetChengfang(num int, level int) int {
	a := 1
	for i := 0; i < level; i++ {
		a *= num
	}
	return a
}
