package graphql

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type user struct{
	ID string `json:"id"`
	Name string `json:"name"`
}

var data map[string]user

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:"User",
		Fields: graphql.Fields{
			"id" : &graphql.Field{
				Type: graphql.String,
			},
			"name":&graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user":&graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				//different file to contains func
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOk := p.Args["id"].(string)
					if isOk {
						return data[idQuery], nil
					}
					return nil, nil
				},
			},
		},
	},
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func Handler() gin.HandlerFunc {
	h:=handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		GraphiQL: true,
	})
	return func(c *gin.Context){
		h.ServeHTTP(c.Writer,c.Request)
	}
}