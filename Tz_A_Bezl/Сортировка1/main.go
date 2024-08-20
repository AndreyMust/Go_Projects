package main

import (
	"fmt"
)


func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1
	// Выбор опорного элемента
	pivotIndex := len(arr) / 2
	// Перестановка опорного элемента в конец
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Разделение массива
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Перемещаем опорный элемент на его окончательное место
	arr[left], arr[right] = arr[right], arr[left]
	// Рекурсивно сортируем левую и правую части
	quicksort(arr[:left])
	quicksort(arr[left+1:])
	return arr
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("Исходный массив:", arr)
	sortedArr := quicksort(arr)
	fmt.Println("Отсортированный массив:", sortedArr)
}
