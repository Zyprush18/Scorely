package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Zyprush18/Scorely/database"
	"github.com/Zyprush18/Scorely/handlers/role"
	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/repository/reporole"
	"github.com/Zyprush18/Scorely/service/servicerole"
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

	// role
	roleRepo := reporole.RolesMysql(initDb)
	roleService := servicerole.NewRoleService(roleRepo)
	roleHandler := role.RoleHandler(roleService, initlog)

	adminMux := http.NewServeMux()

	// role route
	adminMux.HandleFunc("/role", roleHandler.GetRole)
	adminMux.HandleFunc("/add/role", roleHandler.AddRoles)
	adminMux.HandleFunc("/role/{id}", roleHandler.Show)
	adminMux.HandleFunc("/role/{id}/update", roleHandler.Update)
	adminMux.HandleFunc("/role/{id}/delete", roleHandler.Delete)

	fmt.Println("🚀 running on: http://localhost:8000")
	http.ListenAndServe(":8000", adminMux)
}
