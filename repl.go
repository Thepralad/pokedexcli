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
		fmt.Println(word[0])
	}
}
