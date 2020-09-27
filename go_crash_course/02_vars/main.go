package main

import ("fmt"
"math" 
"github.com/sarath1991/go_crash_course/03_package"
)
var t = "sam"
func main() {
	s := "Hello"
	p := 2
	const isCool = true
	fmt.Print(s)
	fmt.Print(p)
	fmt.Printf("%T", p)
	fmt.Print(isCool)
	fmt.Print(t)
	
	fvar :=8.9

	name, email := "sarath", "sarath1991@gmail.com"
	fmt.Print(name, " ", email)
	fmt.Printf("%T", fvar)
	fmt.Println(math.Floor(2.25))
	fmt.Println(stgutil.Reverse("olleh"))


}
