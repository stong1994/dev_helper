package ddl

import (
	"sql2gostruct/consts"
	"sql2gostruct/parser"
	"strings"
)

type CreateDDLListener struct {
	*parser.BaseMySqlParserListener
	Data CreateDDLData
}

func NewCreateDDLListener() *CreateDDLListener {
	return &CreateDDLListener{
		BaseMySqlParserListener: &parser.BaseMySqlParserListener{},
	}
}

func (cdl *CreateDDLListener) EnterTableName(c *parser.TableNameContext) {
	tableName := c.GetText()
	strs := strings.Split(tableName, ".")
	if len(strs) == 2 {
		tableName = strs[1]
		dbName := strs[0]
		if hasQuoted(dbName) {
			dbName = dbName[1 : len(dbName)-1]
		}
		cdl.Data.DBNAME = dbName
	}
	if hasQuoted(tableName) {
		tableName = tableName[1 : len(tableName)-1]
	}
	cdl.Data.TableName = tableName
}

func (cdl *CreateDDLListener) EnterFullColumnName(c *parser.FullColumnNameContext) {
	cdl.Data.Columns = append(cdl.Data.Columns, Column{Name: c.GetText()})
}

func (cdl *CreateDDLListener) EnterColumnDataType(c *parser.ColumnDataTypeContext) {
	l := len(cdl.Data.Columns)
	target := cdl.Data.Columns[l-1]
	target.Type = consts.ColumnTypeMap[strings.ToUpper(c.GetTypeName().GetText())]
	cdl.Data.Columns[l-1] = target
}

func (cdl CreateDDLListener) EnterCommentColumnConstraint(c *parser.CommentColumnConstraintContext) {
	l := len(cdl.Data.Columns)
	target := cdl.Data.Columns[l-1]
	target.Comment = c.GetComment().GetText()
	if hasQuoted(target.Comment) {
		target.Comment = target.Comment[1 : len(target.Comment)-1]
	}
	cdl.Data.Columns[l-1] = target
}

func hasQuoted(str string) bool {
	return strings.HasPrefix(str, ".") || strings.HasPrefix(str, "\"") || strings.HasPrefix(str, "`")
}
