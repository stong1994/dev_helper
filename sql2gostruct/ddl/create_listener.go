package ddl

import (
	"sql2gostruct/consts"
	"sql2gostruct/parser"
	"strings"
)

type CreateDDLListener struct {
	*parser.BaseCreateParserListener
	Data CreateDDLData
}

func NewCreateDDLListener() *CreateDDLListener {
	return &CreateDDLListener{
		BaseCreateParserListener: &parser.BaseCreateParserListener{},
	}
}

func (cdl *CreateDDLListener) EnterCreate_table(c *parser.Create_tableContext) {
	cdl.Data.TableName = c.GetText()
}

func (cdl *CreateDDLListener) EnterTbl_name(c *parser.Tbl_nameContext) {
	tableName := c.GetText()
	strs := strings.Split(tableName, ".")
	if len(strs) == 2 {
		tableName = strs[1]
		dbName := strs[0]
		if strings.HasPrefix(dbName, "`") {
			dbName = dbName[1 : len(dbName)-1]
		}
		cdl.Data.DBNAME = dbName
	}
	if strings.HasPrefix(tableName, "`") {
		tableName = tableName[1 : len(tableName)-1]
	}
	cdl.Data.TableName = tableName
}

func (cdl *CreateDDLListener) EnterCol_name(c *parser.Col_nameContext) {
	cdl.Data.Columns = append(cdl.Data.Columns, Column{Name: c.GetText()})
}

func (cdl *CreateDDLListener) EnterData_type(c *parser.Data_typeContext) {
	l := len(cdl.Data.Columns)
	target := cdl.Data.Columns[l-1]
	target.Type = consts.ColumnTypeMap[strings.ToUpper(c.GetText())]
	cdl.Data.Columns[l-1] = target
}

func (cdl CreateDDLListener) EnterComment_defi(c *parser.Comment_defiContext) {
	l := len(cdl.Data.Columns)
	target := cdl.Data.Columns[l-1]
	target.Comment = c.GetCol_comment().GetText()
	if strings.HasPrefix(target.Comment, "'") || strings.HasPrefix(target.Comment, `""`) {
		target.Comment = target.Comment[1 : len(target.Comment)-1]
	}
	cdl.Data.Columns[l-1] = target
}
