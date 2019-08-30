package main // import "github.com/rakin92/go-gql-starter"

import (
	"context"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/rakin92/go-gql-starter/db"
	"github.com/rakin92/go-gql-starter/handler"
	"github.com/rakin92/go-gql-starter/migrations"
	"github.com/rakin92/go-gql-starter/resolvers"
	"github.com/rakin92/go-gql-starter/schema"
	"github.com/rakin92/go-gql-starter/utils"
)

func main() {
	utils.ConfigureEnv()

	db, err := db.ConnectDB()
	migrations.MigrateDB(db)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	context.Background()

	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(*schema.NewSchema(), &resolvers.Resolvers{DB: db}, opts...)

	mux := http.NewServeMux()
	mux.Handle("/", handler.GraphiQL{})
	mux.Handle("/query", handler.Authenticate(&handler.GraphQL{Schema: schema}))

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Listening to... port 8080")
	if err = s.ListenAndServe(); err != nil {
		panic(err)
	}
}
