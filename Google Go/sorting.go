package main

import (
	"fmt"
	"sync"
	"sort"
)

func sorting(array []int, wg *sync.WaitGroup) {
	sort.Ints(array)
	fmt.Println(array)
	wg.Done()
}

func merge(a []int, b []int) []int {
    final := []int{}
    i := 0
    j := 0
    for i < len(a) && j < len(b) {
        if a[i] < b[j] {
            final = append(final, a[i])
            i++
        } else {
            final = append(final, b[j])
            j++
        }
    }
    for ; i < len(a); i++ {
        final = append(final, a[i])
    }
    for ; j < len(b); j++ {
        final = append(final, b[j])
    }
    return final
}

func main() {
	var wg sync.WaitGroup

	var size int
	fmt.Print("Enter the size of the array:")
	fmt.Scanln(&size)
	arr := make([]int,size)
	fmt.Println("Enter the elements to be sorted (space separated): ")
	for i:=0;i<size;i++ {
		fmt.Scanf("%d", &arr[i])
	}

	partition_size := size/4
	wg.Add(4)

	fmt.Println("The size of partition_size:", partition_size)
	for i:=0;i<3;i++ {
		go sorting(arr[partition_size*i:partition_size*(i+1)], &wg)
	}
	go sorting(arr[partition_size*(3):size], &wg)

	wg.Wait()

	// merging two partitions at a time
	first := merge(arr[:len(arr)/4],arr[len(arr)/4:len(arr)/2])
	second := merge(arr[len(arr)/2:(len(arr)/2)+(len(arr)/4)],arr[(len(arr)/2)+(len(arr)/4):])

	third := merge(first, second)
	fmt.Println("The sorted array is :",third)
	
}	

