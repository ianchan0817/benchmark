package main

import (
	"fmt"
	"time"
)

func main() {
	current_time := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Println("********* Bechmark start *********")
	largets_number := 0
	for i := 0; i < 10000000; i++ {
		if largets_number < i {
			largets_number = i
		}
	}
	duration := time.Now().UnixNano()/int64(time.Millisecond) - current_time
	fmt.Println("It takes", duration, "ms\n**********************************\n")

}

// ********* Bechmark start *********
// It takes 7 ms
// **********************************
