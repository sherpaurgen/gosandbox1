package main

import (
	"fmt"
)

func main() {
	arr := []int{1,2,3,4,5,6,7,8,9,10,11}
	rem := len(arr)%4
	fmt.Printf("len of arr: %v &  rem :%v \n",len(arr),rem)
	fmt.Println(arr[len(arr)-rem:])
	fmt.Println(arr[0:1])
}