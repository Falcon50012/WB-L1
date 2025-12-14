// Удалить i-ый элемент из слайса.
// Продемонстрируйте корректное удаление без утечки памяти.
// Подсказка:
// можно сдвинуть хвост слайса на место удаляемого элемента
// (copy(slice[i:], slice[i+1:])) и уменьшить длину слайса на 1.

package main

import "fmt"

func main() {
	sl := []*int{new(int), new(int), new(int), new(int), new(int)}
	*sl[0], *sl[1], *sl[2], *sl[3], *sl[4] = 1, 2, 3, 4, 5

	fmt.Print("До удаления: ")
	for _, ptr := range sl {
		fmt.Print(*ptr, " ")
	}
	fmt.Println()

	idx := 2
	copy(sl[idx:], sl[idx+1:])
	sl[len(sl)-1] = nil
	sl = sl[:len(sl)-1]

	fmt.Print("После удаления: ")
	for _, ptr := range sl {
		fmt.Print(*ptr, " ")
	}
	fmt.Println()
}
