package main

import "fmt"

type Observer interface {
	getId() string
	updateValue(string)
}

type Topic interface {
	register(observer Observer)
	broadcast()
}

type Item struct {
	observers []Observer
	name      string
	available bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) UpdateAvailable() {
	fmt.Printf("Item %s is available\n", i.name)
	i.available = true
	i.broadcast()
}

func (i *Item) broadcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

func (i *Item) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

//Ahora el observador que recibira los cambios
type EmailClient struct {
	id string
}

func (eC EmailClient) getId() string {
	return eC.id
}

func (eC EmailClient) updateValue(value string) {
	fmt.Printf("Sending email. . .  %s available for the client %s\n", value, eC.id)
}

func main() {
	macbook := NewItem("Mac Book Air")
	firstObserver := &EmailClient{
		id: "1234",
	}
	secondObserver := &EmailClient{id: "abcd"}
	macbook.register(firstObserver)
	macbook.register(secondObserver)
	macbook.UpdateAvailable()
}
