package apis

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"todo-web/dataHandle"
	"todo-web/err"
	"todo-web/models"
	"todo-web/server/manage"
	"todo-web/server/server"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
)

func TodoApiApplication(db *gorm.DB) manage.Application {
	app := manage.NewApplication("/api/todo", "api", "")

	app.AsignAddon(userConfime(db))
	app.AsignViewer(
		manage.QuickNewViewer(
			"",
			db,
			manage.NewHandle(manage.GET , getAllTodos),
			manage.NewHandle(manage.POST , postNewTodo),
			manage.NewHandle(manage.PUT, putExistTodo),
			manage.NewHandle(manage.DELETE , deleteTodo)))

	app.AsignModels(&models.Todo{})
	return app
}

func userConfime(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		e, exist := c.Get("err")
		u, existu := c.Get("token")
		if exist && e != err.NoExcetion {
			c.JSON(http.StatusBadRequest, dataHandle.FailureResult(e.(err.Exception)))
			c.Abort()
			return
		} else if !existu {
			c.JSON(http.StatusBadRequest,
				dataHandle.FailureFuncResult(
					err.TokenFailure,
					"Not found User In token"))
			c.Abort()
			return
		}

		var user models.UserModel = models.FromUserClaims(*(u.(*server.UserClaims)))
		var userResult []models.UserModel

		db.Where(&user).Find(&userResult)
		if len(userResult) == 0 {
			c.JSON(http.StatusBadRequest, dataHandle.FailureFuncResult(
				err.NotFoundUser, "user: "+user.EmailAddress))
			c.Abort()
			return
		}
		c.Set("user", userResult[0])
	}
}

const (
	TODO_ALL = iota + 0
	TODO_DONE
	TODO_UNDONE
)

func getAllTodos(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exist := c.Get("user")
		if !exist {
			userNotExist(c)
		}
		pageSize, _ := strconv.Atoi(c.DefaultQuery("size", "0"))
		pageCount, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
		filter, _ := strconv.Atoi(c.DefaultQuery("filter", "0"))
		keyword:=c.DefaultQuery("key","")

		var todos []models.Todo
		var u = user.(models.UserModel)
		exampleTodo := models.Todo{User: u, UserID: u.ID}
		switch filter {
		case TODO_ALL:

		case TODO_DONE:
			exampleTodo.Status = (models.Done)

		case TODO_UNDONE:
			exampleTodo.Status = (models.UnDone)

		}
		var startPos int
		temp := db.Where(&exampleTodo)

		if keyword!=""{
			temp=temp.Where("`todos`.`title` LIKE ? or `todos`.`body` LIKE ?","%"+keyword+"%","%"+keyword+"%")
		}

		if pageCount != 0 && pageSize != 0 {
			startPos = pageSize * (pageCount - 1)
			temp.Limit(pageSize).Offset(startPos).Find(&todos)
		} else {
			temp.Find(&todos)
		}

		c.JSON(http.StatusOK, dataHandle.OkResult(todos))
	}
}

func postNewTodo(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		duration, _ := strconv.Atoi(c.PostForm("duration"))

		user, exist := c.Get("user")
		if !exist {
			userNotExist(c)
		}
		var u = user.(models.UserModel)
		var todo models.Todo
		e := c.MustBindWith(&todo, binding.Form)
		if e == nil {
			u.Todos = append(u.Todos, todo)
			todo.User = u
			todo.UserID = u.ID
			todo.Status = (models.UnDone)
			todo.DeathLine = time.Now().
				Add(time.Duration(duration * int(time.Hour) * 24))

			db.Save(&u)
			db.Save(&todo)

			c.JSON(http.StatusOK, dataHandle.OkResult(true))
			return
		}
		c.JSON(http.StatusBadRequest, dataHandle.FailureFuncResult(
			err.TargetParmsNotExist, e.Error()))
	}
}

func putExistTodo(d *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		foundTodo(c, d, func(t models.Todo) dataHandle.Result {
			if t.Status == (models.Done) {

				t.Status = (models.UnDone)
			} else {
				t.Status = (models.Done)
			}
			d.Save(t)
			return dataHandle.OkResult(true)
		})
	}
}
func deleteTodo(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		foundTodo(c, db, func(t models.Todo) dataHandle.Result {
			db.Delete(t)
			return dataHandle.OkResult(true)
		})
	}
}
func userNotExist(c *gin.Context) {
	c.JSON(http.StatusBadRequest,
		dataHandle.FailureFuncResult(
			err.NotFoundUser,
			"user not Found"))
	c.Abort()
}
func foundTodo(c *gin.Context, db *gorm.DB, opr func(models.Todo) dataHandle.Result) {
	user, exist := c.Get("user")
	if !exist {
		userNotExist(c)
	}
	var results []dataHandle.Result = make([]dataHandle.Result, 0)

	idsStr := c.Query("todoID")
	idStrs := strings.Split(idsStr, ",")

	var u = user.(models.UserModel)
	db.
		Where(&models.Todo{User: u, UserID: u.ID}).
		Find(&u.Todos)

	for _, v := range idStrs {
		id, e := strconv.Atoi(v)
		if e != nil {
			results = append(results, dataHandle.FailureFuncResult(err.TargetParmsNotExist,
				fmt.Sprintf("%s is not int", v)))
			continue
		}

		var todo models.Todo
		found := false

		for _, v := range u.Todos {
			if v.ID == uint(id) {
				todo = v
				found = true
				break
			}
		}
		if found {
			results = append(results, opr(todo))
		} else {
			results = append(results, dataHandle.FailureFuncResult(
				err.NotFoundToDo,
				fmt.Sprintf("Todo Id: %v not found in User: %s",
					id, u.EmailAddress)))
		}
	}
	c.JSON(200, dataHandle.OkResult(results))
}
