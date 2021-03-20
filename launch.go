package main

import (
	"todo-web/apis"
	"todo-web/server/server"
)

func main() {
	server.NewServer()
	db:=server.InitDatabase()

	server.UseIOCMid(server.JWTVerifyIOC)
	
	server.AddApplication(apis.TodoApiApplication(db))
	server.AddApplication(apis.UserApiApplication(db))
	server.AddApplication(apis.TeskApiApplication(db))

	server.Build(":8080")
}