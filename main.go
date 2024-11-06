package main

import (
	"fmt"
	"log"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"monkey/repl"
	"os"
)

const MONKEY_FACE = `
  .--. .-"       "-. .--.
 / .. \/  .-. .-.  \/ .. \
| |  '|  /   Y   \  |' |  |
| \   \  \ 0 | 0 /  /  /  |
 \ '- ,\.-"""""""-./, -' /
  ''-' /_   ^ ^   _\ '-''
      |  \._   _./  |
      \   \ '~' /   /
       '._ '-=-' _.'
          '-----'
`

func main() {
  if (len(os.Args) == 2) {
    content, err := os.ReadFile(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    l := lexer.New(string(content))
    p := parser.New(l)
		if len(p.Errors()) != 0 {
			printParserErrors(p.Errors())
      return
		}

    v := evaluator.Eval(p.ParseProgram(), object.NewEnvironment())
    if v != nil {
      fmt.Printf(v.Inspect())
      fmt.Printf("\n")
    }
    return
  }

	username := os.Getenv("USER")     // Works on Unix/Linux/Mac
  if username == "" {
    username = os.Getenv("USERNAME") // Works on Windows
  }
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

func printParserErrors(errors []string) {
  fmt.Printf(MONKEY_FACE)
  fmt.Printf("Woops! We ran into some monkey business here!\n")
  fmt.Printf(" parser errors:\n")
	for _, msg := range errors {
		fmt.Printf("\t"+msg+"\n")
	}
}
