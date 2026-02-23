package repl

import (
	"bufio"
	"fmt"
	"freedom/evaluator"
	"freedom/lexer"
	"freedom/parser"
	"io"
)

const PROMPT = ">>"

const MONKEY_FACE = `
.--.
.-"
/ .. \/
"-.
.-. .-.
Y
__,__
.--.
\/ .. \
| |'|/\|'| |
| \\\ 0 | 0 /// |
\ '- ,\.-"""""""-./, -' /
''-' /_
^ ^
_\ '-''
|\.__./|
\\ '~' //
'._ '-=-' _.'
'-----'
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
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
		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, "Evaluated\n")
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		} else {
			io.WriteString(out, "Parsed\n")
			io.WriteString(out, program.String())
			io.WriteString(out, "\n")
		}

	}
}
func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
