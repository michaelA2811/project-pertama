package main

import "fmt"

func main() {
	// var name string = "John"
	// var age int = 25
	// fmt.Println("hello world")
	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println("Genap")
	// 	} else {
	// 		fmt.Println("Ganjil")
	// 	}

	// }
	nama := []string{"Budi", "Anies", "John", "Doe", "Lorem"}
	for i := 0; i < len(nama); i++ {
		fmt.Println(nama[i])
	}
}
