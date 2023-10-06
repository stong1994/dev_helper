sql2gostruct用于将sql转化为go结构体。

## 使用步骤
1. Run Server
```bash
go run web/main.go
```
2. 访问页面 http://localhost:9094/
3. 输入sql

## 开发步骤
使用antlr4生成go文件
```bash
 cd antlr4
 antlr4 -Dlanguage=Go -o ../parser -visitor MySqlParser.g4 MySqlLexer.g4
```
### antlr4 gui
```bash
antlr4-parse MySqlParser.g4 MySqlLexer.g4 ddlStatement -gui 
```
mysql antlr4实现请参考[antlr4官方例子](https://github.com/antlr/grammars-v4/tree/master/sql/mysql/Positive-Technologies)