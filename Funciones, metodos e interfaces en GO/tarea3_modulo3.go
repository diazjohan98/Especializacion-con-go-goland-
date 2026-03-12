package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Print(a.food)
}

func (a Animal) Move() {
	fmt.Print(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Print(a.noise)
}

func main() {
	animals := map[string]Animal{
		"cow":   {food: "grass", locomotion: "walk", noise: "moo"},
		"bird":  {food: "worms", locomotion: "fly", noise: "peep"},
		"snake": {food: "mice", locomotion: "slither", noise: "hsss"},
	}

	for {
		var animalName, action string
		fmt.Print("> ")

		_, err := fmt.Scan(&animalName, &action)
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		targetAnimal, exists := animals[animalName]
		if !exists {
			fmt.Println("Animal not found. Please choose 'cow', 'bird', or 'snake'.")
			continue
		}

		switch action {
		case "eat":
			targetAnimal.Eat()
		case "move":
			targetAnimal.Move()
		case "speak":
			targetAnimal.Speak()
		default:
		fmt.Println("Action not recognized. Please type 'eat', 'move', or 'speak'.")
		
		}

	}
}