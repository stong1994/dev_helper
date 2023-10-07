
## 根据sql创建表语句，生成openapi文档

1. `antlr4 -Dlanguage=Go -o parser MySqlParser.g4 MySqlLexer.g4`

## 部署
1. run server
   ```bash
    go run .
   ```
2. 打开浏览器： `localhost:9094`
3. 粘贴建表语句
4. 复制内容
   1. 在浏览器右侧点击复制按钮
   2. 在项目根目录查看openapi.json
   3. 访问 localhost:9094/openapi.json
5. 将内容/地址导入至支持openapi协议的客户端

## 使用细节
1. 布尔类型使用mysql的TINYINT类型