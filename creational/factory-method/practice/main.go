package main

import "fmt"

type Animal interface {
	Talk()
	Go()
}

type Dog struct {}

func (Dog) Talk() {
	fmt.Println("Gâu Gâu")
}

func (Dog) Go() {
	fmt.Println("Dog goes ...")
}

type Cat struct {}

func (Cat) Talk() {
	fmt.Println("Meow Meow")
}

func (Cat) Go() {
	fmt.Println("Cat goes ...")
}

type AnimalService struct {
	animal Animal
}

func (a AnimalService) GoAndTalk() {
	a.animal.Go()
	a.animal.Talk()
}

func CreateNewAnimal(name string) Animal {
	switch name {
	case "dog":
		return new(Dog)
	case "cat":
		return new(Cat)
	}

	return new(Dog)
}

func main() {
	d := AnimalService{
		animal: CreateNewAnimal("dog"),
	}

	d.GoAndTalk()

	c := AnimalService{
		animal: CreateNewAnimal("cat"),
	}

	c.GoAndTalk()
}
