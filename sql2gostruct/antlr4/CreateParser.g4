parser grammar CreateParser;

options
   { tokenVocab = MySqlLexer;}

stat
   : create_table EOF
   ;

create_table: CREATE (TEMPORARY)? TABLE (IF NOT EXISTS)? tbl_name '(' create_definition (',' create_definition)* ')' table_options? partition_options?;

create_definition: col_name column_definition;
//create_definition:  col_name column_definition
//    | ('INDEX' | 'KEY') (index_name)? (index_type)? '(' key_part (',' key_part)* ')' (index_option)*
//    | ('FULLTEXT' | 'SPATIAL') ('INDEX' | 'KEY') (index_name)? '(' key_part (',' key_part)* ')' (index_option)*
//    | ('CONSTRAINT' (symbol)?)? 'PRIMARY' 'KEY' (index_type)? '(' key_part (',' key_part)* ')' (index_option)*
//    | ('CONSTRAINT' (symbol)?)? 'UNIQUE' ('INDEX' | 'KEY') (index_name)? (index_type)? '(' key_part (',' key_part)* ')' (index_option)*
//    | ('CONSTRAINT' (symbol)?)? 'FOREIGN' 'KEY' (index_name)? '(' col_name (',' col_name)* ')' reference_definition
//    | check_constraint_definition;

column_definition: data_type show_length? (NOT 'NULL' | 'NULL')? (DEFAULT (literal | '(' expr ')'))? ('AUTO_INCREMENT')? ('UNIQUE')? ((key_defi? comment_defi? )| (comment_defi? key_defi?));
//column_definition: data_type (show_length)? (NOT 'NULL' | 'NULL')? ('DEFAULT' (literal | '(' expr ')'))? ('VISIBLE' | 'INVISIBLE')? ('AUTO_INCREMENT')? ('UNIQUE' ('KEY'))? (('PRIMARY')? 'KEY')? ('COMMENT' '\'' string '\'')? ('COLLATE' collation_name)? ('COLUMN_FORMAT' ('FIXED' | 'DYNAMIC' | 'DEFAULT'))? ('ENGINE_ATTRIBUTE' ('=' string))? ('SECONDARY_ENGINE_ATTRIBUTE' ('=' string))? ('STORAGE' ('DISK' | 'MEMORY'))? (reference_definition)? (check_constraint_definition)?
//    | data_type ('COLLATE' collation_name)? ('GENERATED' 'ALWAYS')? 'AS' '(' expr ')' ('COLLATE' | 'STORED')? (NOT 'NULL' | 'NULL')? ('VISIBLE' | 'INVISIBLE')? ('UNIQUE' ('KEY'))? (('PRIMARY')? 'KEY')? ('COMMENT' '\'' string '\'')? (reference_definition)? (check_constraint_definition)?;

comment_defi: COMMENT col_comment=STRING_LITERAL;
key_defi: ('PRIMARY')? 'KEY';
show_length: '(' DIGTS ')';

data_type: TINYINT |
           SMALLINT |
           MEDIUMINT |
           MIDDLEINT |
           INT |
           INT1 |
           INT2 |
           INT3 |
           INT4 |
           INT8 |
           INTEGER |
           BIGINT |
           REAL |
           DOUBLE |
           PRECISION |
           FLOAT |
           FLOAT4 |
           FLOAT8 |
           DECIMAL | DEC | NUMERIC | DATE | TIME | TIMESTAMP | DATETIME | YEAR | CHAR | VARCHAR | NVARCHAR | NATIONAL | BINARY | VARBINARY | TINYBLOB | BLOB | MEDIUMBLOB | LONG | LONGBLOB | TINYTEXT | TEXT | MEDIUMTEXT | LONGTEXT | ENUM | VARYING | SERIAL ;

key_part: (col_name (length)?)? ('ASC' | 'DESC')?;

index_type: 'USING' ('BTREE' | 'HASH');

index_option: ('KEY_BLOCK_SIZE' ('=' value)
    | index_type
    | 'WITH' 'PARSER' parser_name
    | 'COMMENT' '\'' string '\''
    | ('VISIBLE' | 'INVISIBLE')
    | 'ENGINE_ATTRIBUTE' ('=' string)
    | 'SECONDARY_ENGINE_ATTRIBUTE' ('=' string));

check_constraint_definition: ('CONSTRAINT' (symbol)?)? 'CHECK' '(' expr ')' (NOT 'ENFORCED')?;

reference_definition: 'REFERENCES' tbl_name '(' key_part (',' key_part)* ')' (MATCH 'FULL' | MATCH 'PARTIAL' | MATCH 'SIMPLE')? ('ON' 'DELETE' reference_option)? ('ON' 'UPDATE' reference_option)?;

reference_option: 'RESTRICT' | 'CASCADE' | 'SET' 'NULL' | 'NO' 'ACTION' | 'SET' 'DEFAULT';

table_options: table_option (',' table_option)*;

table_option: 'AUTOEXTEND_SIZE' ('=' value)
    | 'AUTO_INCREMENT' ('=' value)
    | 'AVG_ROW_LENGTH' ('=' value)
    | ('DEFAULT')? 'CHARACTER' 'SET' ('=' charset_name)
    | 'CHECKSUM' ('=' ('0' | '1'))
    | ('DEFAULT')? 'COLLATE' ('=' collation_name)
    | COMMENT ('=')? '\'' string '\''
    // | COMPRESSION ('=' ('ZLIB' | 'LZ4' | 'NONE'))
    | 'CONNECTION' ('=' '\'' connect_string '\'')
    | ('DATA' | 'INDEX') 'DIRECTORY' ('=' '\'' absolute_path_to_directory '\'')
    | 'DELAY_KEY_WRITE' ('=' ('0' | '1'))
    // | 'ENCRYPTION' ('=' ('\'Y\'' | '\'N\'' ))
    | 'ENGINE' ('=' engine_name)
    | 'ENGINE_ATTRIBUTE' ('=' '\'' string '\'')
    | 'INSERT_METHOD' ('=' ('NO' | 'FIRST' | 'LAST'))
    | 'KEY_BLOCK_SIZE' ('=' value)
    | 'MAX_ROWS' ('=' value)
    | 'MIN_ROWS' ('=' value)
    | 'PACK_KEYS' ('=' ('0' | '1' | 'DEFAULT'))
    | 'PASSWORD' ('=' '\'' string '\'')
    | 'ROW_FORMAT' ('=' ('DEFAULT' | 'DYNAMIC' | 'FIXED' | 'COMPRESSED' | 'REDUNDANT' | 'COMPACT'))
    | 'START' 'TRANSACTION'
    | 'SECONDARY_ENGINE_ATTRIBUTE' ('=' '\'' string '\'')
    | 'STATS_AUTO_RECALC' ('=' ('DEFAULT' | '0' | '1'))
    | 'STATS_PERSISTENT' ('=' ('DEFAULT' | '0' | '1'))
    | 'STATS_SAMPLE_PAGES' ('=' value)
    | tablespace_option
    | 'UNION' ('=' '(' tbl_name (',' tbl_name)* ')');

partition_options: 'PARTITION' 'BY'
    (('LINEAR')? 'HASH' '(' expr ')'
    | ('LINEAR')? 'KEY' ('ALGORITHM' '=' ('1' | '2'))? '(' column_list ')'
    | 'RANGE' ('(' expr ')')? 'COLUMNS' '(' column_list ')'
    | 'LIST' ('(' expr ')')? 'COLUMNS' '(' column_list ')'
    ) ('PARTITIONS' num)? ('SUBPARTITION' 'BY'
    (('LINEAR')? 'HASH' '(' expr ')'
    | ('LINEAR')? 'KEY' ('ALGORITHM' '=' ('1' | '2'))? '(' column_list ')'
    ) ('SUBPARTITIONS' num)?
    )? ('(' partition_definition (',' partition_definition)* ')')?;

partition_definition: 'PARTITION' partition_name
    ('VALUES'
        ('LESS' 'THAN' (('(' expr ')' | value_list) | 'MAXVALUE'))
    | 'IN' '(' value_list ')'
    )? (('STORAGE')? 'ENGINE' ('=' engine_name))? ('COMMENT' ('=' '\'' string '\''))? ('DATA' 'DIRECTORY' ('=' '\'' data_dir '\''))? ('INDEX' 'DIRECTORY' ('=' '\'' index_dir '\''))? ('MAX_ROWS' ('=' max_number_of_rows))? ('MIN_ROWS' ('=' min_number_of_rows))? ('TABLESPACE' ('=' tablespace_name))? ('(' subpartition_definition (',' subpartition_definition)* ')')?;

subpartition_definition: 'SUBPARTITION' logical_name
    (('STORAGE')? 'ENGINE' ('=' engine_name))? ('COMMENT' ('=' '\'' string '\''))? ('DATA' 'DIRECTORY' ('=' '\'' data_dir '\''))? ('INDEX' 'DIRECTORY' ('=' '\'' index_dir '\''))? ('MAX_ROWS' ('=' max_number_of_rows))? ('MIN_ROWS' ('=' min_number_of_rows))? ('TABLESPACE' ('=' tablespace_name))? ;

tablespace_option: 'TABLESPACE' tablespace_name ('STORAGE' 'DISK')
    | ('TABLESPACE' tablespace_name)? 'STORAGE' 'MEMORY';


tbl_name: ID;
col_name: ID;
tablespace_name:ID;
index_name: ID;
engine_name: 'ARCHIVE' | 'BLACKHOLE' | 'CSV' | 'FEDERATED' | 'INNODB' | 'MEMORY' | 'MRG_MYISAM' | 'MYISAM' | 'NDB' | 'NDBCLUSTER' | 'PERFORMANCE_SCHEMA' | 'TOKUDB';
symbol: ID;
literal: ID;
string: ID;
collation_name: ID;
parser_name: ID;
connect_string: ID;
charset_name: ID;
value: ID;
length: DECIMAL_LITERAL;
absolute_path_to_directory: ID;
expr: ID;
num: DIGTS;
max_number_of_rows: DIGTS;
min_number_of_rows: DIGTS;
partition_name: ID;
logical_name: ID;
data_dir: ID;
index_dir: ID;
column_list: col_name column_definition (',' col_name column_definition)*;
value_list: ID (',' ID)*;