package main

import (
	"awesomeProject/src/middlewares/router"
)

//TODO: check jwt middleware connection
//TODO: check router connection
//TODO: check graphql connection
func main(){

	//schema
	r := router.Router
	router.SetRouter()
	//r.Use(jwt.JWTMiddleware())
	r.Run(":8080")
}





