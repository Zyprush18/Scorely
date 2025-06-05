package main

import (
	"fmt"
	"net/http"

	"github.com/Zyprush18/Scorely/repository/database"
)

func main(){
	// connect database
	database.Connect()

	fmt.Println("🚀 running on port : 8000")
	http.ListenAndServe(":8000",nil)
}