package main

import (
	"fmt"
	"time"
)

func runloop(num *int,s string){
for i:=1;true;i++ {
*num=*num+1
fmt.Println(s,*num)
time.Sleep(time.Second*1)
}}

func main() {
num := 0
go runloop(&num,"a1")
go runloop(&num,"b1")

//Race condition in Go occurs when two or more goroutines have shared data and interact with it simultaneously.
//here num is shared by 4 goroutines above
time.Sleep(time.Second*10)
fmt.Println("Program end")
}
