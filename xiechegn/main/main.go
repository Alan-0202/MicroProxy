package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func main(){

	go func() {
		for i := 0; i < 100; i++{

			if i == 10{
				runtime.Gosched()
			}
			fmt.Println("The First 1:" + strconv.Itoa(i))
		}
	}()


	go func() {
		for i := 100; i < 200; i++{
			fmt.Println("The Second 2:" + strconv.Itoa(i))
		}
	}()


	time.Sleep(time.Second * 5)


}