package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c Cow) Eat()   { fmt.Println("grass") }
func (c Cow) Move()  { fmt.Println("walk") }
func (c Cow) Speak() { fmt.Println("moo") }

type Bird struct{}

func (b Bird) Eat()   { fmt.Println("worms") }
func (b Bird) Move()  { fmt.Println("fly") }
func (b Bird) Speak() { fmt.Println("peep") }

type Snake struct{}

func (s Snake) Eat()   { fmt.Println("mice") }
func (s Snake) Move()  { fmt.Println("slither") }
func (s Snake) Speak() { fmt.Println("hsss") }

func main() {
	animals := make(map[string]Animal)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		words := strings.Fields(input)

		if len(words) != 3 {
			fmt.Println("Invalid input. Format: newanimal <name> <type> OR query <name> <action>")
			continue
		}

		command := words[0]
		name := words[1]
		param := words[2]

		if command == "newanimal" {
			switch param {
			case "cow":
				animals[name] = Cow{}
				fmt.Println("Created it!")
			case "bird":
				animals[name] = Bird{}
				fmt.Println("Created it!")
			case "snake":
				animals[name] = Snake{}
				fmt.Println("Created it!")
			default:
				fmt.Println("Unknown animal type. Choose cow, bird, or snake.")
			}
		} else if command == "query" {
			targetAnimal, exists := animals[name]
			if !exists {
				fmt.Println("Animal not found.")
				continue
			}

			switch param {
			case "eat":
				targetAnimal.Eat()
			case "move":
				targetAnimal.Move()
			case "speak":
				targetAnimal.Speak()
			default:
				fmt.Println("Unknown action. Choose eat, move, or speak.")
			}
		} else {
			fmt.Println("Unknown command. Please use 'newanimal' or 'query'.")
		}
	}
}
