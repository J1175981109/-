package main

import (
	"fmt"
	"time"
)
var sushu = make(map[int]int)
var x int

func main(){

	for i := 1; i <= 30 ;i++ {
		go Sushu(i)
		time.Sleep(5)
		fmt.Printf("sushu[%d]=%d\n",i,x)
	}

}
func Sushu(n int){

	ch := make(chan int)

	for i := 3; i <= n; i++ {
		for j := 2; j <= i/2; j++ {

			if i % j == 0 {
				break
			} else {
				go func() {
					ch <- i
				}()
				}
			}
		}

	x = <- ch

	sushu[n] = x
}