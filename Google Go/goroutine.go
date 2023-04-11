package main

import (
	"fmt"
	"time"
)
func displayOne() {
	fmt.Println("First line of displayOne()")
	time.Sleep(1*time.Second)	
	fmt.Println("Middle line of displayOne()")
	time.Sleep(1*time.Second)
	fmt.Println("Last line of displayOne()")
}

func displaySecond() {
	fmt.Println("First line of displaySecond()")
	time.Sleep(1*time.Second)	
	fmt.Println("Middle line of displayOne()")
	time.Sleep(1*time.Second)
	fmt.Println("Last line of displaySecond()")
}
func main() {
	go displayOne()
	go displaySecond()
	time.Sleep(5*time.Second)	
	fmt.Println("\nRace condition happens when multiple go routines interleave the statements to be executed")
	fmt.Println("We can see that in some executions the order of calling of functions is different")
	fmt.Println("Because there is a race condition for which fmt.Println() is printed first")
	fmt.Println("\nPlease do try the observation of different outputs on multiple executions")
}
