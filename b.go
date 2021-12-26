package main

import "fmt"

var ch chan int = make (chan int)
func main() {

go func() {
fmt.Println("下山的路又堵起了")
ch<-0
}()
<-ch
}