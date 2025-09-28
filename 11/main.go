// Реализовать пересечение двух неупорядоченных множеств
// (например, двух слайсов) — т.е. вывести элементы,
// присутствующие и в первом, и во втором.

// Пример:
// A = {1,2,3}
// B = {2,3,4}
// Пересечение = {2,3}

package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	res := []int{}
	m := make(map[int]bool)

	for _, num := range nums1 {
		m[num] = true
	}

	for _, num := range nums2 {
		if m[num] {
			res = append(res, num)
		}
		m[num] = false
	}

	fmt.Println(res)
}
