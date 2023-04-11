package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() string {
	return a.food
}

func (a *Animal) Move() string {
	return a.locomotion
}

func (a *Animal) Speak() string {
	return a.noise
}

// Greetings Program
func main() {
	dict := map[string]Animal{
		"cow":   {food: "grass", locomotion: "walk", noise: "moo"},
		"bird":  {food: "worms", locomotion: "fly", noise: "peep"},
		"snake": {food: "mice", locomotion: "slither", noise: "hsss"},
	}

	for {
		animal, action := getData()
		a := dict[animal]

		switch action {
		case "eat":
			fmt.Println(a.Eat())
		case "move":
			fmt.Println(a.Move())
		case "speak":
			fmt.Println(a.Speak())
		}
	}
}

func getData() (string, string) {
	fmt.Println("> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	vars := strings.Fields(scanner.Text())
	return vars[0], vars[1]
}