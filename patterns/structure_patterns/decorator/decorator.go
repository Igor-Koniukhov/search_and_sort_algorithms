package main

import "fmt"

type Pizza interface {
	GetPrice() int
}

type PlainPizza struct{}

func (p *PlainPizza) GetPrice() int {
	return 10
}

// Base decorator
type PizzaDecorator struct {
	Pizza Pizza
}

func NewPizzaDecorator(pizza Pizza) *PizzaDecorator {
	return &PizzaDecorator{Pizza: pizza}
}

// CheeseDecorator adds cheese to the pizza.
type CheeseDecorator struct {
	*PizzaDecorator
}

func (c *CheeseDecorator) GetPrice() int {
	return c.Pizza.GetPrice() + 5
}

func NewCheeseDecorator(pizza Pizza) Pizza {
	return &CheeseDecorator{&PizzaDecorator{pizza}}
}

// PepperoniDecorator adds pepperoni to the pizza.
type PepperoniDecorator struct {
	*PizzaDecorator
}

func (p *PepperoniDecorator) GetPrice() int {
	return p.Pizza.GetPrice() + 7
}

func NewPepperoniDecorator(pizza Pizza) Pizza {
	return &PepperoniDecorator{&PizzaDecorator{pizza}}
}

func main() {
	plainPizza := &PlainPizza{}
	cheesePizza := NewCheeseDecorator(plainPizza)
	pepperoniPizza := NewPepperoniDecorator(cheesePizza)

	fmt.Printf("Plain Pizza Price: $%d\n", plainPizza.GetPrice())
	fmt.Printf("Cheese Pizza Price: $%d\n", cheesePizza.GetPrice())
	fmt.Printf("Pepperoni Pizza Price: $%d\n", pepperoniPizza.GetPrice())
}
