package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/WhoDoIt/GoCompiler/internal/interpreter"
	"github.com/WhoDoIt/GoCompiler/internal/parser"
	"github.com/WhoDoIt/GoCompiler/internal/tokenizer"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Invalid number of arguments")
		os.Exit(1)
	}
	sourceName := os.Args[1]
	fmt.Println("Eval", sourceName)

	// resultName := os.Args[2]
	// fmt.Println("Eval", sourceName, "to", resultName)

	data, err := os.ReadFile(sourceName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	tk, err := tokenizer.Tokenize(data)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for _, v := range tk {
		fmt.Println("[" + strconv.Itoa(int(v.Type)) + ", \"" + v.Content + "\"]")
	}

	fmt.Println()

	expr, err := parser.Parse(tk)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(interpreter.StringVisitor{}.Print(expr))
	fmt.Println(interpreter.NumberEvalVisitor{}.Calculate(expr))
}
