## 目的
根据openapi文件转成go代码。
## 目标结构
├── router
│        ├── {{ModuleName1}}.go
│        ├── {{ModuleName2}}.go
├── controller
│   |── {{ModuleName1}}
│        ├── {{ModuleName1}}.go
│   |── {{ModuleName2}}
│        ├── {{ModuleName2}}.go
├── service
│   |── entity
│        ├── dto
│            ├── {{ModuleName1}}.go
│            ├── {{ModuleName2}}.go
│        ├── do
│            ├── {{ModuleName1}}.go
│            ├── {{ModuleName2}}.go
│   |── {{ModuleName1}}
│        ├── {{ModuleName1}}.go
│   |── {{ModuleName2}}
│        ├── {{ModuleName2}}.go
├── adaptor
│   |── {{ModuleName1}}
│        ├── {{ModuleName1}}.go
│   |── {{ModuleName2}}
│        ├── {{ModuleName2}}.go

## 模块——module
根据openapi的tag生成对应的模块，即openapi的tag与生成代码的模块一一对应
## router
router对每个模块生成一个单独的文件，文件名为ModuleName.go.
文件内容为：
```go
package routes

import "github.com/kataras/iris/v12"

func (r *Router) {{ModuleName}}(accountingRoute iris.Party) {
	{{ModuleName}}Route.Post("{{path}}", r.controllers.{{ModuleName}}Ctl.{{RestName}})
	...
```
其中RestName根据路径最后两个标识生成，如/salary/get => GetSalary

## controller
在controller，每个模块都有一个独立的目录，目录名为{{module_name}}.子文件名为{{module_name}}.go
对于文件内容，首先是Controller定义：
```go
package {{module_name}}

import (
    "{{project.name}}/service"
    "{{project.name}}/service/dto"
    "{{project.name}}/service/consts"
)

type Controller struct {
	services *service.Services
}

func New{{ModuleName}}Controller(services *service.Services) *Controller {
	return &Controller{
		services: services,
	}
}
```
然后是每个api对应的方法。
如果是POST方法，则使用模版：
```go
func (ctl Controller) {{RestName}}(ctx iris.Context) {
    if !controller.CheckPermission(ctx, consts.PermCode{{ModuleName}}Edit) {
        return
    }
	var param dto.{{RestName}}Req
	err := ctx.ReadJSON(&param)
	if err != nil {
		WriterResp(ctx, nil, service.ParamInvalid.WithErr(err))
		return
	}
	data, errno := ctl.services.{{ModuleName}}Svc.{{RestName}}(ctx.Request().Context(), param)
    controller.WriteResp(ctx, data, errno)
}
```
如果是GET方法，则使用模版：
```go
func (ctl Controller) {{RestName}}(ctx iris.Context) {
    if !controller.CheckPermission(ctx, consts.PermCode{{ModuleName}}View) {
        return
    }
	var param dto.{{RestName}}Req
	err := ctx.ReadQuery(&param)
	if err != nil {
		WriterResp(ctx, nil, service.ParamInvalid.WithErr(err))
		return
	}
	data, errno := ctl.services.{{ModuleName}}Svc.{{RestName}}(ctx.Request().Context(), param)
    controller.WriteResp(ctx, data, errno)
}
```

## service
service下每个模块都有自己的目录，规则与Controller相同。
对于文件内容，首先是`adaptor.go`:
```go
import (
    "context"
	"{{project.name}}/service/dto"
	"{{project.name}}/service/do"
)


type Base{{ModuleName}} interface {
    {{RestName}}(ctx context.Context, param dto.{{RestName}}Req) (dto.{{RestName}}Resp, error)
	...
}
```
对于{{module_name}}.go文件内容：
```go
package {{module_name}}

import (
	"{{project.name}}/service"
)

type Service struct {
	adaptor Base{{ModuleName}}
}

func NewService(adaptors Base{{ModuleName}}) *Service {
	return &Service{
		adaptor: adaptor,
	}
}
```
然后是每个API：
```go
func (s Service) {{RestName}}(ctx context.Context, param dto.{{RestName}}Req) (dto.{{RestName}}Resp, service.Errno) {

	data, err := s.adaptor.{{RestName}}(ctx, param)
	if err != nil {
		return nil, service.{{ProjectName}}AccessErr.WithErr(err)
	}

	return data, errno.OK
}
```

## adaptor
目录结构同controller。
```go
package {{module_name}}

import (
	"context"
	"{{project.name}}/config"
    "{{project.name}}/adaptor/common"
	"{{project.name}}/service/dto"

    "github.com/opentracing/opentracing-go"
)

type Adaptor struct {
	tracer         opentracing.Tracer
	bpHost         string
}

func NewAdaptor(tracer opentracing.Tracer, bpHost string) *Adaptor {
	return &Adaptor{
		tracer:         tracer,
		bpHost:         bpHost,
	}
}
```
如果是GET请求：
```go
type {{restName}}Resp struct {
	common.CommonResp
	Data dto.{{RestName}}Resp `json:"data"`
}

func (a *Adaptor) {{RestName}}(ctx context.Context, param dto.{{RestName}}Req) (dto.{{RestName}}Resp, error) {
    var (
        data dto.{{RestName}}Resp
        url  = a.bpHost + "{{path}}"
        resp {{restName}}Resp
        req  = iris.Map{
            "company_id": ctx.Value("company_id"),
			"user_id": ctx.Value("user_id"),
        }
    )
	req = common.MergeQueryParam(req, param)
    if err := thttp.SpanGet(ctx, a.tracer, config.ServerFullName, url, common.GetHttpHeader, req, &resp); err != nil {
        return nil, err
    }
    if resp.NotOK() {
        return nil, resp
    }
	
	return resp.Data, nil
}
```
如果是POST请求：
```go
type {{restName}}Resp struct {
	common.CommonResp
	Data dto.{{RestName}}Resp `json:"data"`
}

func (a *Adaptor) {{RestName}}(ctx context.Context, param dto.{{RestName}}Req) (dto.{{RestName}}Resp, error) {
    var (
        data dto.{{RestName}}Resp
        url  = a.bpHost + "{{path}}"
        resp {{restName}}Resp
        req  = iris.Map{
            "company_id": ctx.Value("company_id"),
			"user_id": ctx.Value("user_id"),
        }
    )
	req = common.MergePostParam(req, param)
    if err := thttp.SpanPost(ctx, a.tracer, config.ServerFullName, url, common.PostHttpHeader, req, &resp); err != nil {
        return nil, err
    }
    if resp.NotOK() {
        return nil, resp
    }
	
	return resp.Data, nil
}
```