package main

import (
	"fmt"
	"time"

	"github.com/sbstp/syncx"
)

func main() {
	writer, reader := syncx.TwistedPair()

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := reader.Read(buf)
			if err != nil {
				fmt.Println(err)
				break
			}
			data := buf[:n]
			fmt.Println("read:", string(data))
		}
	}()

	for i := 0; i < 100; i++ {
		n, err := fmt.Fprintf(writer, "hello world %d!", i)
		fmt.Println("wrote", i, n, err)
	}
	writer.Close()

	time.Sleep(5 * time.Second)

}
