package repl

import (
	"bufio"
	"compiler/compiler"
	"compiler/lexer"
	"compiler/object"
	"compiler/parser"
	"compiler/vm"
	"fmt"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	constants := []object.Object{}
	globals := make([]object.Object, vm.GlobalsSize)
	symbolTable := compiler.NewSymbolTable()

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

		comp := compiler.NewWithState(symbolTable, constants)
		err := comp.Compile(program)
		if err != nil {
			_, _ = fmt.Fprintf(out, "Woops! Compilation failed:\n %s\n", err)
			continue
		}

		code := comp.Bytecode()
		constants = code.Constants

		machine := vm.NewWithGlobalsStore(code, globals)
		err = machine.Run()
		if err != nil {
			_, _ = fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)
			continue
		}

		stackTop := machine.LastPoppedStackElem()
		_, _ = io.WriteString(out, stackTop.Inspect())
		_, _ = io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		_, _ = fmt.Fprintf(out, "\t%s\n", msg)
	}
}
