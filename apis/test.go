package apis

import (
	"net/http"
	"todo-web/dataHandle"
	"todo-web/models"
	"todo-web/server/IOC"
	"todo-web/server/manage"

	"github.com/jinzhu/gorm"
)

func TeskApiApplication(db *gorm.DB) manage.Application {
	app := manage.NewApplication("/api/test", "test", "")

	app.AsignAddonIOC(db,userConfime)

	app.AsignViewer(
		manage.QuickNewViewer(
			"/:id/:value",
			db,
			manage.NewIOCHandle(manage.GET, testHandle)))

	return app
}

type TestPort struct {
	Db *gorm.DB

	PathValue IOC.Value `ioc:"from:path;to:uint;name:id;default:1"`
	NameValue IOC.Value `ioc:"from:path;to:string;name:value;default:abab"`
	UserValue IOC.Value `ioc:"from:context;to:raw;name:user"`
}

func testHandle(t TestPort, c *IOC.ConxtextSeter, Req *http.Request) (dataHandle.Result, dataHandle.Result) {
	var v uint64
	var v2 string
	var user models.UserModel
	t.PathValue.Get(&v)
	t.UserValue.Get(&user)
	t.NameValue.Get(&v2)

	c.Set("ssss", v2)

	return dataHandle.OkResult(v), dataHandle.OkResult(user)
}
