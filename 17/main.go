// Реализовать алгоритм бинарного поиска встроенными методами языка.
// Функция должна принимать отсортированный слайс и искомый элемент,
// возвращать индекс элемента или -1, если элемент не найден.
// Подсказка:
// можно реализовать рекурсивно или итеративно, используя цикл for.

package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func binarySearch(nums []int, target int) int {
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		mid := lo + (hi-lo)/2

		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			lo = mid + 1
		} else if nums[mid] > target {
			hi = mid - 1
		}
	}

	return -1
}

func main() {
	nums := make(sort.IntSlice, 0, 10)

	for i := 0; i < cap(nums); i++ {
		nums = append(nums, rand.Intn(10)+1)
	}

	nums.Sort()
	fmt.Println("Отсортированный слайс:", nums)

	target := rand.Intn(10) + 1
	fmt.Println("Искомый элемент:", target)

	result := binarySearch(nums, target)

	if result == -1 {
		fmt.Println("Искомый элемент отсутствует")
	} else {
		fmt.Println("Индекс элемента:", result)
	}
}
