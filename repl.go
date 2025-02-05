package main

import (
	"fmt"
	"os"
	"bufio"
)

func repl(){
	scanner := bufio.NewScanner(os.Stdin)
	
	for{
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		word := cleanInput(scanner.Text())
		for _, cmd := range getCommand(){
			if word[0] == cmd.cmd{
				cmd.callback()
			}
		}
	}
}

type cliCommand struct{
	cmd string
	description string
	callback func() error
}

func getCommand() map[string]cliCommand{
	return map[string]cliCommand{
		"exit": {cmd: "exit", description: "exits pokedex", callback: commandExit},
		"help": {cmd: "help", description: "show usage", callback: commandHelp},
	}
}
