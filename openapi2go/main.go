package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/codegen"
	"github.com/getkin/kin-openapi/openapi3"
	"io"
	"log"
	"net/http"
	url2 "net/url"
	"regexp"
	"strings"
	"text/template"
	"unicode"
)

func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	http.HandleFunc("/gen_param", func(w http.ResponseWriter, r *http.Request) {
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		fmt.Println(string(bytes))

		type genParam struct {
			Url  string `json:"url"`
			Type string `json:"type"`
		}
		var request genParam
		if err = json.Unmarshal(bytes, &request); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		if request.Url == "" {
			w.Write(nil)
			return
		}

		swagger, err := getSwagger(request.Url)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		fileName := getFileName(swagger.Paths)
		var content string

		switch request.Type {
		case "dto":
			content, err = genParams(swagger)
		case "serviceModule":
			content, err = genServiceModule(swagger)
		case "adaptorModule":
			content, err = genAdapterModule(swagger)
		case "routerModule":
			content, err = genRouterModule(swagger)
		case "controllerModule":
			content, err = genControllerModule(swagger)
		case "permCode":
			content, err = genPermCodeModule(swagger)
		}
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		data := map[string]string{"fileName": fileName, "content": content}
		resp, err := json.Marshal(data)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		if _, err = w.Write(resp); err != nil {
			panic(err)
		}
		fmt.Println("finished!")
	})

	fmt.Println("listening :9095")
	err := http.ListenAndServe(":9095", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func getSwagger(url string) (*openapi3.T, error) {
	loader := openapi3.NewLoader()
	path, err := url2.Parse(url)
	if err != nil {
		return nil, err
	}
	swagger, err := loader.LoadFromURI(path)
	if err != nil {
		return nil, err
	}
	return swagger, nil
}

func genParams(swagger *openapi3.T) (string, error) {
	// no pointer
	for _, v := range swagger.Components.Schemas {
		for _, p := range v.Value.Properties {
			p.Value.Extensions["x-go-type-skip-optional-pointer"] = true
		}
	}
	for _, v := range swagger.Paths {
		if v.Get != nil {
			for _, p := range v.Get.Parameters {
				p.Value.Extensions["x-go-type-skip-optional-pointer"] = true
			}
		}
		if v.Post != nil {
			for _, c := range v.Post.RequestBody.Value.Content {
				for _, p := range c.Schema.Value.Properties {
					p.Value.Extensions["x-go-type-skip-optional-pointer"] = true
				}
			}
		}
	}

	code, err := codegen.Generate(swagger, codegen.Configuration{
		PackageName: "dto",
		Generate: codegen.GenerateOptions{
			Models: true,
		},
	})
	if err != nil {
		return "", fmt.Errorf("error generating code: %w", err)
	}

	code += `
type PagenationResponse interface{}
`
	code = strings.ReplaceAll(code, "Id", "ID")
	return code, nil
}

func getFileName(paths openapi3.Paths) string {
	for k := range paths {
		sps := strings.Split(k, "/")
		for _, sp := range sps {
			if sp != "" {
				return sp
			}
		}
	}
	return ""
}

type RouterModuleAPI struct {
	RestName string
	Path     string
	Get      bool
}

type RouterModuleGen struct {
	API             []RouterModuleAPI
	ModuleName      string
	ProjectName     string
	ModuleNameSnake string
}

type ControllerModuleAPI struct {
	RestName string
	Param    string
	Get      bool
}

type ControllerModuleGen struct {
	API             []ControllerModuleAPI
	ModuleName      string
	ProjectName     string
	ModuleNameSnake string
}

type ServiceModuleAPI struct {
	RestName string
	Param    string
	Resp     string
}

type ServiceModuleGen struct {
	API             []ServiceModuleAPI
	ModuleName      string
	ProjectName     string
	BpProjectName   string
	ModuleNameSnake string
}

type AdapterModuleAPI struct {
	RestName string
	Param    string
	Resp     string
	Get      bool
	Path     string
}

type AdapterModuleGen struct {
	API             []AdapterModuleAPI
	ModuleName      string
	ProjectName     string
	BpProjectName   string
	ModuleNameSnake string
}

func genRouterModule(swagger *openapi3.T) (string, error) {
	tmpl := "template/router_module.tmpl"

	tl := template.New(tmpl)
	data, err := codegen.GetUserTemplateText(tmpl)
	if err != nil {
		return "", err
	}

	tl, err = tl.Parse(data)
	if err != nil {
		return "", err
	}

	firstPath := swagger.Paths.InMatchingOrder()[0]
	gen := RouterModuleGen{
		ModuleName:      getModuleName(firstPath),
		ProjectName:     "eebo.ehr.metabase",
		ModuleNameSnake: getModuleNameSnake(firstPath),
	}

	for _, url := range swagger.Paths.InMatchingOrder() {
		gen.API = append(gen.API, RouterModuleAPI{
			RestName: getRestName(url),
			Get:      swagger.Paths[url].Get != nil,
			Path:     url,
			//Param: codegen.GenerateTypesForSchemas(),
		})
	}

	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	err = tl.Execute(w, gen)
	if err != nil {
		return "", err
	}
	w.Flush()
	return buf.String(), nil
}

func genPermCodeModule(swagger *openapi3.T) (string, error) {
	tmpl := "template/perm_code.tmpl"

	tl := template.New(tmpl)
	data, err := codegen.GetUserTemplateText(tmpl)
	if err != nil {
		return "", err
	}

	tl, err = tl.Parse(data)
	if err != nil {
		return "", err
	}

	firstPath := swagger.Paths.InMatchingOrder()[0]
	gen := ControllerModuleGen{
		ModuleName:      getModuleName(firstPath),
		ProjectName:     "eebo.ehr.metabase",
		ModuleNameSnake: getModuleNameSnake(firstPath),
	}

	for _, url := range swagger.Paths.InMatchingOrder() {
		param, err := GetRequestParamTypeNameBySchema(url, swagger.Paths[url])
		if err != nil {
			return "", err
		}
		gen.API = append(gen.API, ControllerModuleAPI{
			RestName: getRestName(url),
			Get:      swagger.Paths[url].Get != nil,
			//Param: codegen.GenerateTypesForSchemas(),
			Param: param,
		})
	}

	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	err = tl.Execute(w, gen)
	if err != nil {
		return "", err
	}
	w.Flush()
	return buf.String(), nil
}

func genControllerModule(swagger *openapi3.T) (string, error) {
	tmpl := "template/controller_module.tmpl"

	tl := template.New(tmpl)
	data, err := codegen.GetUserTemplateText(tmpl)
	if err != nil {
		return "", err
	}

	tl, err = tl.Parse(data)
	if err != nil {
		return "", err
	}

	firstPath := swagger.Paths.InMatchingOrder()[0]
	gen := ControllerModuleGen{
		ModuleName:      getModuleName(firstPath),
		ProjectName:     "eebo.ehr.metabase",
		ModuleNameSnake: getModuleNameSnake(firstPath),
	}

	for _, url := range swagger.Paths.InMatchingOrder() {
		param, err := GetRequestParamTypeNameBySchema(url, swagger.Paths[url])
		if err != nil {
			return "", err
		}
		gen.API = append(gen.API, ControllerModuleAPI{
			RestName: getRestName(url),
			Get:      swagger.Paths[url].Get != nil,
			//Param: codegen.GenerateTypesForSchemas(),
			Param: param,
		})
	}

	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	err = tl.Execute(w, gen)
	if err != nil {
		return "", err
	}
	w.Flush()
	return buf.String(), nil
}

func genServiceModule(swagger *openapi3.T) (string, error) {
	tmpl := "template/service_module.tmpl"

	tl := template.New(tmpl)
	data, err := codegen.GetUserTemplateText(tmpl)
	if err != nil {
		return "", err
	}

	tl, err = tl.Parse(data)
	if err != nil {
		return "", err
	}

	firstPath := swagger.Paths.InMatchingOrder()[0]
	gen := ServiceModuleGen{
		ModuleName:      getModuleName(firstPath),
		ProjectName:     "eebo.ehr.metabase",
		ModuleNameSnake: getModuleNameSnake(firstPath),
		BpProjectName:   "BpMetabase",
	}

	for _, url := range swagger.Paths.InMatchingOrder() {
		param, err := GetRequestParamTypeNameBySchema(url, swagger.Paths[url])
		if err != nil {
			return "", err
		}
		resp, err := GetRespTypeNameBySchema(url, swagger.Paths[url])
		if err != nil {
			return "", err
		}
		gen.API = append(gen.API, ServiceModuleAPI{
			RestName: getRestName(url),
			//Param: codegen.GenerateTypesForSchemas(),
			Param: param,
			Resp:  resp,
		})
	}

	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	err = tl.Execute(w, gen)
	if err != nil {
		return "", err
	}
	w.Flush()
	return buf.String(), nil
}

func genAdapterModule(swagger *openapi3.T) (string, error) {
	tmpl := "template/adapter_module.tmpl"

	tl := template.New(tmpl)
	data, err := codegen.GetUserTemplateText(tmpl)
	if err != nil {
		return "", err
	}

	tl, err = tl.Parse(data)
	if err != nil {
		return "", err
	}

	firstPath := swagger.Paths.InMatchingOrder()[0]
	gen := AdapterModuleGen{
		ModuleName:      getModuleName(firstPath),
		ProjectName:     "eebo.ehr.metabase",
		ModuleNameSnake: getModuleNameSnake(firstPath),
		BpProjectName:   "BpMetabase",
	}

	for _, url := range swagger.Paths.InMatchingOrder() {
		param, err := GetRequestParamTypeNameBySchema(url, swagger.Paths[url])
		if err != nil {
			return "", err
		}
		resp, err := GetRespTypeNameBySchema(url, swagger.Paths[url])
		if err != nil {
			return "", err
		}
		gen.API = append(gen.API, AdapterModuleAPI{
			Get:      swagger.Paths[url].Get != nil,
			RestName: getRestName(url),
			Path:     url,
			//Param: codegen.GenerateTypesForSchemas(),
			Param: param,
			Resp:  resp,
		})
	}

	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	err = tl.Execute(w, gen)
	if err != nil {
		return "", err
	}
	w.Flush()
	return buf.String(), nil
}

func getModuleName(url string) string {
	sps := strings.Split(url, "/")
	target := sps[0]
	if len(sps) >= 2 {
		target = sps[len(sps)-2]
	}
	return convertToCamelCase(target)
}

func getModuleNameSnake(url string) string {
	sps := strings.Split(url, "/")
	target := sps[0]
	if len(sps) >= 2 {
		target = sps[len(sps)-2]
	}
	return target
}

func getRestName(url string) string {
	sps := strings.Split(url, "/")
	// 取最后两个
	if len(sps) >= 2 {
		sps = sps[len(sps)-2:]
	}
	for i := 0; i < len(sps)/2; i++ {
		sps[i], sps[len(sps)-i-1] = sps[len(sps)-i-1], sps[i]
	}
	return convertToCamelCase(strings.Join(sps, "_"))
}

func convertToCamelCase(word string) string {
	if !strings.Contains(word, "_") {
		var s []rune
		for _, c := range word {
			if len(s) == 0 {
				s = append(s, unicode.ToUpper(c))
			} else {
				s = append(s, c)
			}
		}
		return string(s)
	}
	words := strings.Split(word, "_")
	var rst string
	for _, w := range words {
		rst += convertToCamelCase(w)
	}
	return rst
}

func GetRequestParamTypeNameBySchema(url string, path *openapi3.PathItem) (string, error) {
	if path.Post != nil {
		name, err := generateDefaultOperationID("POST", url, ToCamelCase)
		if err != nil {
			return "", err
		}
		return "dto." + name + "JSONRequestBody", nil
	}
	name, err := generateDefaultOperationID("GET", url, ToCamelCase)
	if err != nil {
		return "", err
	}
	return "dto." + name + "Params", nil
}

func getTypeNameBySchema(url string, path *openapi3.PathItem) (string, error) {
	if path.Post != nil {
		return generateDefaultOperationID("POST", url, ToCamelCase)
	}
	return generateDefaultOperationID("GET", url, ToCamelCase)
}

func generateDefaultOperationID(opName string, requestPath string, toCamelCaseFunc func(string) string) (string, error) {
	var operationId = strings.ToLower(opName)

	if opName == "" {
		return "", fmt.Errorf("operation name cannot be an empty string")
	}

	if requestPath == "" {
		return "", fmt.Errorf("request path cannot be an empty string")
	}

	for _, part := range strings.Split(requestPath, "/") {
		if part != "" {
			operationId = operationId + "-" + part
		}
	}

	return toCamelCaseFunc(operationId), nil
}

func ToCamelCase(str string) string {
	s := strings.Trim(str, " ")

	n := ""
	capNext := true
	for _, v := range s {
		if unicode.IsUpper(v) {
			n += string(v)
		}
		if unicode.IsDigit(v) {
			n += string(v)
		}
		if unicode.IsLower(v) {
			if capNext {
				n += strings.ToUpper(string(v))
			} else {
				n += string(v)
			}
		}
		_, capNext = separatorSet[v]
	}
	return n
}

func GetRespTypeNameBySchema(url string, path *openapi3.PathItem) (string, error) {
	if path.Post != nil {
		name, err := generateDefaultOperationID("POST", url, ToCamelCase)
		if err != nil {
			return "", err
		}
		return "dto." + name + "JSONBody", nil
	}
	//name, err := generateDefaultOperationID("GET", url, ToCamelCase)
	//if err != nil {
	//	return "", err
	//}
	if strings.HasSuffix(url, "/get") {
		return "dto." + getModuleName(url) + "View", nil
	}
	return "dto.PagenationResponse", nil
}

var (
	pathParamRE    *regexp.Regexp
	predeclaredSet map[string]struct{}
	separatorSet   map[rune]struct{}
)

func init() {
	pathParamRE = regexp.MustCompile(`{[.;?]?([^{}*]+)\*?}`)

	predeclaredIdentifiers := []string{
		// Types
		"bool",
		"byte",
		"complex64",
		"complex128",
		"error",
		"float32",
		"float64",
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"rune",
		"string",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"uintptr",
		// Constants
		"true",
		"false",
		"iota",
		// Zero value
		"nil",
		// Functions
		"append",
		"cap",
		"close",
		"complex",
		"copy",
		"delete",
		"imag",
		"len",
		"make",
		"new",
		"panic",
		"print",
		"println",
		"real",
		"recover",
	}
	predeclaredSet = map[string]struct{}{}
	for _, id := range predeclaredIdentifiers {
		predeclaredSet[id] = struct{}{}
	}

	separators := "-#@!$&=.+:;_~ (){}[]"
	separatorSet = map[rune]struct{}{}
	for _, r := range separators {
		separatorSet[r] = struct{}{}
	}
}
