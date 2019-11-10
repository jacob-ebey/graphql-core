package schemas

import (
	"github.com/graphql-go/graphql"
)

var PingPongQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"ping": &graphql.Field{
			Type: graphql.String,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return "Pong", nil
			},
		},
	},
})

var PingPongSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: PingPongQuery,
})
