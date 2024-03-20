package main

import "fmt"

type Car struct {
	Type   string
	Seats  int
	Engine string
	GPS    bool
}

type Builder interface {
	SetType(string) Builder
	SetSeats(int) Builder
	SetEngine(string) Builder
	AddGPS() Builder
	Build() Car
}

type CarBuilder struct {
	car Car
}

func (b *CarBuilder) SetType(t string) Builder {
	b.car.Type = t
	return b
}

func (b *CarBuilder) SetSeats(seats int) Builder {
	b.car.Seats = seats
	return b
}

func (b *CarBuilder) SetEngine(engine string) Builder {
	b.car.Engine = engine
	return b
}

func (b *CarBuilder) AddGPS() Builder {
	b.car.GPS = true
	return b
}

func (b *CarBuilder) Build() Car {
	return b.car
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

func main() {
	sportsCar := NewCarBuilder().
		SetType("Sports").
		SetSeats(2).
		SetEngine("V8").
		AddGPS().
		Build()

	familyCar := NewCarBuilder().
		SetType("Family").
		SetSeats(5).
		SetEngine("V6").
		Build()

	fmt.Println("Sports Car:", sportsCar)
	fmt.Println("Family Car:", familyCar)
}
