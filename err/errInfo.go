package err

import ()

type Exception struct {
	Code         uint
	Message      string
	ExtraMessage string
}

func newException(code uint, message string) func(string) Exception {
	e := Exception{Code: code, Message: message}

	f := func(extra string) Exception {
		e.ExtraMessage = extra
		return e
	}
	return f
}

var (
	//souce not found 1XX

	//not Found In DataBase
	NotFoundInDataBase = newException(101, "Not Found Item In DataBase")
	//Not Found User
	NotFoundUser=newException(102,"Not Found User In User Set")
	//Todo Not Found
	NotFoundToDo=newException(103,"Target Todo Not Found")
	//API Port Not Found
	NotFoundAPI=newException(104,"Target API Port Not Exist")

	
	//Bad Request 2XX
	TargetParmNotExist=newException(201,"target Parm Not provide")
	
	//authentication 3XX
	AuthenticationFailure=newException(301 ,"Failure authentication User")
	PermissionDenied=newException(302,"Permission Denied")
	AccessDenied=newException(303,"Access Denied")

)
