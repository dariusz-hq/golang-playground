package main

import "fmt"

func main() {
	ht := NewHashTable(50)

	ht.set("myProp1", 1)
	fmt.Println(ht.get("myProp1"))

	ht.set("myProp1", 2)
	fmt.Println(ht.get("myProp1"))

	ht.set("myProp2", "blabla")
	fmt.Println(ht.get("myProp2"))
}
