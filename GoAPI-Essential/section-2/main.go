package main

import "fmt"

func main() {
	//// var fullname string;
	//// or var fullName = "Pathipat Mattra";
	// var fullName string = "Pathipat Mattra"
	// var age int = 20
	// use this
	// fullName := "Pathipat Mattra"
	// age := 20
	// fmt.Printf("Hello %s Yay! age = %d\n", fullName, age)

	// fullName = fullName + "neuf"
	// age = age * 2
	// fmt.Println(fullName, age)

	var score int = 62
	var grade string

	if score >= 80 {
		grade = "A"
	} else if score >= 70 {
		grade = "B"
	} else if score >= 60 {
		grade = "C"
	} else {
		grade = "F"
	}

	fmt.Printf("Your grade is %s", grade)
}
