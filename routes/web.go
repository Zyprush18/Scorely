package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Zyprush18/Scorely/database"
	"github.com/Zyprush18/Scorely/handlers/auth"
	"github.com/Zyprush18/Scorely/handlers/class"
	"github.com/Zyprush18/Scorely/handlers/exam"
	"github.com/Zyprush18/Scorely/handlers/examquestion"
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
	"github.com/Zyprush18/Scorely/repository/repoexamquestions"
	"github.com/Zyprush18/Scorely/repository/repoexams"
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
	"github.com/Zyprush18/Scorely/service/serviceexam"
	"github.com/Zyprush18/Scorely/service/serviceexamquest"
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
	// daftar role
	adm := "admin"
	tch := "teacher"
	// sdn := "student"

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
	adminMux.Handle("/api/role", middleware.MiddlewareAuth(http.HandlerFunc(roleHandler.GetRole), adm))
	adminMux.Handle("/api/role/add", middleware.MiddlewareAuth(http.HandlerFunc(roleHandler.AddRoles), adm))
	adminMux.Handle("/api/role/{id}", middleware.MiddlewareAuth(http.HandlerFunc(roleHandler.Show), adm))
	adminMux.Handle("/api/role/{id}/update", middleware.MiddlewareAuth(http.HandlerFunc(roleHandler.Update), adm))
	adminMux.Handle("/api/role/{id}/delete", middleware.MiddlewareAuth(http.HandlerFunc(roleHandler.Delete), adm))

	// user
	userRepo := repouser.NewUserDatabase(initDb)
	userService := userservice.NewUserService(&userRepo)
	userhandler := user.NewHandlerUser(userService, initlog)

	// user route
	adminMux.Handle("/api/user", middleware.MiddlewareAuth(http.HandlerFunc(userhandler.GetAllUser), adm))
	adminMux.Handle("/api/user/add", middleware.MiddlewareAuth(http.HandlerFunc(userhandler.Create), adm))
	adminMux.Handle("/api/user/{id}", middleware.MiddlewareAuth(http.HandlerFunc(userhandler.Show), adm))
	adminMux.Handle("/api/user/{id}/update", middleware.MiddlewareAuth(http.HandlerFunc(userhandler.Update), adm))
	adminMux.Handle("/api/user/{id}/delete", middleware.MiddlewareAuth(http.HandlerFunc(userhandler.Delete), adm))

	// major
	majorrepo := repomajor.ConnectDb(initDb)
	majorservice := majorservice.RepoMajorConn(&majorrepo)
	hanldermajor := major.Handlers(majorservice, initlog)

	// major route
	adminMux.Handle("/api/major", middleware.MiddlewareAuth(http.HandlerFunc(hanldermajor.GetAllData), adm))
	adminMux.Handle("/api/major/add", middleware.MiddlewareAuth(http.HandlerFunc(hanldermajor.Create), adm))
	adminMux.Handle("/api/major/{id}", middleware.MiddlewareAuth(http.HandlerFunc(hanldermajor.Show), adm))
	adminMux.Handle("/api/major/{id}/update", middleware.MiddlewareAuth(http.HandlerFunc(hanldermajor.Updated), adm))
	adminMux.Handle("/api/major/{id}/delete", middleware.MiddlewareAuth(http.HandlerFunc(hanldermajor.Deleted), adm))

	// level
	levelrepo := repolevel.ConnectDb(initDb)
	levelservice := servicelevel.ConnectRepo(&levelrepo)
	handlerlevel := level.ConnectService(levelservice, initlog)

	// route level
	adminMux.Handle("/api/level", middleware.MiddlewareAuth(http.HandlerFunc(handlerlevel.GetAll), adm))
	adminMux.Handle("/api/level/add", middleware.MiddlewareAuth(http.HandlerFunc(handlerlevel.Create), adm))
	adminMux.Handle("/api/level/{id}", middleware.MiddlewareAuth(http.HandlerFunc(handlerlevel.Show), adm))
	adminMux.Handle("/api/level/{id}/update", middleware.MiddlewareAuth(http.HandlerFunc(handlerlevel.Update), adm))
	adminMux.Handle("/api/level/{id}/delete", middleware.MiddlewareAuth(http.HandlerFunc(handlerlevel.Delete), adm))

	// class
	classrepo := repoclass.ConnectDb(initDb)
	serviceclass := classservice.NewClassService(&classrepo)
	handlerclass := class.NewHandlerClass(serviceclass, initlog)

	// route class
	adminMux.Handle("/api/class", middleware.MiddlewareAuth(http.HandlerFunc(handlerclass.GetAll), adm))
	adminMux.Handle("/api/class/add", middleware.MiddlewareAuth(http.HandlerFunc(handlerclass.Create), adm))
	adminMux.Handle("/api/class/{id}", middleware.MiddlewareAuth(http.HandlerFunc(handlerclass.Show), adm))
	adminMux.Handle("/api/class/{id}/update", middleware.MiddlewareAuth(http.HandlerFunc(handlerclass.Update), adm))
	adminMux.Handle("/api/class/{id}/delete", middleware.MiddlewareAuth(http.HandlerFunc(handlerclass.Delete), adm))

	// student
	studentrepo := repostudent.ConnectDb(initDb)
	servicestudent := servicestudent.NewServiceStudent(&studentrepo)
	handlerstudent := student.NewHandlerStudent(servicestudent,initlog)

	// route student
	adminMux.Handle("/api/student", middleware.MiddlewareAuth(http.HandlerFunc(handlerstudent.GetAll), adm))
	adminMux.Handle("/api/student/add", middleware.MiddlewareAuth(http.HandlerFunc(handlerstudent.Create), adm))
	adminMux.Handle("/api/student/{id}", middleware.MiddlewareAuth(http.HandlerFunc(handlerstudent.Show), adm))
	adminMux.Handle("/api/student/{id}/update", middleware.MiddlewareAuth(http.HandlerFunc(handlerstudent.Update), adm))
	adminMux.Handle("/api/student/{id}/delete", middleware.MiddlewareAuth(http.HandlerFunc(handlerstudent.Delete), adm))

	// teacher
	teacherrepo := repoteacher.ConnectDb(initDb)
	servicesteacher := serviceteacher.ConnectRepo(&teacherrepo)
	handlerteacher := teacher.ConnectService(servicesteacher,initlog)

	// route teacher
	adminMux.Handle("/api/teacher", middleware.MiddlewareAuth(http.HandlerFunc(handlerteacher.GetAll), adm))
	adminMux.Handle("/api/teacher/add", middleware.MiddlewareAuth(http.HandlerFunc(handlerteacher.Create), adm))
	adminMux.Handle("/api/teacher/{id}", middleware.MiddlewareAuth(http.HandlerFunc(handlerteacher.Show), adm))
	adminMux.Handle("/api/teacher/{id}/update", middleware.MiddlewareAuth(http.HandlerFunc(handlerteacher.Update), adm))
	adminMux.Handle("/api/teacher/{id}/delete", middleware.MiddlewareAuth(http.HandlerFunc(handlerteacher.Delete), adm))

	// subject
	subjectrepo:= reposubject.ConnectDb(initDb)
	servicesubjects:=subjectservice.ConnectRepo(&subjectrepo)
	subjecthandler := subject.ConnectService(servicesubjects,initlog)

	// route subject
	adminMux.Handle("/api/subject", middleware.MiddlewareAuth(http.HandlerFunc(subjecthandler.GetAll), adm))
	adminMux.Handle("/api/subject/add", middleware.MiddlewareAuth(http.HandlerFunc(subjecthandler.Create), adm))
	adminMux.Handle("/api/subject/{id}", middleware.MiddlewareAuth(http.HandlerFunc(subjecthandler.Show), adm))
	adminMux.Handle("/api/subject/{id}/update", middleware.MiddlewareAuth(http.HandlerFunc(subjecthandler.Update), adm))
	adminMux.Handle("/api/subject/{id}/delete", middleware.MiddlewareAuth(http.HandlerFunc(subjecthandler.Delete), adm))

	// exams
	examrepo := repoexams.ConnectDb(initDb)
	examservice := serviceexam.ConnectRepo(&examrepo)
	handlerexam := exam.ConnServc(examservice,initlog)

	// route exam
	adminMux.Handle("/api/exam", middleware.MiddlewareAuth(http.HandlerFunc(handlerexam.GetALl), adm))
	adminMux.Handle("/api/teacher/exam", middleware.MiddlewareAuth(http.HandlerFunc(handlerexam.FindByIdTeacher), tch))
	adminMux.Handle("/api/exam/{subject_id}/add", middleware.MiddlewareAuth(http.HandlerFunc(handlerexam.Create), adm, tch))
	adminMux.Handle("/api/exam/{id}", middleware.MiddlewareAuth(http.HandlerFunc(handlerexam.Show), adm,tch))
	adminMux.Handle("/api/exam/{id}/update", middleware.MiddlewareAuth(http.HandlerFunc(handlerexam.Update), adm,tch))
	adminMux.Handle("/api/exam/{id}/delete", middleware.MiddlewareAuth(http.HandlerFunc(handlerexam.Delete), adm,tch))

	// exam question
	examquestrepo := repoexamquestions.ConnectDB(initDb)
	examquestservice := serviceexamquest.ConnectRepo(&examquestrepo)
	handlerexamquest := examquestion.ConnectService(examquestservice,initlog)

	// route exam question
	adminMux.Handle("/api/exam/{id_exam}/examquestion", middleware.MiddlewareAuth(http.HandlerFunc(handlerexamquest.GetAll), adm,tch))
	
	fmt.Println("🚀 running on: http://localhost:8000")
	http.ListenAndServe(":8000", adminMux)
}
