package main

import (
	"fmt"
)

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	// 1. Create a new array that contain 3 hobies
	hobbies := [3]string{"Coding", "Tracking", "Exercise"}
	// Print full hobbies in terminal
	fmt.Println(hobbies)

	//2. Print first element from an hobbies
	fmt.Println(hobbies[0])

	//3. Print second and thirs element of an hobbies using slice method
	fmt.Println(hobbies[1:3])

	//4.1 Creating an new array of hobbies and assigning it to a new variable from existing array of hobies
	mainHobbies := hobbies[:2]
	//4.2 Printing new array of hobbies which we created using slice
	fmt.Println(mainHobbies)

	//4.3 re assigning an array which can contain value of 2 and 3 index
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	// 5. Creating a Dynamic array
	courseGole := []string{"Learn GO !", "Learn all the importance concept of GO"}
	fmt.Println(courseGole)

	// 6. Update 2nd Goal and add new goal at last
	courseGole[1] = "Clear all the concept of GO"
	// 6.1 add new goal at last append method is used to add the value in an array and slice
	courseGole = append(courseGole, "Learn GO and stand out of it")

	fmt.Println(courseGole)

	product := []Product{{id: "123", title: "Bag", price: 300}, {id: "131", title: "Shoes", price: 500}}
	product = append(product, Product{id: "321", title: "TShirt", price: 450})

	fmt.Println(product)

}
