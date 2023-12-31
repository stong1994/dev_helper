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
		file, err := os.Create("openapi.json")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		_, err = file.Write(jsonBytes)
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
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
	cs.Schemas = make(openapi3.Schemas)
	cs.Schemas["CommonResp"] = commonResp().NewRef()

	swagger := &openapi3.T{
		OpenAPI:    "3.0.1",
		Components: &cs,
		Info: &openapi3.Info{
			Title:          "eebo.ehr.metabase-元数据管理前台", // todo
			TermsOfService: "",
			Version:        "1.0.0",
		},
		Paths:   make(map[string]*openapi3.PathItem),
		Servers: nil,
	}
	tagMap := map[string]bool{}
	for _, table := range tables {
		tag := table.GetDesc()
		if tag == "" {
			tag = table.TableName
		}
		if _, exist := tagMap[tag]; !exist {
			swagger.Tags = append(swagger.Tags, &openapi3.Tag{
				Name: tag,
			})
		}
		schemaView, viewTag := getEntitySchemaView(table)
		schemaCreate, createTag := getEntitySchemaCreate(table)
		schemaEdit, editTag := getEntitySchemaEdit(table)
		refNameView := fmt.Sprintf("#/components/schemas/%s", viewTag)
		refNameCreate := fmt.Sprintf("#/components/schemas/%s", createTag)
		refNameEdit := fmt.Sprintf("#/components/schemas/%s", editTag)
		swagger.Components.Schemas[viewTag] = openapi3.NewSchemaRef("", schemaView)
		swagger.Components.Schemas[createTag] = openapi3.NewSchemaRef("", schemaCreate)
		swagger.Components.Schemas[editTag] = openapi3.NewSchemaRef("", schemaEdit)

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
				//responses := openapi3.NewResponses()
				//responses["200"] = &openapi3.ResponseRef{
				//	Value: openapi3.NewResponse().WithJSONSchema(fillResp(openapi3.NewSchemaRef(refNameView, nil))),
				//	//Ref: refNameView,
				//}
				op.Responses = getResponse(openapi3.NewSchemaRef(refNameView, nil))
				item.Get = op
			case pathTypeList:
				op := openapi3.NewOperation()
				op.Tags = append(op.Tags, tag)
				op.Summary = fmt.Sprintf("批量获取%s", tag)
				op.Parameters = []*openapi3.ParameterRef{
					{Value: &openapi3.Parameter{
						Name:        "p",
						In:          "query",
						Required:    false,
						Description: "页码，从1开始",
						Schema:      openapi3.NewIntegerSchema().NewRef(),
					}},
					{Value: &openapi3.Parameter{
						Name:        "limit",
						In:          "query",
						Required:    false,
						Description: "每页限制数量",
						Schema:      openapi3.NewIntegerSchema().NewRef(),
					}},
					{Value: &openapi3.Parameter{
						Name:        "search",
						In:          "query",
						Required:    false,
						Description: "搜索内容",
						Schema:      openapi3.NewStringSchema().NewRef(),
					}},
				}

				resp := openapi3.NewArraySchema()
				resp.Items = openapi3.NewSchemaRef(refNameView, nil)

				op.Responses = getResponse(resp.NewRef())
				item.Get = op
			case pathTypeCreate:
				op := openapi3.NewOperation()
				op.Tags = append(op.Tags, tag)
				op.Summary = fmt.Sprintf("新增%s", tag)
				op.RequestBody = &openapi3.RequestBodyRef{Value: openapi3.NewRequestBody().WithJSONSchemaRef(openapi3.NewSchemaRef(refNameCreate, nil))}
				resp := openapi3.NewObjectSchema()
				resp.Properties = map[string]*openapi3.SchemaRef{
					"id": openapi3.NewStringSchema().NewRef(),
				}

				op.Responses = getResponse(resp.NewRef())
				item.Post = op
			case pathTypeUpdate:
				op := openapi3.NewOperation()
				op.Tags = append(op.Tags, tag)
				op.Summary = fmt.Sprintf("更新%s", tag)
				op.RequestBody = &openapi3.RequestBodyRef{Value: openapi3.NewRequestBody().WithJSONSchemaRef(openapi3.NewSchemaRef(refNameEdit, nil))}
				resp := openapi3.NewObjectSchema()
				resp.Properties = map[string]*openapi3.SchemaRef{
					"id": openapi3.NewStringSchema().NewRef(),
				}
				op.Responses = getResponse(resp.NewRef())

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
				resp := openapi3.NewObjectSchema()
				resp.Properties = map[string]*openapi3.SchemaRef{
					"id": openapi3.NewStringSchema().NewRef(),
				}
				op.Responses = getResponse(resp.NewRef())
				item.Post = op
			}
			swagger.Paths[path] = &item
		}
	}

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

func getEntitySchemaView(table CreateDDLData) (*openapi3.Schema, string) {
	s := new(openapi3.Schema)
	s.Type = "object"
	s.Properties = make(openapi3.Schemas)
	for _, column := range table.Columns {
		switch column.Type {
		case consts.TINYINT:
			if strings.Contains(column.Name, "is_") {
				schema := openapi3.NewBoolSchema()
				schema.Description = column.Comment
				s.Properties[column.Name] = openapi3.NewSchemaRef("", schema)
			} else {
				schema := openapi3.NewIntegerSchema()
				schema.Description = column.Comment
				s.Properties[column.Name] = openapi3.NewSchemaRef("", schema)
			}
		case consts.SMALLINT, consts.MEDIUMINT, consts.MIDDLEINT, consts.INT, consts.INT1, consts.INT2,
			consts.INT3, consts.INT4, consts.INT8, consts.INTEGER, consts.BIGINT:
			schema := openapi3.NewIntegerSchema()
			schema.Description = column.Comment
			s.Properties[column.Name] = openapi3.NewSchemaRef("", schema)
		default:
			schema := openapi3.NewStringSchema()
			schema.Description = column.Comment
			s.Properties[column.Name] = openapi3.NewSchemaRef("", schema)
		}
		if column.Name == "add_by_id" {
			schema := openapi3.NewStringSchema()
			schema.Description = "添加人姓名"
			s.Properties["add_by_name"] = openapi3.NewSchemaRef("", schema)
		}
		if column.Name == "update_by_id" {
			schema := openapi3.NewStringSchema()
			schema.Description = "更新人姓名"
			s.Properties["update_by_name"] = openapi3.NewSchemaRef("", schema)
		}
	}
	return s, fmt.Sprintf("%s_VIEW", table.GetDesc())
}

func getEntitySchemaEdit(table CreateDDLData) (*openapi3.Schema, string) {
	s := new(openapi3.Schema)
	s.Type = "object"
	s.Properties = make(openapi3.Schemas)
	for _, column := range table.Columns {
		if ignoreFieldsInEditModel(column.Name) {
			continue
		}
		switch column.Type {
		case consts.TINYINT:
			if strings.Contains(column.Name, "is_") {
				schema := openapi3.NewBoolSchema()
				schema.Description = column.Comment
				s.Properties[column.Name] = openapi3.NewSchemaRef("", schema)
			} else {
				schema := openapi3.NewIntegerSchema()
				schema.Description = column.Comment
				s.Properties[column.Name] = openapi3.NewSchemaRef("", schema)
			}

		case consts.SMALLINT, consts.MEDIUMINT, consts.MIDDLEINT, consts.INT, consts.INT1,
			consts.INT2, consts.INT3, consts.INT4, consts.INT8, consts.INTEGER, consts.BIGINT:
			schema := openapi3.NewIntegerSchema()
			schema.Description = column.Comment
			s.Properties[column.Name] = openapi3.NewSchemaRef("", schema)
		default:
			schema := openapi3.NewStringSchema()
			schema.Description = column.Comment
			s.Properties[column.Name] = openapi3.NewSchemaRef("", schema)
		}
	}
	return s, fmt.Sprintf("%s_EDIT", table.GetDesc())
}

func getEntitySchemaCreate(table CreateDDLData) (*openapi3.Schema, string) {
	s := new(openapi3.Schema)
	s.Type = "object"
	s.Properties = make(openapi3.Schemas)
	for i, column := range table.Columns {
		if i == 0 { // always declare the id column first
			continue
		}
		if ignoreFieldsInEditModel(column.Name) {
			continue
		}
		switch column.Type {
		case consts.TINYINT:
			if strings.Contains(column.Name, "is_") {
				schema := openapi3.NewBoolSchema()
				schema.Description = column.Comment
				s.Properties[column.Name] = openapi3.NewSchemaRef("", schema)
			} else {
				schema := openapi3.NewIntegerSchema()
				schema.Description = column.Comment
				s.Properties[column.Name] = openapi3.NewSchemaRef("", schema)
			}

		case consts.SMALLINT, consts.MEDIUMINT, consts.MIDDLEINT, consts.INT, consts.INT1,
			consts.INT2, consts.INT3, consts.INT4, consts.INT8, consts.INTEGER, consts.BIGINT:
			schema := openapi3.NewIntegerSchema()
			schema.Description = column.Comment
			s.Properties[column.Name] = openapi3.NewSchemaRef("", schema)
		default:
			schema := openapi3.NewStringSchema()
			schema.Description = column.Comment
			s.Properties[column.Name] = openapi3.NewSchemaRef("", schema)
		}
	}
	return s, fmt.Sprintf("%s_CREATE", table.GetDesc())
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

func ignoreFieldsInEditModel(field string) bool {
	if field == "add_by_id" || field == "update_by_id" || field == "update_dt" || field == "add_dt" || field == "is_delete" || field == "is_deleted" {
		return true
	}
	return false
}

func fillResp(ref *openapi3.SchemaRef) *openapi3.Schema {
	rst := openapi3.NewObjectSchema()
	code := openapi3.NewIntegerSchema()
	code.Title = "状态码"
	code.Description = "200-ok"
	rst.Properties = map[string]*openapi3.SchemaRef{
		"errormsg":   openapi3.NewStringSchema().NewRef(),
		"resultcode": code.NewRef(),
		"data":       ref,
	}
	return rst
}

func commonResp() *openapi3.Schema {
	rst := openapi3.NewObjectSchema()
	code := openapi3.NewIntegerSchema()
	code.Title = "状态码"
	code.Description = "200-ok"
	rst.Properties = map[string]*openapi3.SchemaRef{
		"errormsg":   openapi3.NewStringSchema().NewRef(),
		"resultcode": code.NewRef(),
		"data":       openapi3.NewSchema().NewRef(),
	}
	return rst
}

func getResponse(ref *openapi3.SchemaRef) openapi3.Responses {
	responses := openapi3.NewResponses()

	responses["200"] = &openapi3.ResponseRef{
		Value: openapi3.NewResponse().WithJSONSchema(fillResp(ref)),
	}
	return responses
}
