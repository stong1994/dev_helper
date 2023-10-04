sql2gostruct用于将sql转化为go结构体。

## 使用步骤
1. Run Server
```bash
go run .
```
2. 访问页面 http://localhost:9090/
3. 输入sql

## 开发步骤
使用antlr4生成go文件
```bash
 cd antlr4
 antlr4 -Dlanguage=Go -o ../parser CreateParser.g4 MySqlLexer.g4
```