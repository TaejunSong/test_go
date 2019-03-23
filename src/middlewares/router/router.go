package router

import (
	"awesomeProject/src/api/graphql"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

//TODO: what is the init in golang
func init(){
	Router = gin.Default()
}

func SetRouter(){
	Router.POST("/graphql", graphql.Handler())
	Router.GET("/graphql", graphql.Handler())
}