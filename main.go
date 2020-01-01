package main

import (
	"apijoy/config"
	"apijoy/model"
	"apijoy/mutations"
	"apijoy/object"
	"apijoy/queries"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
)

func main() {
	// database connection
	db := config.DbConnection()
	// connection in each packages
	queries.Database = db
	mutations.Database = db
	model.Database = db
	object.Database = db

	// create graphql schema config
	schemaConfig := graphql.SchemaConfig{
		// query object
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:        "Query",
			Fields:      queries.GetRootFields(),
			Description: "List query yang bisa digunakan",
		}),

		// mutation object
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:        "Mutation",
			Fields:      mutations.GetRootFields(),
			Description: "List mutation yang bisa digunakan",
		}),
	}

	// build the schema
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Printf("failed to build schema : %v", err)
	}

	// handler
	httpHandler := handler.New(&handler.Config{
		Schema:   &schema,
		GraphiQL: true,
	})

	http.Handle("/", manageCors(httpHandler))
	// start server
	port := "8888"
	log.Println("server ready at :" + port)
	http.ListenAndServe(":"+port, nil)
}

// manage cors for allowance access from react
func manageCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}
