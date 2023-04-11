package main
import (
	"fmt"
)

func GenDisplaceFn(a, v0, s0 float64) func(time float64) float64{
	return func(time float64) float64 {
		return (0.5 * a * time * time) + (v0 * time) + (s0)
	}
}

func main() {
	var acc float64
	var v0 float64
	var s0 float64

	fmt.Print("Enter the value of  initial displacement: ")
	fmt.Scan(&s0)

	fmt.Print("Enter the value of initial velocity: ")
	fmt.Scan(&v0)

	fmt.Print("Enter the value of acceleration: ")
	fmt.Scan(&acc)

	fn := GenDisplaceFn(acc, v0, s0)

	var time float64
	fmt.Print("Enter the value of time: ")
	fmt.Scan(&time)

	fmt.Print("The final displacement is: ")
	fmt.Println(fn(time))
}