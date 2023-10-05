package ddl

func CreateTableMethod(ddl string) (string, error) {
	rst, err := NewCreateTableParser(ddl, CrateTableAdaptorAntlr{}).Parse()
	return rst, err
}
