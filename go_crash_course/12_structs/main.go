package main

import ("fmt"
"strconv"
)

type Person struct {
	firstName, lastName string
	age int
}
// value recievers
func (p Person) greeting () string{
	return "My name is "+p.firstName+" "+p.lastName+" and age is "+strconv.Itoa(p.age)
}

func (p *Person) hasBirthday() {
	p.age++; 
}

func (p *Person) getMarried(spouseName string) {
	p.lastName = spouseName

}

func main() {
	p1 :=Person{firstName: "Samantha", lastName : "Brad", age : 25}
	fmt.Println(p1)

	p2 := Person{"Samantha", "Brad", 25}
	fmt.Println(p2)
	p1.age++
	fmt.Println(p1.age)

	fmt.Println(p1.greeting())

	p1.hasBirthday()

	fmt.Println(p1.greeting())

	p1.getMarried("sam")

	fmt.Println(p1)

}