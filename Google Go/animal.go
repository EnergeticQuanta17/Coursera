package main
import (
	"fmt"
)

type Animal struct{
	food string
	locomotion string
	noise string
}

func (a Animal) Eat() {
	fmt.Println("The animal is eating", a.food)
}

func (a Animal) Move() {
	fmt.Println("The animal is "+ a.locomotion+"ing")
}

func (a Animal) Speak() {
	fmt.Println("The animal is making noise:", a.noise)
}

func function_call(instance Animal, info string) {
	switch(info) {
		case "eat":
			instance.Eat()
		case "move":
			instance.Move()
		case "speak":
			instance.Speak()
		default:
			fmt.Println("Incorrect information request!")
	}
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	var name string
	var info string

	for {
		fmt.Print("> ")
		fmt.Scanln(&name, &info)

		switch(name) {
			case "cow":
				instance := cow
				function_call(instance, info)
			case "bird":
				instance := bird
				function_call(instance, info)
			case "snake":
				instance := snake
				function_call(instance, info)
			default:
				fmt.Println("Incorrect name of animal!")
		}

	}
}