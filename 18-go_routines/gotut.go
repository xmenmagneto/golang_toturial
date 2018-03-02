package main

import (
	"time"
	"fmt"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 1000)
	}
}

func main() {
	say("Hey")
	go say("There")
}