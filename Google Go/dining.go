package main

import (
	"fmt"
	"sync"
)

func remove(s []int, i int) []int {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

type ChopS struct { sync.Mutex }

type Philo struct {
	index int
	leftCS, rightCS *ChopS
}

var currently_eating []int
var append_currently_eating sync.Mutex

func (p * Philo) ask_host(channel chan int) {
	// fmt.Println("Philosopher", p.index, "asking host,", currently_eating)
	for {
		if len(currently_eating) ==0 {
			append_currently_eating.Lock()
			currently_eating = append(currently_eating, p.index)
			fmt.Println("Array isnide ask_host -->", currently_eating)
			append_currently_eating.Unlock()
			channel <- 1
			return 
		}
		if len(currently_eating)>=2 {
			continue
		}
		for _, value := range currently_eating {
			if((value+1)%5!=p.index && (value+4)%5!=p.index) {
				append_currently_eating.Lock()
				currently_eating = append(currently_eating, p.index)
				fmt.Println("Array isnide ask_host -->", currently_eating)
				append_currently_eating.Unlock()
				channel <- 1
				return
			}
		}
	}
	
}

func (p *Philo) eat(wg *sync.WaitGroup) {
	fmt.Println("Entered -->", p)
	mychannel := make(chan int)
	for i:=0;i<3;i++ {
		go p.ask_host(mychannel)
		<- mychannel

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("Philosopher", p.index, "is eating -->", currently_eating)

		// fmt.Println("Array -->", currently_eating)
		append_currently_eating.Lock()
		if currently_eating[0]==p.index {
			currently_eating = remove(currently_eating, 0)	
		} else {
			currently_eating = remove(currently_eating, 1)
		}
		fmt.Println("Philosopher", p.index, "finished eating -->", currently_eating)
		append_currently_eating.Unlock()
		
		p.leftCS.Unlock()
		p.rightCS.Unlock()
	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	CSticks := make([]*ChopS, 5)

	for i:=0;i<5;i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i:=0;i<5;i++ {
		philos[i] = & Philo{i, CSticks[i], CSticks[(i+1)%5]}
	}

	wg.Add(5)
	for i:=0;i<5;i++ {
		go philos[i].eat(&wg)
	}
	wg.Wait()
}