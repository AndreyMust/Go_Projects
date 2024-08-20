package main

import (
	"fmt"
	"sync"
)

func asyncQuicksort(arr []int, wg *sync.WaitGroup) {
	defer wg.Done()

	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1
	pivotIndex := len(arr) / 2
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	var subWg sync.WaitGroup
	subWg.Add(2)

	go asyncQuicksort(arr[:left], &subWg)
	go asyncQuicksort(arr[left+1:], &subWg)

	subWg.Wait()
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("Исходный массив:", arr)

	var wg sync.WaitGroup
	wg.Add(1)

	go asyncQuicksort(arr, &wg)
	wg.Wait()
	fmt.Println("Отсортированный массив:", arr)
}
