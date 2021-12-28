package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)


func main() {

	inputReader := bufio.NewReader(os.Stdin)
	animalSlice := []animal{}

	for {
		fmt.Println("Example command to create animal:")
		fmt.Println("newanimal rocky cow")
		fmt.Println("Example command to get action of created animal:use speak, move or eat")
		fmt.Println("query rocky speak")
		fmt.Print(">")

		userQuery, err := inputReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

				userQuery = strings.Trim(userQuery, "\n")
		querySlice := strings.Split(userQuery, " ")
		if len(querySlice) > 3 || len(querySlice) < 3 {
			fmt.Println("Invalid query")
		}
		req := querySlice[0]
        //fmt.Printf("......%T %v",querySlice,querySlice[2])
		switch req {
		case "newanimal":

			if querySlice[2] == "cow" {
				animalSlice = append(animalSlice, cow{name: querySlice[1]})
				fmt.Printf("Created\n\n")
			} else if querySlice[2] == "snake" {
				animalSlice = append(animalSlice, snake{name: querySlice[1]})
				fmt.Printf("Created\n\n")
			} else if querySlice[2] == "bird" {
				animalSlice = append(animalSlice, bird{name: querySlice[1]})
				fmt.Printf("Created\n\n")
			} else {
				fmt.Println("Invalid Query")
			}
		case "query":
			for _, animal := range animalSlice {
				if animal.getName() == querySlice[1] {
					if querySlice[2] == "move" {
						animal.move()
					} else if querySlice[2] == "eat" {
						animal.eat()
					} else if querySlice[2] == "speak" {
						animal.speak()
					}
				} else {
					//fmt.Println("Not found!")
				}
			}

		default:
			fmt.Println("Invalid user query")

		}
	}

}

type animal interface {
	eat()
	move()
	speak()
	getName() string
}

type cow struct{ name string }
type snake struct{ name string }
type bird struct{ name string }

func (c cow) getName() string {
	return c.name
}
func (s snake) getName() string {
	return s.name
}
func (b bird) getName() string {
	return b.name
}

func (c cow) eat() {
	fmt.Printf("%s eats grass\n\n", c.name)
}

func (c cow) move() {
	fmt.Printf("%s walks\n\n", c.name)
}

func (c cow) speak() {
	fmt.Printf("%s moos\n\n", c.name)
}

func (s snake) eat() {
	fmt.Printf("%s eats mice\n\n", s.name)
}

func (s snake) move() {
	fmt.Printf("%s slither\n\n", s.name)
}

func (s snake) speak() {
	fmt.Printf("%s hsss\n\n", s.name)
}

func (b bird) eat() {
	fmt.Printf("%s eats worm\n\n", b.name)
}

func (b bird) move() {
	fmt.Printf("%s flys\n\n", b.name)
}

func (b bird) speak() {
	fmt.Printf("%s peeps\n\n", b.name)
}