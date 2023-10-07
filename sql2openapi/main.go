package main

import (
	"encoding/json"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/getkin/kin-openapi/openapi3"
	"io"
	"log"
	"net/http"
	"os"
	"sql2openapi/consts"
	"sql2openapi/parser"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	http.HandleFunc("/ddl_create_table", func(w http.ResponseWriter, r *http.Request) {
		bytes, err := io.ReadAll(r.Body)
		//fmt.Println(string(bytes))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		type createDDL struct {
			DDL string `json:"ddl"`
		}
		var request createDDL
		if err = json.Unmarshal(bytes, &request); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		if request.DDL == "" {
			w.Write(nil)
			return
		}

		tables, err := GetTables(request.DDL)
		if err != nil {
			//w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		for i, v := range tables {
			tables[i] = preHandle(v)
		}

		swagger := GenSchema(tables)
		jsonBytes, err := swagger.MarshalJSON()
		if err != nil {
			log.Fatal(err)
		}

		// Write the JSON to a file
		file, err := os.Create("schema.json")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		_, err = file.Write(jsonBytes)
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
		//data, _ := json.Marshal(map[string]interface{}{"struct": rst})
		if _, err = w.Write(jsonBytes); err != nil {
			panic(err)
		}
		fmt.Println("finished!")
	})

	fmt.Println("listening :9094")
	err := http.ListenAndServe(":9094", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func GetTables(ddl string) ([]CreateDDLData, error) {
	lexer := parser.NewMySqlLexer(antlr.NewInputStream(ddl))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	createParser := parser.NewMySqlParser(stream)

	errListener := &ErrListener{}
	createParser.RemoveErrorListeners() // 默认会使用ConsoleErrorListener，需要移除。
	createParser.AddErrorListener(errListener)
	createParser.GetInterpreter().SetPredictionMode(antlr.PredictionModeLLExactAmbigDetection)

	listener := NewCreateTableVisitor()
	antlr.NewParseTreeWalker().Walk(listener, createParser.Root())
	return listener.Data, errListener.Error()
}

func GenSchema(tables []CreateDDLData) *openapi3.T {
	cs := openapi3.NewComponents()
	swagger := &openapi3.T{
		OpenAPI:    "3.0.1",
		Components: &cs,
		Info: &openapi3.Info{
			Title:          "eebo.ehr.metabase-元数据管理前台",
			TermsOfService: "",
			Version:        "1.0.0",
		},
		Paths:   make(map[string]*openapi3.PathItem),
		Servers: nil,
	}
	tagMap := map[string]bool{}
	schemas := make(map[string]*openapi3.SchemaRef)
	for _, table := range tables {
		tag := getTag(table.Comment)
		if tag == "" {
			tag = table.TableName
		}
		if _, exist := tagMap[tag]; !exist {
			swagger.Tags = append(swagger.Tags, &openapi3.Tag{
				Name: tag,
			})
		}
		component := getComponent(table)
		//refName := fmt.Sprintf("#/components/schemas/%s", tag)
		schemas[tag] = component

		for i, path := range getPaths(table.TableName) {
			item := openapi3.PathItem{}
			switch i {
			case pathTypeGet:
				op := openapi3.NewOperation()
				op.Tags = append(op.Tags, tag)
				op.Summary = fmt.Sprintf("获取%s", tag)
				op.Parameters = []*openapi3.ParameterRef{{Value: &openapi3.Parameter{
					Name:     "id",
					In:       "query",
					Required: true,
					Schema:   openapi3.NewStringSchema().NewRef(),
				}}}
				responses := openapi3.NewResponses()
				responses["200"] = &openapi3.ResponseRef{
					Value: openapi3.NewResponse().WithJSONSchemaRef(component),
				}
				op.Responses = responses
				item.Get = op
			case pathTypeList:
				op := openapi3.NewOperation()
				op.Tags = append(op.Tags, tag)
				op.Summary = fmt.Sprintf("批量获取%s", tag)
				op.Parameters = []*openapi3.ParameterRef{
					{Value: &openapi3.Parameter{
						Name:     "p",
						In:       "query",
						Required: false,
						Schema:   openapi3.NewIntegerSchema().NewRef(),
					}},
					{Value: &openapi3.Parameter{
						Name:     "limit",
						In:       "query",
						Required: false,
						Schema:   openapi3.NewIntegerSchema().NewRef(),
					}},
				}
				responses := openapi3.NewResponses()

				resp := openapi3.NewArraySchema()
				resp.Properties = map[string]*openapi3.SchemaRef{"application/json": component}
				responses["200"] = &openapi3.ResponseRef{
					Value: openapi3.NewResponse().WithJSONSchemaRef(&openapi3.SchemaRef{
						Value: openapi3.NewArraySchema(),
					}),
				}
				op.Responses = responses
				item.Get = op
			case pathTypeCreate:
				op := openapi3.NewOperation()
				op.Tags = append(op.Tags, tag)
				op.Summary = fmt.Sprintf("新增%s", tag)
				op.RequestBody = &openapi3.RequestBodyRef{Value: openapi3.NewRequestBody().WithJSONSchemaRef(component)}
				item.Post = op
			case pathTypeUpdate:
				op := openapi3.NewOperation()
				op.Tags = append(op.Tags, tag)
				op.Summary = fmt.Sprintf("更新%s", tag)
				op.RequestBody = &openapi3.RequestBodyRef{Value: openapi3.NewRequestBody().WithJSONSchemaRef(component)}
				item.Post = op
			case pathTypeDelete:
				op := openapi3.NewOperation()
				op.Tags = append(op.Tags, tag)
				op.Summary = fmt.Sprintf("删除%s", tag)
				schema := openapi3.NewStringSchema()
				schema.Title = "id"
				op.RequestBody = &openapi3.RequestBodyRef{Value: &openapi3.RequestBody{
					Required: true,
					Content:  openapi3.NewContentWithJSONSchema(schema),
				}}
				item.Post = op
			}
			swagger.Paths[path] = &item
		}
	}
	swagger.Components.Schemas = schemas
	return swagger
}

func preHandle(table CreateDDLData) CreateDDLData {
	if strings.HasPrefix(table.TableName, "t_") {
		table.TableName = strings.TrimPrefix(table.TableName, "t_")
	}
	for i, v := range table.Columns {
		if strings.HasPrefix(v.Name, "c_") {
			table.Columns[i].Name = strings.TrimPrefix(v.Name, "c_")
		}
	}
	return table
}

func getComponent(table CreateDDLData) *openapi3.SchemaRef {
	sr := new(openapi3.SchemaRef)
	sr.Value = new(openapi3.Schema)
	sr.Value.Type = "object"
	sr.Value.Properties = make(openapi3.Schemas)
	for _, column := range table.Columns {
		switch column.Type {
		case consts.TINYINT, consts.SMALLINT, consts.MEDIUMINT, consts.MIDDLEINT,
			consts.INT,
			consts.INT1,
			consts.INT2,
			consts.INT3,
			consts.INT4,
			consts.INT8,
			consts.INTEGER,
			consts.BIGINT:
			schema := openapi3.NewIntegerSchema()
			schema.Description = column.Comment
			sr.Value.Properties[column.Name] = openapi3.NewSchemaRef("", schema)
		default:
			schema := openapi3.NewStringSchema()
			schema.Description = column.Comment
			sr.Value.Properties[column.Name] = openapi3.NewSchemaRef("", schema)

		}

	}
	return sr
}

func getTag(comment string) string {
	if comment == "" {
		return ""
	}
	if strings.HasSuffix(comment, "表") {
		return strings.TrimSuffix(comment, "表")
	}
	return comment
}

func getPaths(tableName string) []string {
	return []string{
		fmt.Sprintf("/%s/get", tableName),
		fmt.Sprintf("/%s/list", tableName),
		fmt.Sprintf("/%s/create", tableName),
		fmt.Sprintf("/%s/update", tableName),
		fmt.Sprintf("/%s/delete", tableName),
	}
}

const (
	pathTypeGet = iota
	pathTypeList
	pathTypeCreate
	pathTypeUpdate
	pathTypeDelete
)
