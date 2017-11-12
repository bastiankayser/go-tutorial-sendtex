package main

import (
	"time"
	"fmt"
	"sync"
)
var wg sync.WaitGroup

func cleanup() {
	defer wg.Done()
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup:", r)
	}
	
}

func says(s string) {
	defer cleanup()
	for i:=0;i<3;i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
		if i == 2 {
			panic("oh dear, a 2")
		}
	}
}



func main() {
	wg.Add(1)
	go says("Hey")
	wg.Add(1)
	go says("there")
	wg.Wait()
	
}