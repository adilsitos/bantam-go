package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/adilsitos/bantam-go/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Panic("could not initilize program")
	}

	fmt.Printf("Hello %s! This is the bantam \n A toy programming language to learn how pratt parsing works", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
