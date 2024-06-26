package repl_interpreter

import (
	"bufio"
	"compiler/evaluator"
	"compiler/lexer"
	"compiler/object"
	"compiler/parser"
	"fmt"
	"io"
)

const PROMPT = ">> "

// var fibonacci = fn(x) { if (x == 0) { 0 } else { if (x == 1) { 1 } else { fibonacci(x - 1) + fibonacci(x - 2); } } };
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		_, _ = fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			_, _ = io.WriteString(out, evaluated.Inspect())
			_, _ = io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		_, _ = fmt.Fprintf(out, "\t%s\n", msg)
	}
}
