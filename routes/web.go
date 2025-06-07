package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Zyprush18/Scorely/handlers"
	"github.com/Zyprush18/Scorely/repository/database"
	"github.com/Zyprush18/Scorely/utils"
)

func RunApp()  {
	// connect database
	err := database.Connect()
	if err != nil {
		utils.Logfile(err.Error())
		log.Fatalln(err.Error())
	}

	http.HandleFunc("/add/role", handlers.AddRoles)

	fmt.Println("🚀 running on port: 8000")
	http.ListenAndServe(":8000",nil)
}