package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Zyprush18/Scorely/database"
	"github.com/Zyprush18/Scorely/handlers/level"
	"github.com/Zyprush18/Scorely/handlers/major"
	"github.com/Zyprush18/Scorely/handlers/role"
	"github.com/Zyprush18/Scorely/handlers/user"
	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/repository/repolevel"
	"github.com/Zyprush18/Scorely/repository/repomajor"
	"github.com/Zyprush18/Scorely/repository/reporole"
	"github.com/Zyprush18/Scorely/repository/repouser"
	"github.com/Zyprush18/Scorely/service/majorservice"
	"github.com/Zyprush18/Scorely/service/servicelevel"
	"github.com/Zyprush18/Scorely/service/servicerole"
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
	adminMux.HandleFunc("/role", roleHandler.GetRole)
	adminMux.HandleFunc("/role/add", roleHandler.AddRoles)
	adminMux.HandleFunc("/role/{id}", roleHandler.Show)
	adminMux.HandleFunc("/role/{id}/update", roleHandler.Update)
	adminMux.HandleFunc("/role/{id}/delete", roleHandler.Delete)

	// user
	userRepo := repouser.NewUserDatabase(initDb)
	userService := userservice.NewUserService(&userRepo)
	userhandler := user.NewHandlerUser(userService, initlog)

	// user route
	adminMux.HandleFunc("/user", userhandler.GetAllUser)
	adminMux.HandleFunc("/user/add", userhandler.Create)
	adminMux.HandleFunc("/user/{id}", userhandler.Show)
	adminMux.HandleFunc("/user/{id}/update", userhandler.Update)
	adminMux.HandleFunc("/user/{id}/delete", userhandler.Delete)

	// major
	majorrepo := repomajor.ConnectDb(initDb)
	majorservice := majorservice.RepoMajorConn(&majorrepo)
	hanldermajor := major.Handlers(majorservice, initlog)

	// major route
	adminMux.HandleFunc("/major", hanldermajor.GetAllData)
	adminMux.HandleFunc("/major/add",hanldermajor.Create)
	adminMux.HandleFunc("/major/{id}",hanldermajor.Show)
	adminMux.HandleFunc("/major/{id}/update",hanldermajor.Updated)
	adminMux.HandleFunc("/major/{id}/delete",hanldermajor.Deleted)

	// level
	levelrepo := repolevel.ConnectDb(initDb)
	levelservice := servicelevel.ConnectRepo(&levelrepo)
	handlerlevel := level.ConnectService(levelservice, initlog)

	// route level
	adminMux.HandleFunc("/level", handlerlevel.GetAll)
	adminMux.HandleFunc("/level/add", handlerlevel.Create)
	adminMux.HandleFunc("/level/{id}", handlerlevel.Show)

	fmt.Println("🚀 running on: http://localhost:8000")
	http.ListenAndServe(":8000", adminMux)
}
