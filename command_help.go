package main

import (
	"fmt"
)

func commandHelp() error{
	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Usage:")
	cmds := getCommand()
	for _, cmd := range cmds{
		fmt.Printf("%v: %v\n", cmd.cmd, cmd.description)
	}
	return nil
}
