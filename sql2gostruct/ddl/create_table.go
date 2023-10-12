package ddl

import (
	"fmt"
	"sql2gostruct/consts"
	"strconv"
	"strings"
	"unicode"
)

const SpaceCharacter = "    "

type Column struct {
	Name    string
	Type    consts.ColumnType
	Comment string
}

type CreateDDLData struct {
	TableName string
	DBNAME    string
	Columns   []Column
}

type CreateTableAdaptor interface {
	Parse(ddl string) (CreateDDLData, error)
}

type CreateTableParser struct {
	ddl    string
	parser CreateTableAdaptor
}

func NewCreateTableParser(ddl string, parser CreateTableAdaptor) CreateTableParser {
	return CreateTableParser{
		ddl:    ddl,
		parser: parser,
	}
}

func (ctp CreateTableParser) Parse() (string, error) {
	data, err := ctp.parser.Parse(ctp.ddl)
	if err != nil {
		return "", err
	}
	tableName := getTableStructName(data.TableName)
	columnStr := getColumns(data.Columns)
	return fmt.Sprintf(`type %s struct {
%s
}`, tableName, columnStr), nil
}

func getTableStructName(name string) string {
	if strings.HasPrefix(name, "t_") {
		name = strings.Trim(name, "t_")
	}
	return convertToCamelCase(name)
}

func convertToCamelCase(word string) string {
	if !strings.Contains(word, "_") {
		var s []rune
		for _, c := range word {
			if len(s) == 0 {
				s = append(s, unicode.ToUpper(c))
			} else {
				s = append(s, c)
			}
		}
		return string(s)
	}
	words := strings.Split(word, "_")
	var rst string
	for _, w := range words {
		rst += convertToCamelCase(w)
	}
	return rst
}

func getColumns(columns []Column) string {
	var lines []string
	columnWidths := [4]int{}
	for _, column := range columns {
		items := getColumn(column)
		if len(items) != 4 {
			panic("some code changed?")
		}
		for i, v := range items {
			if len(v) > columnWidths[i] {
				columnWidths[i] = len(v)
			}
		}
	}
	for _, column := range columns {
		items := getColumn(column)
		line := SpaceCharacter
		for i, v := range items {
			padding := strings.Repeat(" ", columnWidths[i]-len(v))
			line += v + padding + SpaceCharacter
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

// 返回值eg: ID    string    gorm:"column:c_id"
func getColumn(column Column) [4]string {
	return [4]string{parseColumnName(column.Name), parseMysqlTypeCus(column), getDBFieldName(column.Name), getFieldComment(column.Comment)}
}

func getFieldComment(comment string) string {
	if comment == "" {
		return ""
	}
	return "// " + comment
}

func parseColumnName(name string) string {
	if strings.HasPrefix(name, "c_") {
		name = strings.TrimPrefix(name, "c_")
	}
	name = strings.ReplaceAll(name, "id", "ID")
	return convertToCamelCase(name)
}

func parseMysqlTypeCus(col Column) string {
	if strings.HasPrefix(parseColumnName(col.Name), "is_") {
		return "bool"
	}
	return parseMysqlType(col.Type)
}

// https://zontroy.com/mysql-to-go-type-mapping/
func parseMysqlType(typ consts.ColumnType) string {
	switch typ {
	case consts.TINYINT:
		return "int8"
	case consts.SMALLINT:
		return "int16"
	case consts.MEDIUMINT:
		return "int"
	case consts.MIDDLEINT:
		return "int"
	case consts.INT:
		return "int"
	case consts.INT1:
		return "int"
	case consts.INT2:
		return "int"
	case consts.INT3:
		return "int"
	case consts.INT4:
		return "int"
	case consts.INT8:
		return "int"
	case consts.INTEGER:
		return "int"
	case consts.BIGINT:
		return "int64"
	case consts.REAL:
		return "float64"
	case consts.DOUBLE:
		return "float64"
	case consts.PRECISION:
		return "float64"
	case consts.FLOAT:
		return "float64"
	case consts.FLOAT4:
		return "float64"
	case consts.FLOAT8:
		return "float64"
	case consts.DECIMAL:
		return "float64"
	case consts.DEC:
		return "float64"
	case consts.NUMERIC:
		return "float64"
	case consts.DATE:
		return ""
	case consts.TIME:
		return "time.Time"
	case consts.TIMESTAMP:
		return "time.Time"
	case consts.DATETIME:
		return "time.Time"
	case consts.YEAR:
		return "int"
	case consts.CHAR:
		return "string"
	case consts.VARCHAR:
		return "string"
	case consts.NVARCHAR:
		return "string"
	case consts.NATIONAL:
		return "string"
	case consts.BINARY:
		return "string"
	case consts.VARBINARY:
		return "string"
	case consts.TINYBLOB:
		return "string"
	case consts.BLOB:
		return "string"
	case consts.MEDIUMBLOB:
		return "string"
	case consts.LONG:
		return "string"
	case consts.LONGBLOB:
		return "string"
	case consts.TINYTEXT:
		return "string"
	case consts.TEXT:
		return "string"
	case consts.MEDIUMTEXT:
		return "string"
	case consts.LONGTEXT:
		return "string"
	case consts.ENUM:
		return "int"
	case consts.VARYING:
		return "string"
	case consts.SERIAL:
		return "string"
	}
	panic("unknown type:" + strconv.Itoa(int(typ)))
}

func getDBFieldName(field string) string {
	return fmt.Sprintf("`gorm:\"column:%s\"`", field)
}
