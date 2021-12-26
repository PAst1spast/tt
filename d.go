package main

import "fmt"

func main() {
	ch := Numbers() //
	for i := 0; i < 1000000; i++ {
		prime := <-ch //
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime)
	}
}
func Numbers() chan int {        //Numbers()用来生成数字
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}
//素数筛源自https://zhuanlan.zhihu.com/p/91006073
func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}


