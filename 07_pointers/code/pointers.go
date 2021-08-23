package main

import "fmt"


type Coordinates struct {
	X, Y float64
}

var c = Coordinates{X: 10, Y: 20}

func main() {
	// print out pointer and value
	var name string
	var namePointer *string

	name = "Lawrence"
	namePointer = &name
	var nameValue =  *namePointer

	// print out pointer and value
	fmt.Println("Name:", name)
	fmt.Println("Name *:", namePointer)
	fmt.Println("Name Value", nameValue)

	//structs with pointers
	coordinatesMemoryAddress := &c 
	coordinatesMemoryAddress.X = 200

	// structs with pointers
	fmt.Println(coordinatesMemoryAddress)


}



// // ******************************************************

// func main() {
// 	var name string = "Beyonce"
// 	var namePointer *string = &name
// 	var nameValue = *namePointer

// 	fmt.Println("Name:", name)
// 	fmt.Println("Name *:", namePointer)
// 	fmt.Println("Name Value:", nameValue)

// }

// // ******************************************************

// func changeName(n string) {
// 	n = strings.ToUpper(n)
// }

// func main() {
// 	name := "Elvis"
// 	changeName(name)
// 	fmt.Println(name)
// }

// // ******************************************************
