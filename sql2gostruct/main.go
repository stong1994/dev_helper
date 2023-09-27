package main

import (
	"fmt"
	"strings"
	"unicode"
)

const SpaceCharacter = "    "

func ConvertDDL2Struct(ddl string) string {
	tableName := getTableStructName(getTableName(ddl))
	columnStr := getColumns(ddl)
	return fmt.Sprintf(`type %s struct {
%s
}`, tableName, columnStr)
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

func getTableName(ddl string) string {
	ddl = strings.TrimPrefix(ddl, "\n")
	firstLine := strings.ToLower(strings.Split(ddl, "\n")[0])
	firstLine = strings.Join(strings.Fields(firstLine), " ") // 合并空格

	// 去最后一个字符串
	fs := strings.Fields(firstLine)
	tableName := strings.Fields(firstLine)[len(fs)-1]
	if !strings.Contains(tableName, ".") {
		return tableName
	}
	return strings.Split(tableName, ".")[1]
}

func getColumns(ddl string) string {
	columnContent := getColumnContent(ddl)
	columns := strings.Split(columnContent, ",")
	var lines []string
	columnWidths := [3]int{}
	for _, column := range columns {
		items := getColumn(column)
		if len(items) != 3 {
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
func getColumn(column string) [3]string {
	fields := strings.Fields(column)
	if len(fields) < 2 {
		panic("invalid column:" + column)
	}
	return [3]string{parseColumnName(fields[0]), parseMysqlType(fields[1]), getDBFieldName(fields[0])}
}

func parseColumnName(name string) string {
	if strings.HasPrefix(name, "c_") {
		name = strings.TrimPrefix(name, "c_")
	}
	name = strings.ReplaceAll(name, "id", "ID")
	return convertToCamelCase(name)
}

func parseMysqlType(typ string) string {
	if strings.Contains(typ, "char") {
		return "string"
	}
	if strings.Contains(typ, "int") {
		return "int"
	}
	if strings.Contains(typ, "date") || strings.Contains(typ, "time") {
		return "time.Time"
	}
	panic("unknown type:" + typ)
}

func getColumnContent(ddl string) string {
	first, last := -1, -1
	for i, v := range ddl {
		if v == '(' {
			first = i
			break
		}
	}
	for i := len(ddl) - 1; i >= 0; i-- {
		if ddl[i] == ')' {
			last = i
			break
		}
	}
	if first == -1 || last == -1 {
		panic("invalid columns:" + ddl)
	}
	return ddl[first+1 : last]
}

func getDBFieldName(field string) string {
	return fmt.Sprintf(`gorm:"column:%s"`, field)
}
