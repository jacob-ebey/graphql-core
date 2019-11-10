package core_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/graphql-go/graphql"

	core "github.com/jacob-ebey/graphql-core"
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
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

var SchemaType, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: QueryType,
})

func TestExecutorExecutesSchema(t *testing.T) {
	executor := core.GraphQLExecutor{
		Schema: SchemaType,
	}

	res := executor.Execute(context.TODO(), core.GraphQLRequest{
		OperationName: "Test",
		Query:         "query Test { ping }",
	})

	if res.HasErrors() {
		t.Fatal(res.Errors)
	}

	if res.Data.(map[string]interface{})["ping"] != "Pong" {
		t.Fatal("Response did not contain the expected value.")
	}
}

type TestHook struct {
	PreCallCount  int
	PostCallCount int
}

func (hook *TestHook) PreExecute(ctx context.Context, req core.GraphQLRequest) context.Context {
	hook.PreCallCount += 1

	return ctx
}

func (hook *TestHook) PostExecute(ctx context.Context, req core.GraphQLRequest, res *graphql.Result) {
	hook.PostCallCount += 1
}

func TestExecutorHooks(t *testing.T) {
	testHook := &TestHook{}

	executor := core.GraphQLExecutor{
		Schema: SchemaType,
		RunBefore: []core.PreExecuteHook{
			testHook,
		},
		RunAfter: []core.PostExecuteHook{
			testHook,
		},
	}

	res := executor.Execute(context.TODO(), core.GraphQLRequest{
		OperationName: "Test",
		Query:         "query Test { ping }",
	})

	if res.HasErrors() {
		t.Fatal(res.Errors)
	}

	if testHook.PreCallCount != 1 {
		t.Fatal("Expected before hook to be called one time, was called " + strconv.Itoa(testHook.PreCallCount) + " time(s).")
	}

	if testHook.PostCallCount != 1 {
		t.Fatal("Expected after hook to be called one time, was called " + strconv.Itoa(testHook.PostCallCount) + " time(s).")
	}
}
