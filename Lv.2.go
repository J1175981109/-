package main

import (
	"fmt"
)

import "sync"

var (
   myres = make(map[int]int, 20)
    mu    sync.Mutex
)

func factorial(n int) {
	ch := make(chan int)
	var res = 1
	var x int
    for i := 1; i <= n; i++ {
		        res *= i
		   }
    	go func(){
    		ch <- res
		}()
    x = <- ch

	    myres[n] = x

	}
func main() {
	    for i := 1; i <= 20; i++ {
		        go factorial(i)
		    }

	for i, v := range myres {
		fmt.Printf("myres[%d] = %d\n", i, v)
	}

	}
