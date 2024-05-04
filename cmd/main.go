package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	for idx := 0; idx < 100; idx++ {
		idx := idx
		go func() {
			ch <- fmt.Sprintf("angka ke %d ", idx)
		}()
	}
	for i := range <-ch {
		fmt.Println(i)
	}
}
