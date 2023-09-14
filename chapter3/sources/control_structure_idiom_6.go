package main

import (
	"fmt"
	"time"
)

func main() {
	exit := make(chan interface{})

	go func() {
		for {
			select {
			case <-time.After(time.Second):
				fmt.Println("tick")
			case <-exit:
				fmt.Println("exiting...")
				break
			}
		}
		fmt.Println("exit!")
	}()
	time.Sleep(time.Second * 3)
	exit <- struct{}{}

	// 等子进程退出
	time.Sleep(time.Second * 3)
}