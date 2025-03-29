package main

import (
	"fmt"
	"os"
)

func main() {

	loadExistingList()

	if len(os.Args) < 3 {
		fmt.Println("Not enough arguments to do anything")
		return
	}

	command := os.Args[1]
	task := os.Args[2]

	if verifyCommand(command) == "" {
		fmt.Println("Not a valid Command for this program")
		return
	}

	fmt.Println("Command to run: ", command)
	fmt.Println("This is what we want to add to list", task)
	if command == "add" {
		addFunctionality()
	}
}

func verifyCommand(command string) string {

	differentCommands := [5]string{"add", "remove", "list", "update", "delete"}
	for i := range differentCommands {
		if command == differentCommands[i] {
			return differentCommands[i]
		}
	}
	return string("")
}

func addFunctionality() {
	fmt.Println("Adding to the list of things to do...")

}

func loadExistingList() {

}
