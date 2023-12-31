package consts

type ColumnType int

const (
	TINYINT ColumnType = iota + 1
	SMALLINT
	MEDIUMINT
	MIDDLEINT
	INT
	INT1
	INT2
	INT3
	INT4
	INT8
	INTEGER
	BIGINT
	REAL
	DOUBLE
	PRECISION
	FLOAT
	FLOAT4
	FLOAT8
	DECIMAL
	DEC
	NUMERIC
	DATE
	TIME
	TIMESTAMP
	DATETIME
	YEAR
	CHAR
	VARCHAR
	NVARCHAR
	NATIONAL
	BINARY
	VARBINARY
	TINYBLOB
	BLOB
	MEDIUMBLOB
	LONG
	LONGBLOB
	TINYTEXT
	TEXT
	MEDIUMTEXT
	LONGTEXT
	ENUM
	VARYING
	SERIAL
)

var ColumnTypeMap = map[string]ColumnType{
	"TINYINT":    TINYINT,
	"SMALLINT":   SMALLINT,
	"MEDIUMINT":  MEDIUMINT,
	"MIDDLEINT":  MIDDLEINT,
	"INT":        INT,
	"INT1":       INT1,
	"INT2":       INT2,
	"INT3":       INT3,
	"INT4":       INT4,
	"INT8":       INT8,
	"INTEGER":    INTEGER,
	"BIGINT":     BIGINT,
	"REAL":       REAL,
	"DOUBLE":     DOUBLE,
	"PRECISION":  PRECISION,
	"FLOAT":      FLOAT,
	"FLOAT4":     FLOAT4,
	"FLOAT8":     FLOAT8,
	"DECIMAL":    DECIMAL,
	"DEC":        DEC,
	"NUMERIC":    NUMERIC,
	"DATE":       DATE,
	"TIME":       TIME,
	"TIMESTAMP":  TIMESTAMP,
	"DATETIME":   DATETIME,
	"YEAR":       YEAR,
	"CHAR":       CHAR,
	"VARCHAR":    VARCHAR,
	"NVARCHAR":   NVARCHAR,
	"NATIONAL":   NATIONAL,
	"BINARY":     BINARY,
	"VARBINARY":  VARBINARY,
	"TINYBLOB":   TINYBLOB,
	"BLOB":       BLOB,
	"MEDIUMBLOB": MEDIUMBLOB,
	"LONG":       LONG,
	"LONGBLOB":   LONGBLOB,
	"TINYTEXT":   TINYTEXT,
	"TEXT":       TEXT,
	"MEDIUMTEXT": MEDIUMTEXT,
	"LONGTEXT":   LONGTEXT,
	"ENUM":       ENUM,
	"VARYING":    VARYING,
	"SERIAL":     SERIAL,
}
