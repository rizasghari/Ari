package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/rizasghari/ari/lexer"
	"github.com/rizasghari/ari/parser"
	"github.com/rizasghari/ari/repl"
)

func main() {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	if program != nil {
		for _, stmt := range program.Statements {
			log.Printf("TestLetStatements stmt: %v", stmt)
		}
	}

	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Arı 🐝 programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
