package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
func S1(){
	fmt.Println("啊！")
	wg.Done()
}
func S2(){
	time.Sleep(1 * time.Second)
	fmt.Println("好痛苦！")
	wg.Done()
}
func S3(){
	time.Sleep(2 * time.Second)
	fmt.Println("channel好难啊！")
	wg.Done()
}
func S4(){
	time.Sleep(3 * time.Second)
	fmt.Println("还有好多没学的东西啊！")
	wg.Done()
}
func S5(){
	time.Sleep(4 * time.Second)
	fmt.Println("怎么办(⊙o⊙)？")
	wg.Done()
}
func S6(){
	time.Sleep(5 * time.Second)
	fmt.Println("药凉了吗/(ㄒoㄒ)/~~")
	wg.Done()
}
func S7(){
	time.Sleep(6 * time.Second)
	fmt.Println("没办法了")
	wg.Done()
}
func S8(){
	time.Sleep(7 * time.Second)
	fmt.Println("只能")
	wg.Done()
}
func S9(){
	time.Sleep(8 * time.Second)
	fmt.Println("边学")
	wg.Done()
}
func S10(){
	time.Sleep(9 * time.Second)
	fmt.Println("边做了/(ㄒoㄒ)/~~")
	wg.Done()
}
func main(){
	wg.Add(10)
	go S1()
	go S2()
	go S3()
	go S4()
	go S5()
	go S6()
	go S7()
	go S8()
	go S9()
	go S10()
	wg.Wait()
}