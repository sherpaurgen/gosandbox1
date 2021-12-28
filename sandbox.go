package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)


func main() {
	fmt.Println("Example command to get data of animal with name cow,bird,snake & action speak, move or eat")
	fmt.Println("cow move")
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		userQuery, err := inputReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		userQuery = strings.Trim(userQuery, "\n")
		querySlice := strings.Split(userQuery, " ")
		if len(querySlice) != 2 {
			fmt.Println("Invalid query")
			break
		}
		req := querySlice[0]
		act := querySlice[1]
		//fmt.Printf("......%T %v",querySlice,querySlice[2])
		if req=="cow" {
			c1 := Animal{
				food:       "grass",
				locomotion: "walk",
				noise:      "moo",
			}
			if act =="speak" {
				c1.speak()
			} else if act=="move" {
				c1.move()
			} else if act=="eat"{
				c1.eat()
			} else {
				fmt.Println("Invalid cow action")
			continue
			}
		} else if req=="bird" {
			b1 := Animal{
				food:       "work",
				locomotion: "fly",
				noise:      "peep",
			}
			if act =="speak" { b1.speak()
			} else if act=="move" {
				b1.move()
			} else if act=="eat" {b1.eat()} else {
				fmt.Println("Invalid bird action")
				continue
			}
		} else if req=="snake" {
			s1 := Animal{
				food:       "mice",
				locomotion: "slither",
				noise:      "hsss",
			}
			if act =="speak" { s1.speak() } else if act=="move" {
				s1.move()
			} else if act=="eat" {s1.eat()} else {
				fmt.Println("Invalid snake action")
			}
		} else {
			fmt.Println("Invalid animal: only use cow, bird or snake")
			continue
		}
	}

}

type Animal struct{
	food string
    locomotion string
    noise string
}




func (c Animal) eat() {
	fmt.Printf(" %v\n\n",c.food)
}

func (c Animal) move() {
	fmt.Printf("%v\n\n",c.locomotion)
}

func (c Animal) speak() {
	fmt.Printf("%v\n\n",c.noise)
}

