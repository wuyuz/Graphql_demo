package graph

import "gragql_demo/graph/model"

//go:generate go run github.com/99designs/gqlgen
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	users []*model.User  // 这里可以自定义的添加一些属性，可以在解析器中调用
}
