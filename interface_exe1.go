//anonymous function and closure in go
package main

import "fmt"

type shape struct {
	length float32
	breadth float32
	height float32
}
func (s1 shape) area() float32 {
	return s1.length*s1.breadth
}
func (s1 shape) volume() float32 {
	return s1.length*s1.breadth*s1.height
}


type myinterface interface{
	area() float32
	volume() float32
}



func main() {
var land myinterface
land = shape{length:2,breadth: 3,height:4}
fmt.Println("area:",land.area())
fmt.Println("volume:",land.volume())

}