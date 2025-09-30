// Реализовать алгоритм быстрой сортировки массива
// встроенными средствами языка. Можно использовать рекурсию.
// Подсказка:
// напишите функцию quickSort([]int) []int
// которая сортирует срез целых чисел.
// Для выбора опорного элемента можно взять середину или первый элемент.

package main

import (
	"fmt"
	"math/rand"
)

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivotIdx := len(nums) / 2
	pivot := nums[pivotIdx]

	var less, greater, mid []int

	for i := range nums {
		if nums[i] < pivot {
			less = append(less, nums[i])
		} else if nums[i] > pivot {
			greater = append(greater, nums[i])
		} else {
			mid = append(mid, nums[i])
		}
	}

	return append(append(quickSort(less), mid...), quickSort(greater)...)
}

func main() {
	nums := make([]int, 0, 10)

	for i := 0; i < cap(nums); i++ {
		nums = append(nums, rand.Intn(10)+1)
	}

	sortNums := quickSort(nums)

	fmt.Println("Исходный слайс:", nums)
	fmt.Println("Сортированный слайс:", sortNums)
}
