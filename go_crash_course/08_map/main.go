package main

import "fmt"

func main() {
	mp :=map[string]string{"bob": "bob@gmail.com", "mike": "mike@gmail.com"}
	
	var newMp = make(map[string]string)

	newMp["sam"] = "sam@gmail.com"
	fmt.Println(mp)
	fmt.Println(len(newMp))
	delete(newMp, "sam")
	fmt.Println(newMp["sam"])
}