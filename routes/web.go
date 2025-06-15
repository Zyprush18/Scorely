package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Zyprush18/Scorely/database"
	"github.com/Zyprush18/Scorely/handlers/role"
	"github.com/Zyprush18/Scorely/handlers/user"
	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/repository/reporole"
	"github.com/Zyprush18/Scorely/repository/repouser"
	"github.com/Zyprush18/Scorely/service/servicerole"
	"github.com/Zyprush18/Scorely/service/userservice"
)

func RunApp() {
	pathlog := "./../log/app.log"
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
	adminMux.HandleFunc("/add/role", roleHandler.AddRoles)
	adminMux.HandleFunc("/role/{id}", roleHandler.Show)
	adminMux.HandleFunc("/role/{id}/update", roleHandler.Update)
	adminMux.HandleFunc("/role/{id}/delete", roleHandler.Delete)

	// user
	userRepo := repouser.NewUserDatabase(initDb)
	userService := userservice.NewUserService(&userRepo)
	userhandler := user.NewHandlerUser(userService, initlog)

	// user route
	adminMux.HandleFunc("/add/user", userhandler.Create)
	adminMux.HandleFunc("/user/{id}", userhandler.Show)

	fmt.Println("🚀 running on: http://localhost:8000")
	http.ListenAndServe(":8000", adminMux)
}
