package server

import (
	"fmt"
	"todo-web/server/manage"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Controller *gin.RouterGroup

var serverController *gin.Engine
var database *gorm.DB
func NewServer() {
	serverController = gin.Default()
}
func InitDatabase()*gorm.DB{
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charsetutf8&parseTime=True&loc=Local",  manage.SQLUser, manage.SQLPasswd, manage.DatabaseName))
	if err != nil {
		fmt.Printf("数据库连接失败，%s", err.Error())
		database=nil
	}
	 database= db.Debug()

	 return database
}

func UseMid(handle ...gin.HandlerFunc) {
	serverController.Use(handle...)
}

func AddApplication(app manage.Application){
	app.AsignApplication(serverController,database)
}

func NewController(path string) Controller {
	return serverController.Group(path)
}

func Build(address...string) error {
	return serverController.Run(address...)
}