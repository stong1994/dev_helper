// Code generated from CreateParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // CreateParser

import "github.com/antlr4-go/antlr/v4"

// CreateParserListener is a complete listener for a parse tree produced by CreateParser.
type CreateParserListener interface {
	antlr.ParseTreeListener

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterCreate_table is called when entering the create_table production.
	EnterCreate_table(c *Create_tableContext)

	// EnterCreate_definition is called when entering the create_definition production.
	EnterCreate_definition(c *Create_definitionContext)

	// EnterColumn_definition is called when entering the column_definition production.
	EnterColumn_definition(c *Column_definitionContext)

	// EnterComment_defi is called when entering the comment_defi production.
	EnterComment_defi(c *Comment_defiContext)

	// EnterKey_defi is called when entering the key_defi production.
	EnterKey_defi(c *Key_defiContext)

	// EnterShow_length is called when entering the show_length production.
	EnterShow_length(c *Show_lengthContext)

	// EnterData_type is called when entering the data_type production.
	EnterData_type(c *Data_typeContext)

	// EnterKey_part is called when entering the key_part production.
	EnterKey_part(c *Key_partContext)

	// EnterIndex_type is called when entering the index_type production.
	EnterIndex_type(c *Index_typeContext)

	// EnterIndex_option is called when entering the index_option production.
	EnterIndex_option(c *Index_optionContext)

	// EnterCheck_constraint_definition is called when entering the check_constraint_definition production.
	EnterCheck_constraint_definition(c *Check_constraint_definitionContext)

	// EnterReference_definition is called when entering the reference_definition production.
	EnterReference_definition(c *Reference_definitionContext)

	// EnterReference_option is called when entering the reference_option production.
	EnterReference_option(c *Reference_optionContext)

	// EnterTable_options is called when entering the table_options production.
	EnterTable_options(c *Table_optionsContext)

	// EnterTable_option is called when entering the table_option production.
	EnterTable_option(c *Table_optionContext)

	// EnterPartition_options is called when entering the partition_options production.
	EnterPartition_options(c *Partition_optionsContext)

	// EnterPartition_definition is called when entering the partition_definition production.
	EnterPartition_definition(c *Partition_definitionContext)

	// EnterSubpartition_definition is called when entering the subpartition_definition production.
	EnterSubpartition_definition(c *Subpartition_definitionContext)

	// EnterTablespace_option is called when entering the tablespace_option production.
	EnterTablespace_option(c *Tablespace_optionContext)

	// EnterTbl_name is called when entering the tbl_name production.
	EnterTbl_name(c *Tbl_nameContext)

	// EnterCol_name is called when entering the col_name production.
	EnterCol_name(c *Col_nameContext)

	// EnterTablespace_name is called when entering the tablespace_name production.
	EnterTablespace_name(c *Tablespace_nameContext)

	// EnterIndex_name is called when entering the index_name production.
	EnterIndex_name(c *Index_nameContext)

	// EnterEngine_name is called when entering the engine_name production.
	EnterEngine_name(c *Engine_nameContext)

	// EnterSymbol is called when entering the symbol production.
	EnterSymbol(c *SymbolContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterCollation_name is called when entering the collation_name production.
	EnterCollation_name(c *Collation_nameContext)

	// EnterParser_name is called when entering the parser_name production.
	EnterParser_name(c *Parser_nameContext)

	// EnterConnect_string is called when entering the connect_string production.
	EnterConnect_string(c *Connect_stringContext)

	// EnterCharset_name is called when entering the charset_name production.
	EnterCharset_name(c *Charset_nameContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterLength is called when entering the length production.
	EnterLength(c *LengthContext)

	// EnterAbsolute_path_to_directory is called when entering the absolute_path_to_directory production.
	EnterAbsolute_path_to_directory(c *Absolute_path_to_directoryContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterNum is called when entering the num production.
	EnterNum(c *NumContext)

	// EnterMax_number_of_rows is called when entering the max_number_of_rows production.
	EnterMax_number_of_rows(c *Max_number_of_rowsContext)

	// EnterMin_number_of_rows is called when entering the min_number_of_rows production.
	EnterMin_number_of_rows(c *Min_number_of_rowsContext)

	// EnterPartition_name is called when entering the partition_name production.
	EnterPartition_name(c *Partition_nameContext)

	// EnterLogical_name is called when entering the logical_name production.
	EnterLogical_name(c *Logical_nameContext)

	// EnterData_dir is called when entering the data_dir production.
	EnterData_dir(c *Data_dirContext)

	// EnterIndex_dir is called when entering the index_dir production.
	EnterIndex_dir(c *Index_dirContext)

	// EnterColumn_list is called when entering the column_list production.
	EnterColumn_list(c *Column_listContext)

	// EnterValue_list is called when entering the value_list production.
	EnterValue_list(c *Value_listContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitCreate_table is called when exiting the create_table production.
	ExitCreate_table(c *Create_tableContext)

	// ExitCreate_definition is called when exiting the create_definition production.
	ExitCreate_definition(c *Create_definitionContext)

	// ExitColumn_definition is called when exiting the column_definition production.
	ExitColumn_definition(c *Column_definitionContext)

	// ExitComment_defi is called when exiting the comment_defi production.
	ExitComment_defi(c *Comment_defiContext)

	// ExitKey_defi is called when exiting the key_defi production.
	ExitKey_defi(c *Key_defiContext)

	// ExitShow_length is called when exiting the show_length production.
	ExitShow_length(c *Show_lengthContext)

	// ExitData_type is called when exiting the data_type production.
	ExitData_type(c *Data_typeContext)

	// ExitKey_part is called when exiting the key_part production.
	ExitKey_part(c *Key_partContext)

	// ExitIndex_type is called when exiting the index_type production.
	ExitIndex_type(c *Index_typeContext)

	// ExitIndex_option is called when exiting the index_option production.
	ExitIndex_option(c *Index_optionContext)

	// ExitCheck_constraint_definition is called when exiting the check_constraint_definition production.
	ExitCheck_constraint_definition(c *Check_constraint_definitionContext)

	// ExitReference_definition is called when exiting the reference_definition production.
	ExitReference_definition(c *Reference_definitionContext)

	// ExitReference_option is called when exiting the reference_option production.
	ExitReference_option(c *Reference_optionContext)

	// ExitTable_options is called when exiting the table_options production.
	ExitTable_options(c *Table_optionsContext)

	// ExitTable_option is called when exiting the table_option production.
	ExitTable_option(c *Table_optionContext)

	// ExitPartition_options is called when exiting the partition_options production.
	ExitPartition_options(c *Partition_optionsContext)

	// ExitPartition_definition is called when exiting the partition_definition production.
	ExitPartition_definition(c *Partition_definitionContext)

	// ExitSubpartition_definition is called when exiting the subpartition_definition production.
	ExitSubpartition_definition(c *Subpartition_definitionContext)

	// ExitTablespace_option is called when exiting the tablespace_option production.
	ExitTablespace_option(c *Tablespace_optionContext)

	// ExitTbl_name is called when exiting the tbl_name production.
	ExitTbl_name(c *Tbl_nameContext)

	// ExitCol_name is called when exiting the col_name production.
	ExitCol_name(c *Col_nameContext)

	// ExitTablespace_name is called when exiting the tablespace_name production.
	ExitTablespace_name(c *Tablespace_nameContext)

	// ExitIndex_name is called when exiting the index_name production.
	ExitIndex_name(c *Index_nameContext)

	// ExitEngine_name is called when exiting the engine_name production.
	ExitEngine_name(c *Engine_nameContext)

	// ExitSymbol is called when exiting the symbol production.
	ExitSymbol(c *SymbolContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitCollation_name is called when exiting the collation_name production.
	ExitCollation_name(c *Collation_nameContext)

	// ExitParser_name is called when exiting the parser_name production.
	ExitParser_name(c *Parser_nameContext)

	// ExitConnect_string is called when exiting the connect_string production.
	ExitConnect_string(c *Connect_stringContext)

	// ExitCharset_name is called when exiting the charset_name production.
	ExitCharset_name(c *Charset_nameContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitLength is called when exiting the length production.
	ExitLength(c *LengthContext)

	// ExitAbsolute_path_to_directory is called when exiting the absolute_path_to_directory production.
	ExitAbsolute_path_to_directory(c *Absolute_path_to_directoryContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitNum is called when exiting the num production.
	ExitNum(c *NumContext)

	// ExitMax_number_of_rows is called when exiting the max_number_of_rows production.
	ExitMax_number_of_rows(c *Max_number_of_rowsContext)

	// ExitMin_number_of_rows is called when exiting the min_number_of_rows production.
	ExitMin_number_of_rows(c *Min_number_of_rowsContext)

	// ExitPartition_name is called when exiting the partition_name production.
	ExitPartition_name(c *Partition_nameContext)

	// ExitLogical_name is called when exiting the logical_name production.
	ExitLogical_name(c *Logical_nameContext)

	// ExitData_dir is called when exiting the data_dir production.
	ExitData_dir(c *Data_dirContext)

	// ExitIndex_dir is called when exiting the index_dir production.
	ExitIndex_dir(c *Index_dirContext)

	// ExitColumn_list is called when exiting the column_list production.
	ExitColumn_list(c *Column_listContext)

	// ExitValue_list is called when exiting the value_list production.
	ExitValue_list(c *Value_listContext)
}
