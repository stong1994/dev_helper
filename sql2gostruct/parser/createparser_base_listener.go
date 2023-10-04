// Code generated from CreateParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // CreateParser

import "github.com/antlr4-go/antlr/v4"

// BaseCreateParserListener is a complete listener for a parse tree produced by CreateParser.
type BaseCreateParserListener struct{}

var _ CreateParserListener = &BaseCreateParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCreateParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCreateParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCreateParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCreateParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseCreateParserListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseCreateParserListener) ExitStat(ctx *StatContext) {}

// EnterCreate_table is called when production create_table is entered.
func (s *BaseCreateParserListener) EnterCreate_table(ctx *Create_tableContext) {}

// ExitCreate_table is called when production create_table is exited.
func (s *BaseCreateParserListener) ExitCreate_table(ctx *Create_tableContext) {}

// EnterCreate_definition is called when production create_definition is entered.
func (s *BaseCreateParserListener) EnterCreate_definition(ctx *Create_definitionContext) {}

// ExitCreate_definition is called when production create_definition is exited.
func (s *BaseCreateParserListener) ExitCreate_definition(ctx *Create_definitionContext) {}

// EnterColumn_definition is called when production column_definition is entered.
func (s *BaseCreateParserListener) EnterColumn_definition(ctx *Column_definitionContext) {}

// ExitColumn_definition is called when production column_definition is exited.
func (s *BaseCreateParserListener) ExitColumn_definition(ctx *Column_definitionContext) {}

// EnterComment_defi is called when production comment_defi is entered.
func (s *BaseCreateParserListener) EnterComment_defi(ctx *Comment_defiContext) {}

// ExitComment_defi is called when production comment_defi is exited.
func (s *BaseCreateParserListener) ExitComment_defi(ctx *Comment_defiContext) {}

// EnterKey_defi is called when production key_defi is entered.
func (s *BaseCreateParserListener) EnterKey_defi(ctx *Key_defiContext) {}

// ExitKey_defi is called when production key_defi is exited.
func (s *BaseCreateParserListener) ExitKey_defi(ctx *Key_defiContext) {}

// EnterShow_length is called when production show_length is entered.
func (s *BaseCreateParserListener) EnterShow_length(ctx *Show_lengthContext) {}

// ExitShow_length is called when production show_length is exited.
func (s *BaseCreateParserListener) ExitShow_length(ctx *Show_lengthContext) {}

// EnterData_type is called when production data_type is entered.
func (s *BaseCreateParserListener) EnterData_type(ctx *Data_typeContext) {}

// ExitData_type is called when production data_type is exited.
func (s *BaseCreateParserListener) ExitData_type(ctx *Data_typeContext) {}

// EnterKey_part is called when production key_part is entered.
func (s *BaseCreateParserListener) EnterKey_part(ctx *Key_partContext) {}

// ExitKey_part is called when production key_part is exited.
func (s *BaseCreateParserListener) ExitKey_part(ctx *Key_partContext) {}

// EnterIndex_type is called when production index_type is entered.
func (s *BaseCreateParserListener) EnterIndex_type(ctx *Index_typeContext) {}

// ExitIndex_type is called when production index_type is exited.
func (s *BaseCreateParserListener) ExitIndex_type(ctx *Index_typeContext) {}

// EnterIndex_option is called when production index_option is entered.
func (s *BaseCreateParserListener) EnterIndex_option(ctx *Index_optionContext) {}

// ExitIndex_option is called when production index_option is exited.
func (s *BaseCreateParserListener) ExitIndex_option(ctx *Index_optionContext) {}

// EnterCheck_constraint_definition is called when production check_constraint_definition is entered.
func (s *BaseCreateParserListener) EnterCheck_constraint_definition(ctx *Check_constraint_definitionContext) {
}

// ExitCheck_constraint_definition is called when production check_constraint_definition is exited.
func (s *BaseCreateParserListener) ExitCheck_constraint_definition(ctx *Check_constraint_definitionContext) {
}

// EnterReference_definition is called when production reference_definition is entered.
func (s *BaseCreateParserListener) EnterReference_definition(ctx *Reference_definitionContext) {}

// ExitReference_definition is called when production reference_definition is exited.
func (s *BaseCreateParserListener) ExitReference_definition(ctx *Reference_definitionContext) {}

// EnterReference_option is called when production reference_option is entered.
func (s *BaseCreateParserListener) EnterReference_option(ctx *Reference_optionContext) {}

// ExitReference_option is called when production reference_option is exited.
func (s *BaseCreateParserListener) ExitReference_option(ctx *Reference_optionContext) {}

// EnterTable_options is called when production table_options is entered.
func (s *BaseCreateParserListener) EnterTable_options(ctx *Table_optionsContext) {}

// ExitTable_options is called when production table_options is exited.
func (s *BaseCreateParserListener) ExitTable_options(ctx *Table_optionsContext) {}

// EnterTable_option is called when production table_option is entered.
func (s *BaseCreateParserListener) EnterTable_option(ctx *Table_optionContext) {}

// ExitTable_option is called when production table_option is exited.
func (s *BaseCreateParserListener) ExitTable_option(ctx *Table_optionContext) {}

// EnterPartition_options is called when production partition_options is entered.
func (s *BaseCreateParserListener) EnterPartition_options(ctx *Partition_optionsContext) {}

// ExitPartition_options is called when production partition_options is exited.
func (s *BaseCreateParserListener) ExitPartition_options(ctx *Partition_optionsContext) {}

// EnterPartition_definition is called when production partition_definition is entered.
func (s *BaseCreateParserListener) EnterPartition_definition(ctx *Partition_definitionContext) {}

// ExitPartition_definition is called when production partition_definition is exited.
func (s *BaseCreateParserListener) ExitPartition_definition(ctx *Partition_definitionContext) {}

// EnterSubpartition_definition is called when production subpartition_definition is entered.
func (s *BaseCreateParserListener) EnterSubpartition_definition(ctx *Subpartition_definitionContext) {
}

// ExitSubpartition_definition is called when production subpartition_definition is exited.
func (s *BaseCreateParserListener) ExitSubpartition_definition(ctx *Subpartition_definitionContext) {}

// EnterTablespace_option is called when production tablespace_option is entered.
func (s *BaseCreateParserListener) EnterTablespace_option(ctx *Tablespace_optionContext) {}

// ExitTablespace_option is called when production tablespace_option is exited.
func (s *BaseCreateParserListener) ExitTablespace_option(ctx *Tablespace_optionContext) {}

// EnterTbl_name is called when production tbl_name is entered.
func (s *BaseCreateParserListener) EnterTbl_name(ctx *Tbl_nameContext) {}

// ExitTbl_name is called when production tbl_name is exited.
func (s *BaseCreateParserListener) ExitTbl_name(ctx *Tbl_nameContext) {}

// EnterCol_name is called when production col_name is entered.
func (s *BaseCreateParserListener) EnterCol_name(ctx *Col_nameContext) {}

// ExitCol_name is called when production col_name is exited.
func (s *BaseCreateParserListener) ExitCol_name(ctx *Col_nameContext) {}

// EnterTablespace_name is called when production tablespace_name is entered.
func (s *BaseCreateParserListener) EnterTablespace_name(ctx *Tablespace_nameContext) {}

// ExitTablespace_name is called when production tablespace_name is exited.
func (s *BaseCreateParserListener) ExitTablespace_name(ctx *Tablespace_nameContext) {}

// EnterIndex_name is called when production index_name is entered.
func (s *BaseCreateParserListener) EnterIndex_name(ctx *Index_nameContext) {}

// ExitIndex_name is called when production index_name is exited.
func (s *BaseCreateParserListener) ExitIndex_name(ctx *Index_nameContext) {}

// EnterEngine_name is called when production engine_name is entered.
func (s *BaseCreateParserListener) EnterEngine_name(ctx *Engine_nameContext) {}

// ExitEngine_name is called when production engine_name is exited.
func (s *BaseCreateParserListener) ExitEngine_name(ctx *Engine_nameContext) {}

// EnterSymbol is called when production symbol is entered.
func (s *BaseCreateParserListener) EnterSymbol(ctx *SymbolContext) {}

// ExitSymbol is called when production symbol is exited.
func (s *BaseCreateParserListener) ExitSymbol(ctx *SymbolContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseCreateParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseCreateParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterString is called when production string is entered.
func (s *BaseCreateParserListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseCreateParserListener) ExitString(ctx *StringContext) {}

// EnterCollation_name is called when production collation_name is entered.
func (s *BaseCreateParserListener) EnterCollation_name(ctx *Collation_nameContext) {}

// ExitCollation_name is called when production collation_name is exited.
func (s *BaseCreateParserListener) ExitCollation_name(ctx *Collation_nameContext) {}

// EnterParser_name is called when production parser_name is entered.
func (s *BaseCreateParserListener) EnterParser_name(ctx *Parser_nameContext) {}

// ExitParser_name is called when production parser_name is exited.
func (s *BaseCreateParserListener) ExitParser_name(ctx *Parser_nameContext) {}

// EnterConnect_string is called when production connect_string is entered.
func (s *BaseCreateParserListener) EnterConnect_string(ctx *Connect_stringContext) {}

// ExitConnect_string is called when production connect_string is exited.
func (s *BaseCreateParserListener) ExitConnect_string(ctx *Connect_stringContext) {}

// EnterCharset_name is called when production charset_name is entered.
func (s *BaseCreateParserListener) EnterCharset_name(ctx *Charset_nameContext) {}

// ExitCharset_name is called when production charset_name is exited.
func (s *BaseCreateParserListener) ExitCharset_name(ctx *Charset_nameContext) {}

// EnterValue is called when production value is entered.
func (s *BaseCreateParserListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseCreateParserListener) ExitValue(ctx *ValueContext) {}

// EnterLength is called when production length is entered.
func (s *BaseCreateParserListener) EnterLength(ctx *LengthContext) {}

// ExitLength is called when production length is exited.
func (s *BaseCreateParserListener) ExitLength(ctx *LengthContext) {}

// EnterAbsolute_path_to_directory is called when production absolute_path_to_directory is entered.
func (s *BaseCreateParserListener) EnterAbsolute_path_to_directory(ctx *Absolute_path_to_directoryContext) {
}

// ExitAbsolute_path_to_directory is called when production absolute_path_to_directory is exited.
func (s *BaseCreateParserListener) ExitAbsolute_path_to_directory(ctx *Absolute_path_to_directoryContext) {
}

// EnterExpr is called when production expr is entered.
func (s *BaseCreateParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseCreateParserListener) ExitExpr(ctx *ExprContext) {}

// EnterNum is called when production num is entered.
func (s *BaseCreateParserListener) EnterNum(ctx *NumContext) {}

// ExitNum is called when production num is exited.
func (s *BaseCreateParserListener) ExitNum(ctx *NumContext) {}

// EnterMax_number_of_rows is called when production max_number_of_rows is entered.
func (s *BaseCreateParserListener) EnterMax_number_of_rows(ctx *Max_number_of_rowsContext) {}

// ExitMax_number_of_rows is called when production max_number_of_rows is exited.
func (s *BaseCreateParserListener) ExitMax_number_of_rows(ctx *Max_number_of_rowsContext) {}

// EnterMin_number_of_rows is called when production min_number_of_rows is entered.
func (s *BaseCreateParserListener) EnterMin_number_of_rows(ctx *Min_number_of_rowsContext) {}

// ExitMin_number_of_rows is called when production min_number_of_rows is exited.
func (s *BaseCreateParserListener) ExitMin_number_of_rows(ctx *Min_number_of_rowsContext) {}

// EnterPartition_name is called when production partition_name is entered.
func (s *BaseCreateParserListener) EnterPartition_name(ctx *Partition_nameContext) {}

// ExitPartition_name is called when production partition_name is exited.
func (s *BaseCreateParserListener) ExitPartition_name(ctx *Partition_nameContext) {}

// EnterLogical_name is called when production logical_name is entered.
func (s *BaseCreateParserListener) EnterLogical_name(ctx *Logical_nameContext) {}

// ExitLogical_name is called when production logical_name is exited.
func (s *BaseCreateParserListener) ExitLogical_name(ctx *Logical_nameContext) {}

// EnterData_dir is called when production data_dir is entered.
func (s *BaseCreateParserListener) EnterData_dir(ctx *Data_dirContext) {}

// ExitData_dir is called when production data_dir is exited.
func (s *BaseCreateParserListener) ExitData_dir(ctx *Data_dirContext) {}

// EnterIndex_dir is called when production index_dir is entered.
func (s *BaseCreateParserListener) EnterIndex_dir(ctx *Index_dirContext) {}

// ExitIndex_dir is called when production index_dir is exited.
func (s *BaseCreateParserListener) ExitIndex_dir(ctx *Index_dirContext) {}

// EnterColumn_list is called when production column_list is entered.
func (s *BaseCreateParserListener) EnterColumn_list(ctx *Column_listContext) {}

// ExitColumn_list is called when production column_list is exited.
func (s *BaseCreateParserListener) ExitColumn_list(ctx *Column_listContext) {}

// EnterValue_list is called when production value_list is entered.
func (s *BaseCreateParserListener) EnterValue_list(ctx *Value_listContext) {}

// ExitValue_list is called when production value_list is exited.
func (s *BaseCreateParserListener) ExitValue_list(ctx *Value_listContext) {}
