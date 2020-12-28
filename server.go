package main

import (
	"github.com/go-pg/pg/v9"
	"gragql_demo/graph"
	"gragql_demo/graph/generated"
	"gragql_demo/graph/postgres"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"




func main() {

	DB := postgres.New(&pg.Options{
		User:     "postgres",
		Password: "password",
		Database: "postgres",
	})

	defer DB.Close()
	// 实现数据库日志接口钩子
	DB.AddQueryHook(postgres.DBLogger{})

	// 初始化配置
	c := generated.Config{Resolvers: &graph.Resolver{
		MeetupsRepo: postgres.MeetupsRepo{DB: DB},
		UsersRepo:   postgres.UsersRepo{DB: DB},
	}}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))


	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
