package main

import (
	"fmt"
	"time"

	"github.com/sbstp/syncx"
)

func main() {
	e := syncx.NewEvent()

	for i := 0; i < 10; i++ {
		go func(x int) {
			e.Wait()
			fmt.Println("done", x, e.IsSet())
		}(i)
	}

	time.Sleep(time.Second * 2)

	fmt.Println(e.IsSet())
	e.Set()
	e.Set()

	time.Sleep(time.Second * 5)
}
