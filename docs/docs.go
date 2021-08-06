// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "豆豆",
            "url": "http://www.swagger.io/support",
            "email": "951828649@qq.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user/QueryUserInfoById": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户接口"
                ],
                "summary": "获取个人用户信息",
                "responses": {
                    "200": {
                        "description": "{\"RetCode\":0,\"UserInfo\":{},\"Action\":\"GetAllUserResponse\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/queryUser": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户接口"
                ],
                "summary": "获取所有用户",
                "responses": {
                    "200": {
                        "description": "{\"RetCode\":0,\"UserInfo\":{},\"Action\":\"GetAllUserResponse\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "这里写接口服务的host",
	BasePath:    "这里写base path",
	Schemes:     []string{},
	Title:       "hello world 测试相关接口",
	Description: "go 初步学习",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
