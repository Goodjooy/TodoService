package apis

import (
	"fmt"
	"regexp"
	"todo-web/dataHandle"
	"todo-web/err"
	"todo-web/models"
	"todo-web/server/manage"
	"todo-web/server/server"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var emailPattern = regexp.MustCompile("^([a-zA-Z0-9]+([-|.])?)+@([a-zA-Z0-9]+(-[a-zA-Z0-9]+)?\\.)+[a-zA-Z]{2,}$")

func UserApiApplication(db *gorm.DB) manage.Application {
	app := manage.NewApplication("/api/user", "user", "")
	app.AsignModels(&models.UserModel{})

	app.AsignViewer(manage.QuickNewViewer(
		"/login",
		db,
		manage.NewHandle(manage.POST, userLogin)))

	app.AsignViewer(manage.QuickNewViewer(
		"/sign-up", db,
		manage.NewHandle(manage.POST, signUp)))
	return app
}

func userLogin(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("paswd")

		encodedPassword := manage.DateSHA256Hash(password)

		var exampleUser models.UserModel
		var resultUsers []models.UserModel
		var user models.UserModel

		exampleUser.EmailAddress = email
		exampleUser.PassWord = encodedPassword

		db.Where(&exampleUser).Find(&resultUsers)
		if len(resultUsers) == 0 {
			c.JSON(400,
				dataHandle.FailureFuncResult(err.NotFoundUser,
					"user "+email+" not found"))
			return
		}
		user=resultUsers[0]
		token, e := server.GenerateToken(models.FromUser(user))
		if e != err.NoExcetion {
			c.JSON(400,
				dataHandle.FailureResult(e))
			return
		}
		c.SetCookie("token",token,3600,"/","localhost",false,false)
		c.JSON(200, dataHandle.OkResult(true))
	}
}

func signUp(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("paswd")
		name := c.PostForm("name")

		var exampleUser models.UserModel
		var userResult []models.UserModel
		var user models.UserModel

		exampleUser.EmailAddress = email
		db.Where(&exampleUser).Find(&userResult)

		if len(userResult) == 0 && emailPattern.MatchString(email) {
			user.EmailAddress = email
			user.PassWord = manage.DateSHA256Hash(password)
			user.Name = name

			db.Create(&user)

			c.JSON(200, dataHandle.OkResult(email))
		} else {
			c.JSON(400,
				dataHandle.FailureFuncResult(
					err.CreateNewUserFailure,
					fmt.Sprintf(
						"Email: %s | Name: %s", email, name)))
		}
	}
}
