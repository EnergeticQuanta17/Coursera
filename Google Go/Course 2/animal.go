package main

import "fmt"

func main() {
	var command, name, info string
	animalMap := make(map[string]Animal)

	cow := Cow{"grass", "walk", "moo"}
	bird := Bird{"worms", "fly", "peep"}
	snake := Snake{"mice", "slither", "hsss"}

	for {
		fmt.Printf(">")
		fmt.Scanln(&command, &name, &info)

		switch command {
		case "newanimal":
			switch info {
			case "cow":
				animalMap[name] = cow
			case "bird":
				animalMap[name] = bird
			case "snake":
				animalMap[name] = snake
			default:
				fmt.Println("Invalid animal type!")
				continue
			}
			fmt.Println("Created it!")
		case "query":
			if animalMap[name] != nil && checkAct(info) {
				animal := animalMap[name]
				doQuery(animal, info)
			} else {
				fmt.Println("Animal not found")
			}
		}
	}
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food, locomotion, noise string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.noise)
}

type Bird struct {
	food, locomotion, noise string
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b Bird) Speak() {
	fmt.Println(b.noise)
}

type Snake struct {
	food, locomotion, noise string
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s Snake) Speak() {
	fmt.Println(s.noise)
}

func checkAct(act string) bool {
	if act == "eat" || act == "move" || act == "speak" {
		return true
	}
	return false
}

func doQuery(animal Animal, act string) {
	switch act {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	}
}
