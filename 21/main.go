package main

import "fmt"

type computer interface {
	insertInSquarePort()
}
type windows struct{}
type mac struct{}
type client struct{}
type windowsAdapter struct {
	windowMachine *windows
}

func (m *mac) insertInSquarePort() {
	fmt.Println("Insert square port into mac machine")
}

func (w *windowsAdapter) insertInSquarePort() {
	w.windowMachine.insertInCirclePort()
}

func (w *windows) insertInCirclePort() {
	fmt.Println("Insert circle port into windows machine")
}

func (c *client) insertSquareUsbInComputer(com computer) {
	com.insertInSquarePort()
}

func main() {
	client := &client{}
	mac := &mac{}
	client.insertSquareUsbInComputer(mac)
	windowsMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowMachine: windowsMachine,
	}
	client.insertSquareUsbInComputer(windowsMachineAdapter)
}
