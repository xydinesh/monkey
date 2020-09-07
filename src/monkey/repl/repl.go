package repl

import (
	"bufio"
	"fmt"
	"monkey/lexer"
	"monkey/token"
	"os"
	"strings"
)

// Start function for using REPL
func Start() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("~> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare("exit", text) == 0 {
			os.Exit(0)
		}
		l := lexer.New(text)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
