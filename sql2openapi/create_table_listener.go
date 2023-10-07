package main

import (
	"sql2openapi/consts"
	"sql2openapi/parser"
	"strings"
)

type Column struct {
	Name    string
	Type    consts.ColumnType
	Comment string
}

func (c *Column) IsValid() bool {
	return c.Name != "" && c.Type != consts.Unknown
}

func (c *Column) Reset() {
	c.Name = ""
	c.Type = consts.Unknown
	c.Comment = ""
}

type CreateDDLData struct {
	TableName string
	Columns   []Column
	Comment   string
}

type CreateTableListener struct {
	*parser.BaseMySqlParserListener
	Data      []CreateDDLData
	curTable  *CreateDDLData
	curColumn *Column
}

func NewCreateTableVisitor() *CreateTableListener {
	return &CreateTableListener{
		BaseMySqlParserListener: &parser.BaseMySqlParserListener{},
	}
}

func (ctl *CreateTableListener) EnterColumnCreateTable(c *parser.ColumnCreateTableContext) {
	tableName := c.TableName().GetText()
	strs := strings.Split(tableName, ".")
	if len(strs) == 2 {
		tableName = strs[1]
	}
	if hasQuoted(tableName) {
		tableName = tableName[1 : len(tableName)-1]
	}
	ctl.curTable = &CreateDDLData{TableName: tableName}
}

func (ctl *CreateTableListener) ExitColumnCreateTable(c *parser.ColumnCreateTableContext) {
	if ctl.curTableValid() {
		ctl.Data = append(ctl.Data, *ctl.curTable)
	}
}

func (ctl *CreateTableListener) ExitColumnDeclaration(c *parser.ColumnDeclarationContext) {
	if ctl.curColumn.IsValid() {
		ctl.curTable.Columns = append(ctl.curTable.Columns, *ctl.curColumn)
	}
}

func (ctl *CreateTableListener) EnterColumnDataType(c *parser.ColumnDataTypeContext) {
	ctl.curColumn.Type = consts.ColumnTypeMap[strings.ToUpper(c.GetTypeName().GetText())]
}

func (ctl *CreateTableListener) EnterCommentColumnConstraint(c *parser.CommentColumnConstraintContext) {
	ctl.curColumn.Comment = c.GetComment().GetText()
}

func (ctl *CreateTableListener) EnterColumnDeclaration(c *parser.ColumnDeclarationContext) {
	ctl.resetColumn()

	name := c.FullColumnName().GetText()
	if hasQuoted(name) {
		name = name[1 : len(name)-1]
	}
	ctl.curColumn.Name = name
}

func (ctl *CreateTableListener) EnterTableOptionComment(c *parser.TableOptionCommentContext) {
	comment := c.STRING_LITERAL().GetText()
	if hasQuoted(comment) {
		comment = comment[1 : len(comment)-1]
	}
	ctl.curTable.Comment = comment
}

func (ctl *CreateTableListener) resetColumn() {
	if ctl.curColumn == nil {
		ctl.curColumn = new(Column)
	} else {
		ctl.curColumn.Reset()
	}
}

func (ctl *CreateTableListener) curTableValid() bool {
	if ctl.curTable == nil || ctl.curTable.TableName == "" || len(ctl.curTable.Columns) == 0 {
		return false
	}
	return true
}

func hasQuoted(str string) bool {
	return strings.HasPrefix(str, "'") || strings.HasPrefix(str, "\"") || strings.HasPrefix(str, "`")
}
