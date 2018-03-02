package main

import "fmt"

func foo() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}


func main() {
	foo()
}