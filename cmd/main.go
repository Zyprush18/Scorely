package main

import (
	"github.com/Zyprush18/Scorely/config"
	"github.com/Zyprush18/Scorely/routes"
)

func main(){
	config.Configfunc()
	redisClient := config.ConnectRedis()
	defer redisClient.Close()
	routes.RunApp(redisClient)
}