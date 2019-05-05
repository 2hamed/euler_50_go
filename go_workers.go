package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

func main() {
	ch := make(chan int)

	var ops int32
	for i := 0; i < 5; i++ {
		go func(ch <-chan int) {
			for c := range ch {
				fmt.Println("goroutine 1:", c)
				time.Sleep(50 * time.Millisecond)
				atomic.AddInt32(&ops, 1)
			}
		}(ch)
	}
	go func() {
		for {
			ch <- rand.Int()
			// time.Sleep(100 * time.Millisecond)
		}
	}()
	<-time.After(5 * time.Second)

	fmt.Printf("OPs count: %d", atomic.LoadInt32(&ops))

}
