package manage

import (
	"todo-web/server/IOC"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const (
	GET     = "GET"
	POST    = "POST"
	PUT     = "PUT"
	PATCH   = "PATCH"
	DELETE  = "DELETE"
	OPTIONS = "OPTIONS"
)

type Method string

type Viewer struct {
	URLPattern     string
	SupportMethods []Method

	handle map[Method][]gin.HandlerFunc
}

type FeedBackGenerate func(*gorm.DB) gin.HandlerFunc

type Handle struct {
	method  Method
	handles []FeedBackGenerate
}

func NewHandle(method Method, handles ...FeedBackGenerate) Handle {
	return Handle{method: method, handles: handles}
}

func NewIOCHandle(method Method, fn ...IOC.FuncHandler) Handle {
	handle := Handle{method: method}
	for _, f := range fn {
		var t FeedBackGenerate = func(d *gorm.DB) gin.HandlerFunc {
			return IOC.DoIOC(f, d)
		}
		handle.handles = append(handle.handles, t)
	}
	return handle
}

func NewViewer(URLPattern string, db *gorm.DB) Viewer {
	V := Viewer{URLPattern: URLPattern}
	V.handle = make(map[Method][]gin.HandlerFunc)
	return V
}

func QuickNewViewer(URLPattern string, db *gorm.DB, handles ...Handle) Viewer {
	v := NewViewer(URLPattern, db)

	for _, handle := range handles {
		var temp []gin.HandlerFunc

		for _, fn := range handle.handles {
			temp = append(temp, fn(db))
		}
		v.AsgnMethod(handle.method, temp...)
	}
	return v
}

func (v *Viewer) AsgnMethod(method Method, handles ...gin.HandlerFunc) {
	v.SupportMethods = append(v.SupportMethods, method)

	v.handle[method] = append(v.handle[method], handles...)
}
