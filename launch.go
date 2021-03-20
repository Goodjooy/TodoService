package main

import (
	"todo-web/apis"
	"todo-web/server/server"
)

func main() {
	server.NewServer()
	db:=server.InitDatabase()

	server.UseMid(server.JWTVerify)
	
	server.AddApplication(apis.TodoApiApplication(db))
	server.AddApplication(apis.UserApiApplication(db))
	server.AddApplication(apis.TeskApiApplication(db))

	server.Build(":8080")
}