package apis

import (
	"todo-web/dataHandle"
	"todo-web/err"
	"todo-web/models"
	"todo-web/server/IOC"
	"todo-web/server/server"

	"github.com/jinzhu/gorm"
)

type UserConfirmParm struct {
	ErrInfo  IOC.Value `ioc:"from:context;to:raw;name:err"`
	UserInfo IOC.Value `ioc:"from:context;to:raw;name:token`
}

func userConfime(parm UserConfirmParm, db *gorm.DB, context *IOC.ConxtextSeter) {
	var e err.Exception
	var u server.UserClaims

	parm.ErrInfo.Get(&e)
	parm.UserInfo.Get(&u)

	var user models.UserModel = models.FromUserClaims(u)
	var userResult []models.UserModel

	db.Where(&user).Find(&userResult)
	if len(userResult) == 0 {
		panic(dataHandle.FailureFuncResult(
			err.NotFoundUser, "user: "+user.EmailAddress))
	}
	context.Set("user", userResult[0])
}
