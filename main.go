package main

import (
	"fmt"
	"os"

	"github.com/aceagles/EagleLang/repl"
)

func main() {
	fmt.Println("Welcome to the EagleLang REPL!")
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
