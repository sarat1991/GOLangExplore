package main

import "fmt"

func main() {

	varArr := []int{ 1, 2, 3, 4, 5, 6}

	for index, value := range varArr {
		fmt.Println(index," ", value)
	}

	// ignoring the index

	for _ , value :=range varArr{
		fmt.Println("value: ", value)
	}

	var mp = make(map[string]string)

	mp["sam"] = "sam@gmail.com"
	mp["mike"] = "mike@gmail.com"

	for k,v := range mp {
		fmt.Println(k," ", v)
	}
    for k := range mp {
		fmt.Println("key is ", k)
	}

	for _,v := range mp {
		fmt.Println("Value is ", v)
	}
}