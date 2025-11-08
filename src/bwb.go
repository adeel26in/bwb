package main

import "fmt"

func main() {

	fmt.Println("Welcome to GoBitwarden!")

	var choice string

	fmt.Print("Enter your choice: ")

	fmt.Scan(&choice)

	switch choice {
	case "createp":
		createproject()
	case "listp":
		listprojects()
	case "getp":
		getprojects()
	case "listsec":
		listsecrets()
	case "getsec":
		getsecret()
	default:
		fmt.Println("Wrong choice!")
	}

}
