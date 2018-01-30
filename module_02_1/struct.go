package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Student struct {
	Human
	speciality string
}

func main() {
	P := Student{Human: Human{name: "Steve", age: 20}, speciality: "Math"}
	fmt.Println(P.name)
	fmt.Println(P.age)
	fmt.Println(P.speciality)
}
