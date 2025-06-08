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
	// connect database
	initDb,err := database.Connect()
	if err != nil {
		helper.Logfile(err.Error())
		log.Fatalln("Connection Refused")
	}

	// role
	roleRepo := reporole.RolesMysql(initDb)
	roleService := servicerole.RoleService(roleRepo)
	roleHandler := role.RoleHandler(roleService)

	// role route
	http.HandleFunc("/add/role", roleHandler.AddRoles)

	fmt.Println("🚀 running on port: 8000")
	http.ListenAndServe(":8000", nil)
}
