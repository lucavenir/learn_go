package main

import "fmt"

type owner struct {
	name string
	age  int
}

func (o owner) canDrive() bool {
	return o.age >= 18
}

type gasEngine struct {
	kmplt  uint8
	liters uint8
	owner
}

func (e gasEngine) kilometersLeft() uint16 {
	return uint16(e.liters * e.kmplt)
}

type electricEngine struct {
	kmpkwh uint8
	kwh    uint8
	owner
}

func (e electricEngine) kilometersLeft() uint16 {
	return uint16(e.kwh * e.kmpkwh)
}

type engine interface {
	kilometersLeft() uint16
}

func canMakeItTo(engine engine, kilometers uint16) bool {
	return engine.kilometersLeft() > kilometers
}

func main() {
	gas := gasEngine{16, 16, owner{"iipip", 18}}
	fmt.Println(canMakeItTo(gas, 100))

	el := electricEngine{60, 20, owner{"asdasd", 20}}
	fmt.Println(canMakeItTo(el, 100))
}
