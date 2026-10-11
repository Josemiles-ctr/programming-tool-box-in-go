package goBasics

import "fmt"
type Person struct {
	Name string
	Age  int
}
func (p Person) Greet(){
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func (p Person) Eat(food string){
	fmt.Printf("%s is eating %s\n", p.Name, food)
}

