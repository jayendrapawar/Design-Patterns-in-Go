package main

import "fmt"

type BasePizza interface {
	getPrice() int
}

type VeggieMania struct{}

func (p *VeggieMania) getPrice() int {
	return 15
}

type TomatoToppings struct {
	basePizza BasePizza
}

func (t *TomatoToppings) getPrice() int {
	return t.basePizza.getPrice() + 7
}

type CheeseToppings struct {
	basePizza BasePizza
}

func (c *CheeseToppings) getPrice() int {
	return c.basePizza.getPrice() + 10
}

func main() {
	p := VeggieMania{}
	t := TomatoToppings{basePizza: &p}
	c := CheeseToppings{basePizza: &t}
	price := c.getPrice()
	fmt.Println("The price of the pizza is:", price)
	fmt.Println("The price of the pizza is:", price)
}