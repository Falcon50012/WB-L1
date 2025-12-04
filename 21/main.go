package main

import "fmt"

type PCGamepad interface {
	ConnectToPC() string
}

type PSGamepad interface {
	ConnectToPS() string
}

type DualSense struct{}

func (ds *DualSense) ConnectToPS() string {
	return "DualSense успешно подключен"
}

type PSToPCAdapter struct {
	DualSense
}

func (a *PSToPCAdapter) ConnectToPC() string {
	psMessage := a.ConnectToPS()
	return fmt.Sprintf("%s к ПК через адаптер", psMessage)
}

func ConnectController(pad PCGamepad) {
	fmt.Println(pad.ConnectToPC())
}

func main() {
	adapter := &PSToPCAdapter{DualSense{}}

	ConnectController(adapter)
}
