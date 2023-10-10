package main

import (
	"encoding/json"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/codegen"
	"github.com/getkin/kin-openapi/openapi3"
	"io"
	"log"
	"net/http"
	url2 "net/url"
	"strings"
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
			Url string `json:"url"`
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

		content, err := genParams(swagger)
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
	//for _, v := range swagger.Paths {
	//	if v.Post != nil {
	//		for _, vv := range v.Post.RequestBody.Value.Content {
	//			vv.Schema.Ref, err = url2.QueryUnescape(vv.Schema.Ref)
	//			if err != nil {
	//				return "", fmt.Errorf("unescape failed: %w", err)
	//			}
	//		}
	//	}
	//	if v.Get != nil {
	//		for _, vv := range v.Get.Responses {
	//			for _, vvv := range vv.Value.Content {
	//				for _, vvvv := range vvv.Schema.Value.Properties {
	//					vvvv.Ref, err = url2.QueryUnescape(vvvv.Ref)
	//					if err != nil {
	//						return "", fmt.Errorf("unescape failed: %w", err)
	//					}
	//
	//				}
	//			}
	//		}
	//	}
	//
	//}
	//
	//for i, v := range swagger.Components.Schemas {
	//
	//}
	code, err := codegen.Generate(swagger, codegen.Configuration{
		PackageName: "dto",
		Generate: codegen.GenerateOptions{
			Models: true,
		},
		Compatibility: codegen.CompatibilityOptions{
			OldMergeSchemas:                    false,
			OldEnumConflicts:                   false,
			OldAliasing:                        false,
			DisableFlattenAdditionalProperties: false,
			DisableRequiredReadOnlyAsPointer:   true,
			AlwaysPrefixEnumValues:             false,
			ApplyChiMiddlewareFirstToLast:      false,
			ApplyGorillaMiddlewareFirstToLast:  false,
			CircularReferenceLimit:             0,
		},
	})
	if err != nil {
		return "", fmt.Errorf("error generating code: %w", err)
	}
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
