package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type user struct {
	Username string
	Nickname string
	Sex      uint8
	Birthday time.Time
}

// main 因为user结构体的各个字段首字母没有大写，
// 导致无法被外界访问，大写一下就好了
func main() {
	u := user{
		Username: "坤坤",
		Nickname: "阿坤",
		Sex:      20,
		Birthday: time.Now(),
	}
	bs, err := json.Marshal(&u)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(bs))
}
