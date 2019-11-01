package main

import (
	"fmt"
	"time"
)

func Sushu(n int) bool{
	var flag  = 0
	for i := 2; i <= n/2; i++{
		if n %i == 0{
			flag = 1
			break
		}
	}
	if flag == 1 {
		return false
	}else{
		fmt.Println(n)
	return true
	}
}

func main() {
	for i := 1; i <= 10000; i++ {
		go Sushu(i)
time.Sleep(5)
	}
}