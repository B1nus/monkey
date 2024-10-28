package main

import (
	"fmt"
	"monkey/repl"
	"os"
)

func main() {
	username := os.Getenv("USER")     // Works on Unix/Linux/Mac
  if username == "" {
    username = os.Getenv("USERNAME") // Works on Windows
  }
	fmt.Printf("Hello %s! This is the Monkey programming language!", username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}


//   .--. .-"       "-. .--.
//  / .. \/  .-. .-.  \/ .. \
// | |  '|  /   Y   \  |' |  |
// | \   \  \ 0 | 0 /  /  /  |
//  \ '- ,\.-"""""""-./, -' /
//   ''-' /_   ^ ^   _\ '-''
//       |  \._   _./  |
//       \   \ '~' /   /
//        '._ '-=-' _.'
//           '-----'
