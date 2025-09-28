// Имеется последовательность строк:
// ("cat", "cat", "dog", "cat", "tree").
// Создать для неё собственное множество.
// Ожидается:
// получить набор уникальных слов.
// Для примера, множество = {"cat", "dog", "tree"}.

package main

import "fmt"

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]bool)

	for _, s := range strings {
		m[s] = true
	}

	uniqueS := make([]string, 0, len(m))

	for key := range m {
		uniqueS = append(uniqueS, key)
	}

	fmt.Println(uniqueS)
}
