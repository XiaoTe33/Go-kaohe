package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	go Work("goroutine1")
	go Work("goroutine2")
	go Work("goroutine3")
	wg.Wait()
	fmt.Println("successful")
}

func Work(workName string) {
	time.Sleep(time.Second) // 模拟逻辑处理

	fmt.Println(workName)
	wg.Done()
}
