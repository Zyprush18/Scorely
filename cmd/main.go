package main

import (
	"github.com/Zyprush18/Scorely/config"
	"github.com/Zyprush18/Scorely/routes"
)

func main(){
	config.Configfunc()
	routes.RunApp()
}