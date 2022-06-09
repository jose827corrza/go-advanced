package main

import "fmt"

/*
Esta interface no se va a preocupar cual clase la usa, ni cuantas la usan
*/
type IProduct interface {
	getStock() int
	setStock(n int)
	getName() string
	setName(name string)
}

func (c *Computer) getStock() int {
	return c.stock
}

func (c *Computer) setStock(n int) {
	c.stock = n
}

func (c *Computer) getName() string {
	return c.name
}

func (c *Computer) setName(name string) {
	c.name = name
}

type Computer struct {
	name  string
	stock int
}

type Laptop struct {
	Computer
}

type Desktop struct {
	Computer
}

/*
Aca esta el poder, por que me pide retornar un Iproduct y le devuelvo un Laptop
*/
func newLaptop() IProduct {
	return &Laptop{Computer{name: "Macbook Air", stock: 1}}
}

func newDesktop() IProduct {
	return &Desktop{Computer{name: "Asus", stock: 1}}
}

func GetComputerFactory(computerType string) (IProduct, error) {
	if computerType == "Laptop" {
		return newLaptop(), nil
	}
	if computerType == "Desktop" {
		return newDesktop(), nil
	}
	return nil, fmt.Errorf("Invalid PC type")
}

func PrintNameAndStock(p IProduct) {
	fmt.Printf("Product name: %s, and stock %d\n", p.getName(), p.getStock())
}

func main() {
	laptop, _ := GetComputerFactory("Laptop")
	desktop, _ := GetComputerFactory("Desktop")
	PrintNameAndStock(laptop)
	PrintNameAndStock(desktop)
}
