package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Zyprush18/Scorely/database"
	"github.com/Zyprush18/Scorely/handlers/class"
	"github.com/Zyprush18/Scorely/handlers/level"
	"github.com/Zyprush18/Scorely/handlers/major"
	"github.com/Zyprush18/Scorely/handlers/role"
	"github.com/Zyprush18/Scorely/handlers/student"
	"github.com/Zyprush18/Scorely/handlers/teacher"
	"github.com/Zyprush18/Scorely/handlers/user"
	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/repository/repoclass"
	"github.com/Zyprush18/Scorely/repository/repolevel"
	"github.com/Zyprush18/Scorely/repository/repomajor"
	"github.com/Zyprush18/Scorely/repository/reporole"
	"github.com/Zyprush18/Scorely/repository/repostudent"
	"github.com/Zyprush18/Scorely/repository/repoteacher"
	"github.com/Zyprush18/Scorely/repository/repouser"
	"github.com/Zyprush18/Scorely/service/classservice"
	"github.com/Zyprush18/Scorely/service/majorservice"
	"github.com/Zyprush18/Scorely/service/servicelevel"
	"github.com/Zyprush18/Scorely/service/servicerole"
	"github.com/Zyprush18/Scorely/service/servicestudent"
	"github.com/Zyprush18/Scorely/service/serviceteacher"
	"github.com/Zyprush18/Scorely/service/userservice"
)

func RunApp() {
	pathlog := "./log/app.log"
	initlog:= helper.NewLogger(pathlog)
	// connect database
	initDb,err := database.Connect()
	if err != nil {
		initlog.Logfile(err.Error())
		log.Fatalln("Connection Refused")
	}

	adminMux := http.NewServeMux()

	// role
	roleRepo := reporole.RolesMysql(initDb)
	roleService := servicerole.NewRoleService(roleRepo)
	roleHandler := role.RoleHandler(roleService, initlog)

	// role route
	adminMux.HandleFunc("/api/role", roleHandler.GetRole)
	adminMux.HandleFunc("/api/role/add", roleHandler.AddRoles)
	adminMux.HandleFunc("/api/role/{id}", roleHandler.Show)
	adminMux.HandleFunc("/api/role/{id}/update", roleHandler.Update)
	adminMux.HandleFunc("/api/role/{id}/delete", roleHandler.Delete)

	// user
	userRepo := repouser.NewUserDatabase(initDb)
	userService := userservice.NewUserService(&userRepo)
	userhandler := user.NewHandlerUser(userService, initlog)

	// user route
	adminMux.HandleFunc("/api/user", userhandler.GetAllUser)
	adminMux.HandleFunc("/api/user/add", userhandler.Create)
	adminMux.HandleFunc("/api/user/{id}", userhandler.Show)
	adminMux.HandleFunc("/api/user/{id}/update", userhandler.Update)
	adminMux.HandleFunc("/api/user/{id}/delete", userhandler.Delete)

	// major
	majorrepo := repomajor.ConnectDb(initDb)
	majorservice := majorservice.RepoMajorConn(&majorrepo)
	hanldermajor := major.Handlers(majorservice, initlog)

	// major route
	adminMux.HandleFunc("/api/major", hanldermajor.GetAllData)
	adminMux.HandleFunc("/api/major/add",hanldermajor.Create)
	adminMux.HandleFunc("/api/major/{id}",hanldermajor.Show)
	adminMux.HandleFunc("/api/major/{id}/update",hanldermajor.Updated)
	adminMux.HandleFunc("/api/major/{id}/delete",hanldermajor.Deleted)

	// level
	levelrepo := repolevel.ConnectDb(initDb)
	levelservice := servicelevel.ConnectRepo(&levelrepo)
	handlerlevel := level.ConnectService(levelservice, initlog)

	// route level
	adminMux.HandleFunc("/api/level", handlerlevel.GetAll)
	adminMux.HandleFunc("/api/level/add", handlerlevel.Create)
	adminMux.HandleFunc("/api/level/{id}", handlerlevel.Show)
	adminMux.HandleFunc("/api/level/{id}/update", handlerlevel.Update)
	adminMux.HandleFunc("/api/level/{id}/delete", handlerlevel.Delete)

	// class
	classrepo := repoclass.ConnectDb(initDb)
	serviceclass := classservice.NewClassService(&classrepo)
	handlerclass := class.NewHandlerClass(serviceclass, initlog)

	// route class
	adminMux.HandleFunc("/api/class", handlerclass.GetAll)
	adminMux.HandleFunc("/api/class/add", handlerclass.Create)
	adminMux.HandleFunc("/api/class/{id}", handlerclass.Show)
	adminMux.HandleFunc("/api/class/{id}/update", handlerclass.Update)
	adminMux.HandleFunc("/api/class/{id}/delete", handlerclass.Delete)

	// student
	studentrepo := repostudent.ConnectDb(initDb)
	servicestudent := servicestudent.NewServiceStudent(&studentrepo)
	handlerstudent := student.NewHandlerStudent(servicestudent,initlog)

	// route student
	adminMux.HandleFunc("/api/student", handlerstudent.GetAll)
	adminMux.HandleFunc("/api/student/add", handlerstudent.Create)
	adminMux.HandleFunc("/api/student/{id}", handlerstudent.Show)
	adminMux.HandleFunc("/api/student/{id}/update", handlerstudent.Update)
	adminMux.HandleFunc("/api/student/{id}/delete", handlerstudent.Delete)

	// teacher
	teacherrepo := repoteacher.ConnectDb(initDb)
	servicesteacher := serviceteacher.ConnectRepo(&teacherrepo)
	handlerteacher := teacher.ConnectService(servicesteacher,initlog)

	// route teacher
	adminMux.HandleFunc("/api/teacher", handlerteacher.GetAll)
	adminMux.HandleFunc("/api/teacher/add", handlerteacher.Create)
	
	fmt.Println("🚀 running on: http://localhost:8000")
	http.ListenAndServe(":8000", adminMux)
}
