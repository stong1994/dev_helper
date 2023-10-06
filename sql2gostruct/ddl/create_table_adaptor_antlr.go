package ddl

import (
	"github.com/antlr4-go/antlr/v4"
	"sql2gostruct/parser"
)

type CrateTableAdaptorAntlr struct{}

func (c CrateTableAdaptorAntlr) Parse(ddl string) (CreateDDLData, error) {
	lexer := parser.NewMySqlLexer(antlr.NewInputStream(ddl))

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	createParser := parser.NewMySqlParser(stream)

	errListener := &ErrListener{}
	createParser.RemoveErrorListeners() // 默认会使用ConsoleErrorListener，需要移除。
	createParser.AddErrorListener(errListener)
	createParser.GetInterpreter().SetPredictionMode(antlr.PredictionModeLLExactAmbigDetection)

	listener := NewCreateDDLListener()
	antlr.NewParseTreeWalker().Walk(listener, createParser.Root())

	return listener.Data, errListener.Error()
}
