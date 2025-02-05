package main

import(
	"strings"
)

func main(){
	repl()
}

func cleanInput(text string) []string{
	lower := strings.ToLower(text)
	sliced := strings.Fields(lower)
	return sliced
}
