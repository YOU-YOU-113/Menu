// main.go
package main

import "fmt"

func main() {
	var cmd string
	var doing string = "yes"
	for doing == "yes" {
		//fmt.Println("Please input:")
		fmt.Scan(&cmd)
		switch cmd {
		case "help":
			fmt.Println("This is help cmd")
		case "version":
			fmt.Println("menu program v1.0")
		case "quit":
			{
				fmt.Println("Quit from menu")
				doing = "no"
			}

		default:
			fmt.Println("Wrong cmd!")
		}
	}
}
