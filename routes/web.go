package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Zyprush18/Scorely/database"
	"github.com/Zyprush18/Scorely/handlers/auth"
	"github.com/Zyprush18/Scorely/handlers/class"
	"github.com/Zyprush18/Scorely/handlers/level"
	"github.com/Zyprush18/Scorely/handlers/major"
	"github.com/Zyprush18/Scorely/handlers/role"
	"github.com/Zyprush18/Scorely/handlers/student"
	"github.com/Zyprush18/Scorely/handlers/subject"
	"github.com/Zyprush18/Scorely/handlers/teacher"
	"github.com/Zyprush18/Scorely/handlers/user"
	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/middleware"
	"github.com/Zyprush18/Scorely/repository/repoauth"
	"github.com/Zyprush18/Scorely/repository/repoclass"
	"github.com/Zyprush18/Scorely/repository/repolevel"
	"github.com/Zyprush18/Scorely/repository/repomajor"
	"github.com/Zyprush18/Scorely/repository/reporole"
	"github.com/Zyprush18/Scorely/repository/repostudent"
	"github.com/Zyprush18/Scorely/repository/reposubject"
	"github.com/Zyprush18/Scorely/repository/repoteacher"
	"github.com/Zyprush18/Scorely/repository/repouser"
	"github.com/Zyprush18/Scorely/service/classservice"
	"github.com/Zyprush18/Scorely/service/majorservice"
	"github.com/Zyprush18/Scorely/service/serviceauth"
	"github.com/Zyprush18/Scorely/service/servicelevel"
	"github.com/Zyprush18/Scorely/service/servicerole"
	"github.com/Zyprush18/Scorely/service/servicestudent"
	"github.com/Zyprush18/Scorely/service/serviceteacher"
	"github.com/Zyprush18/Scorely/service/subjectservice"
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

	// login
	authrepo := repoauth.ConnectDb(initDb)
	authservice:= serviceauth.ConnectRepo(&authrepo)
	handlerauth := auth.ConnectService(authservice,initlog)

	// route login
	adminMux.HandleFunc("/api/login",handlerauth.Login)

	// role
	roleRepo := reporole.RolesMysql(initDb)
	roleService := servicerole.NewRoleService(roleRepo)
	roleHandler := role.RoleHandler(roleService, initlog)

	// role route
	adminMux.Handle("/api/role", middleware.MiddlewareAuthAdmin(http.HandlerFunc(roleHandler.GetRole)))
	adminMux.Handle("/api/role/add", middleware.MiddlewareAuthAdmin(http.HandlerFunc(roleHandler.AddRoles)))
	adminMux.Handle("/api/role/{id}", middleware.MiddlewareAuthAdmin(http.HandlerFunc(roleHandler.Show)))
	adminMux.Handle("/api/role/{id}/update", middleware.MiddlewareAuthAdmin(http.HandlerFunc(roleHandler.Update)))
	adminMux.Handle("/api/role/{id}/delete", middleware.MiddlewareAuthAdmin(http.HandlerFunc(roleHandler.Delete)))

	// user
	userRepo := repouser.NewUserDatabase(initDb)
	userService := userservice.NewUserService(&userRepo)
	userhandler := user.NewHandlerUser(userService, initlog)

	// user route
	adminMux.Handle("/api/user", middleware.MiddlewareAuthAdmin(http.HandlerFunc(userhandler.GetAllUser)))
	adminMux.Handle("/api/user/add", middleware.MiddlewareAuthAdmin(http.HandlerFunc(userhandler.Create)))
	adminMux.Handle("/api/user/{id}", middleware.MiddlewareAuthAdmin(http.HandlerFunc(userhandler.Show)))
	adminMux.Handle("/api/user/{id}/update", middleware.MiddlewareAuthAdmin(http.HandlerFunc(userhandler.Update)))
	adminMux.Handle("/api/user/{id}/delete", middleware.MiddlewareAuthAdmin(http.HandlerFunc(userhandler.Delete)))

	// major
	majorrepo := repomajor.ConnectDb(initDb)
	majorservice := majorservice.RepoMajorConn(&majorrepo)
	hanldermajor := major.Handlers(majorservice, initlog)

	// major route
	adminMux.Handle("/api/major", middleware.MiddlewareAuthAdmin(http.HandlerFunc(hanldermajor.GetAllData)))
	adminMux.Handle("/api/major/add", middleware.MiddlewareAuthAdmin(http.HandlerFunc(hanldermajor.Create)))
	adminMux.Handle("/api/major/{id}", middleware.MiddlewareAuthAdmin(http.HandlerFunc(hanldermajor.Show)))
	adminMux.Handle("/api/major/{id}/update", middleware.MiddlewareAuthAdmin(http.HandlerFunc(hanldermajor.Updated)))
	adminMux.Handle("/api/major/{id}/delete", middleware.MiddlewareAuthAdmin(http.HandlerFunc(hanldermajor.Deleted)))

	// level
	levelrepo := repolevel.ConnectDb(initDb)
	levelservice := servicelevel.ConnectRepo(&levelrepo)
	handlerlevel := level.ConnectService(levelservice, initlog)

	// route level
	adminMux.Handle("/api/level", middleware.MiddlewareAuthAdmin(http.HandlerFunc(handlerlevel.GetAll)))
	adminMux.Handle("/api/level/add", middleware.MiddlewareAuthAdmin(http.HandlerFunc(handlerlevel.Create)))
	adminMux.Handle("/api/level/{id}", middleware.MiddlewareAuthAdmin(http.HandlerFunc(handlerlevel.Show)))
	adminMux.Handle("/api/level/{id}/update", middleware.MiddlewareAuthAdmin(http.HandlerFunc(handlerlevel.Update)))
	adminMux.Handle("/api/level/{id}/delete", middleware.MiddlewareAuthAdmin(http.HandlerFunc(handlerlevel.Delete)))

	// class
	classrepo := repoclass.ConnectDb(initDb)
	serviceclass := classservice.NewClassService(&classrepo)
	handlerclass := class.NewHandlerClass(serviceclass, initlog)

	// route class
	adminMux.Handle("/api/class", middleware.MiddlewareAuthAdmin(http.HandlerFunc(handlerclass.GetAll)))
	adminMux.Handle("/api/class/add", middleware.MiddlewareAuthAdmin(http.HandlerFunc(handlerclass.Create)))
	adminMux.Handle("/api/class/{id}", middleware.MiddlewareAuthAdmin(http.HandlerFunc(handlerclass.Show)))
	adminMux.Handle("/api/class/{id}/update", middleware.MiddlewareAuthAdmin(http.HandlerFunc(handlerclass.Update)))
	adminMux.Handle("/api/class/{id}/delete", middleware.MiddlewareAuthAdmin(http.HandlerFunc(handlerclass.Delete)))

	// student
	studentrepo := repostudent.ConnectDb(initDb)
	servicestudent := servicestudent.NewServiceStudent(&studentrepo)
	handlerstudent := student.NewHandlerStudent(servicestudent,initlog)

	// route student
	adminMux.Handle("/api/student", middleware.MiddlewareAuthAdmin(http.HandlerFunc(handlerstudent.GetAll)))
	adminMux.Handle("/api/student/add", middleware.MiddlewareAuthAdmin(http.HandlerFunc(handlerstudent.Create)))
	adminMux.Handle("/api/student/{id}", middleware.MiddlewareAuthAdmin(http.HandlerFunc(handlerstudent.Show)))
	adminMux.Handle("/api/student/{id}/update", middleware.MiddlewareAuthAdmin(http.HandlerFunc(handlerstudent.Update)))
	adminMux.Handle("/api/student/{id}/delete", middleware.MiddlewareAuthAdmin(http.HandlerFunc(handlerstudent.Delete)))

	// teacher
	teacherrepo := repoteacher.ConnectDb(initDb)
	servicesteacher := serviceteacher.ConnectRepo(&teacherrepo)
	handlerteacher := teacher.ConnectService(servicesteacher,initlog)

	// route teacher
	adminMux.Handle("/api/teacher", middleware.MiddlewareAuthAdmin(http.HandlerFunc(handlerteacher.GetAll)))
	adminMux.Handle("/api/teacher/add", middleware.MiddlewareAuthAdmin(http.HandlerFunc(handlerteacher.Create)))
	adminMux.Handle("/api/teacher/{id}", middleware.MiddlewareAuthAdmin(http.HandlerFunc(handlerteacher.Show)))
	adminMux.Handle("/api/teacher/{id}/update", middleware.MiddlewareAuthAdmin(http.HandlerFunc(handlerteacher.Update)))
	adminMux.Handle("/api/teacher/{id}/delete", middleware.MiddlewareAuthAdmin(http.HandlerFunc(handlerteacher.Delete)))

	// subject
	subjectrepo:= reposubject.ConnectDb(initDb)
	servicesubjects:=subjectservice.ConnectRepo(&subjectrepo)
	subjecthandler := subject.ConnectService(servicesubjects,initlog)

	// route subject
	adminMux.Handle("/api/subject", middleware.MiddlewareAuthAdmin(http.HandlerFunc(subjecthandler.GetAll)))
	adminMux.Handle("/api/subject/add", http.HandlerFunc(subjecthandler.Create))
	adminMux.Handle("/api/subject/{id}", middleware.MiddlewareAuthAdmin(http.HandlerFunc(subjecthandler.Show)))
	adminMux.Handle("/api/subject/{id}/update", middleware.MiddlewareAuthAdmin(http.HandlerFunc(subjecthandler.Update)))
	adminMux.Handle("/api/subject/{id}/delete", middleware.MiddlewareAuthAdmin(http.HandlerFunc(subjecthandler.Delete)))
	
	fmt.Println("🚀 running on: http://localhost:8000")
	http.ListenAndServe(":8000", adminMux)
}
