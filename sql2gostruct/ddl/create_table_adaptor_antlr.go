package ddl

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sql2gostruct/parser"
)

type CrateTableAdaptorAntlr struct{}

func (c CrateTableAdaptorAntlr) Parse(ddl string) (CreateDDLData, error) {
	lexer := parser.NewMySqlLexer(antlr.NewInputStream(ddl))

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	fmt.Println(stream.GetAllText())
	createParser := parser.NewMySqlParser(stream)

	listener := NewCreateDDLListener()
	antlr.NewParseTreeWalker().Walk(listener, createParser.Root())
	return listener.Data, nil
}
