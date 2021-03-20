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

	app.AsignAddon(userConfime(db))

	fn, _ := IOC.AddIOC(testHandle)

	app.AsignViewer(
		manage.QuickNewViewer(
			"/:id/:value",
			db,
			manage.NewIOCHandle(manage.GET, fn)))

	return app
}

type TestPort struct {
	Req *http.Request
	Db  *gorm.DB

	PathValue IOC.Value `ioc:"from:path;to:uint;name:id;default:1"`
	UserValue IOC.Value `ioc:"from:context;to:raw;name:user"`
}

func testHandle(t TestPort) (dataHandle.Result, dataHandle.Result) {
	var v uint64
	var user models.UserModel
	t.PathValue.Get(&v)
	t.UserValue.Get(&user)

	return dataHandle.OkResult(v), dataHandle.OkResult(user)
}
