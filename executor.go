package core

import (
	"context"

	"github.com/graphql-go/graphql"
)

type PostExecuteHook interface {
	PostExecute(ctx context.Context, req GraphQLRequest, res *graphql.Result)
}

type PreExecuteHook interface {
	PreExecute(ctx context.Context, req GraphQLRequest) context.Context
}

type GraphQLRequest struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

type GraphQLExecutor struct {
	RunAfter  []PostExecuteHook
	RunBefore []PreExecuteHook
	Schema    graphql.Schema
}

func (executor *GraphQLExecutor) Execute(ctx context.Context, req GraphQLRequest) *graphql.Result {
	if executor.RunBefore != nil {
		for _, before := range executor.RunBefore {
			ctx = before.PreExecute(ctx, req)
		}
	}

	result := graphql.Do(graphql.Params{
		Context:        ctx,
		Schema:         executor.Schema,
		OperationName:  req.OperationName,
		RequestString:  req.Query,
		VariableValues: req.Variables,
	})

	if executor.RunAfter != nil {
		for _, after := range executor.RunAfter {
			after.PostExecute(ctx, req, result)
		}
	}

	return result
}
