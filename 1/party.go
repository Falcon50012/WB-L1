// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action
// от родительской структуры Human (аналог наследования).
// Подсказка: используйте композицию (embedded struct),
// чтобы Action имел все методы Human.

package main

import (
	"flag"
	"fmt"
	"os"
)

type Human struct {
	name           string
	beverage       string
	beverageVolume int
}

func (h *Human) Drink() {
	if h.beverageVolume > 0 {
		ending := ending(h.beverageVolume)
		fmt.Printf("Ещё %d литр%s напитка класса %s, %s одобряет\n", h.beverageVolume, ending, h.beverage, h.name)
		h.beverageVolume--
	}
}

func (h *Human) Sleep() {
	fmt.Printf("Напиток класса %s закончился, %s отдыхает\n", h.beverage, h.name)
}

type Action struct {
	*Human
}

func main() {
	name := flag.String("n", "Иван Иваныч", "enter name")
	beverage := flag.String("b", "молоко", "enter beverage type")
	beverageVolume := flag.Int("v", 1000, "enter beverage volume")
	flag.Parse()

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%v Некрасива жизнь без %s...", r, *beverage)
			os.Exit(1)
		}
	}()

	if *beverageVolume <= 0 {
		panic("Критическая ситуация!")
	}

	party := Action{
		&Human{
			name:           *name,
			beverage:       *beverage,
			beverageVolume: *beverageVolume,
		},
	}

	for party.beverageVolume > 0 {
		party.Drink()
	}

	party.Sleep()
}

func ending(n int) string {
	n = n % 100

	if n >= 11 && n <= 14 {
		return "ов"
	}

	switch n % 10 {
	case 1:
		return ""
	case 2, 3, 4:
		return "а"
	default:
		return "ов"
	}
}
