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
	pathlog := "./log/app.log"
	initlog:= helper.NewLogger(pathlog)
	// connect database
	initDb,err := database.Connect()
	if err != nil {
		initlog.Logfile(err.Error())
		log.Fatalln("Connection Refused")
	}

	// role
	roleRepo := reporole.RolesMysql(initDb)
	roleService := servicerole.RoleService(roleRepo)
	roleHandler := role.RoleHandler(roleService, initlog)

	// role route
	http.HandleFunc("/add/role", roleHandler.AddRoles)

	fmt.Println("🚀 running on port: 8000")
	http.ListenAndServe(":8000", nil)
}
