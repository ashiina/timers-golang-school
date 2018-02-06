package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) WalkTo(destination string) {
	fmt.Println("I'm walking to", destination)
}

func (h Human) Introduce() {
	fmt.Println("My name is", h.name, "and I am", h.age, "years old.")
}

type Student struct {
	Human
	speciality string
}

func (s *Student) SetSpeciality(speciality string) {
	s.speciality = speciality
}

func (s Student) Introduce() {
	fmt.Println("My name is", s.name, "and I am", s.age, "years old. I study", s.speciality)
}

type Walker interface {
	WalkTo(destination string)
}

func main() {
	bob := Human{name: "Bob", age: 60}
	bob.Introduce()

	steve := Student{Human: Human{name: "Steve", age: 30}, speciality: "Design"}
	steve.Introduce()
	steve.SetSpeciality("Business")
	steve.Introduce()

	WalkToNewYork(bob)
	WalkToNewYork(steve)
}

func WalkToNewYork(w Walker) {
	w.WalkTo("New York City")
}
