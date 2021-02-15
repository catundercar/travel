package main

import "fmt"

type Speaker interface {
	Speak()
}

type Eat interface {
	Eat()
}

type Human struct {
	Name string
}

func (h *Human) Speak() {
	fmt.Println(h.Name)
}

func (h *Human) Eat() {
	fmt.Println("eat")
}

type Cat struct {
}

func (c *Cat) Eat()  {
	fmt.Println("eat")
}

func main() {
	h := &Human{Name: "laoliu"}
	Eatting(h)

	c := &Cat{}
	Eatting(c)
}

func Eatting(eat Eat) {
	eat.Eat()
	s, ok := eat.(Speaker)
	if !ok {
		fmt.Println("not impl Speaker")
		return
	}
	s.Speak()
}