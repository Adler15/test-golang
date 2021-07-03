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
        "termsOfService": "https://github.com/Adler15",
        "contact": {
            "name": "Victor",
            "url": "https://github.com/Adler15",
            "email": "chenfeilong1115@hotmail.com"
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
        "/create": {
            "post": {
                "description": "新增学生数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户的增删查改"
                ],
                "summary": "新增学生信息接口",
                "parameters": [
                    {
                        "description": "Student",
                        "name": "student",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/student.StudentTest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "RowsAffected",
                        "schema": {
                            "type": "int"
                        }
                    }
                }
            }
        },
        "/createTest": {
            "post": {
                "description": "新增测试类",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "测试类的增删查改"
                ],
                "summary": "新增测试类",
                "parameters": [
                    {
                        "description": "Test",
                        "name": "test",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/test.Test"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功的条数",
                        "schema": {
                            "type": "int"
                        }
                    }
                }
            }
        },
        "/delete": {
            "delete": {
                "description": "删除学生数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户的增删查改"
                ],
                "summary": "删除学生信息接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Student ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "RowsAffected",
                        "schema": {
                            "type": "int"
                        }
                    }
                }
            }
        },
        "/deleteTest": {
            "delete": {
                "description": "删除测试类",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "测试类的增删查改"
                ],
                "summary": "删除测试类",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Test ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功的条数",
                        "schema": {
                            "type": "int"
                        }
                    }
                }
            }
        },
        "/findAll": {
            "get": {
                "description": "查询全部学生数据",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户的增删查改"
                ],
                "summary": "获取全部学生数据接口",
                "responses": {
                    "200": {
                        "description": "All Student",
                        "schema": {
                            "$ref": "#/definitions/student.StudentTest"
                        }
                    }
                }
            }
        },
        "/findAllTest": {
            "get": {
                "description": "查询全部测试类",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "测试类的增删查改"
                ],
                "summary": "查询全部测试类",
                "responses": {
                    "200": {
                        "description": "All Test",
                        "schema": {
                            "$ref": "#/definitions/test.Test"
                        }
                    }
                }
            }
        },
        "/findByName": {
            "get": {
                "description": "通过名字查询学生数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户的增删查改"
                ],
                "summary": "通过名字获取某学生信息接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Student Name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Student",
                        "schema": {
                            "$ref": "#/definitions/student.StudentTest"
                        }
                    }
                }
            }
        },
        "/findByWhere": {
            "get": {
                "description": "通过条件查询测试类",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "测试类的增删查改"
                ],
                "summary": "通过条件查询测试类",
                "parameters": [
                    {
                        "description": "Where of Test",
                        "name": "where",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/test.Test"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Where Test",
                        "schema": {
                            "$ref": "#/definitions/test.Test"
                        }
                    }
                }
            }
        },
        "/findOne": {
            "get": {
                "description": "查询一个学生数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户的增删查改"
                ],
                "summary": "获取某学生信息接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Student",
                        "schema": {
                            "$ref": "#/definitions/student.StudentTest"
                        }
                    }
                }
            }
        },
        "/findOneTest": {
            "get": {
                "description": "查询一条测试类",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "测试类的增删查改"
                ],
                "summary": "查询一条测试类",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Test ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Test",
                        "schema": {
                            "$ref": "#/definitions/test.Test"
                        }
                    }
                }
            }
        },
        "/update": {
            "put": {
                "description": "更新学生数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户的增删查改"
                ],
                "summary": "更新学生信息接口",
                "parameters": [
                    {
                        "description": "Student",
                        "name": "student",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/student.StudentTest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "RowsAffected",
                        "schema": {
                            "type": "int"
                        }
                    }
                }
            }
        },
        "/updateTest": {
            "put": {
                "description": "更新测试类",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "测试类的增删查改"
                ],
                "summary": "更新测试类",
                "parameters": [
                    {
                        "description": "Test",
                        "name": "test",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/test.Test"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功的条数",
                        "schema": {
                            "type": "int"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "student.StudentTest": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "age": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "test.Test": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "test1": {
                    "type": "string"
                },
                "test2": {
                    "type": "string"
                },
                "test3": {
                    "type": "string"
                },
                "test4": {
                    "type": "string"
                },
                "test5": {
                    "type": "string"
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
	Host:        "localhost:9195",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "The Golang Project of Victor",
	Description: "Victor的golang测试项目",
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
