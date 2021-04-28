package main

import "fmt"

type person struct {
	name string
	age int
}

func main(){
	fmt.Println(person{"=3=", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{age: 30})

	s := person{"030", 25}
	fmt.Println(s.name)
	fmt.Println(s.age)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
