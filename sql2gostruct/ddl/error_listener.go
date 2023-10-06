package ddl

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"strings"
)

type ErrListener struct {
	antlr.DefaultErrorListener
	errList []string
}

func (el *ErrListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int,
	msg string, e antlr.RecognitionException) {
	el.errList = append(el.errList, fmt.Sprintf("pos: %d:%d, msg: %s", line, column, msg))
}

func (el *ErrListener) String() string {
	return strings.Join(el.errList, ",")
}

func (el *ErrListener) Error() error {
	if len(el.errList) == 0 {
		return nil
	}
	return fmt.Errorf(el.String())
}
