package IOC

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var extraHandler map[reflect.Type]func(*gin.Context, *gorm.DB) reflect.Value
var request = reflect.TypeOf(&http.Request{})
var database = reflect.TypeOf(&gorm.DB{})
var valueTypeName = reflect.TypeOf(Value{})
var templateMapPtr = reflect.TypeOf(&TenmplateData{})
var contextSeterPtr = reflect.TypeOf(&ConxtextSeter{})

func initExtraHandler() {
	extraHandler = make(map[reflect.Type]func(*gin.Context, *gorm.DB) reflect.Value)

	extraHandler[request] = handleReuest
	extraHandler[database] = handleDB
	extraHandler[contextSeterPtr] = handleContextSeter

}
func handleReuest(c *gin.Context, db *gorm.DB) reflect.Value {
	return reflect.ValueOf(c.Request)
}

func handleDB(c *gin.Context, db *gorm.DB) reflect.Value {
	return reflect.ValueOf(db)
}
func handleContextSeter(c *gin.Context, db *gorm.DB) reflect.Value {
	v := newConxtextSeter()
	return reflect.ValueOf(v)
}

func getHandler(t reflect.Type) func(*gin.Context, *gorm.DB) reflect.Value {
	f, ok := extraHandler[t]
	if !ok {
		panic(fmt.Errorf("not handler found for %s", t.String()))
	}
	return f
}
