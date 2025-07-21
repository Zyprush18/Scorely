package config

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)


func Configfunc()  {
	rootpath,err := filepath.Abs(".")
	if err != nil {
		log.Println("Not Found File .Env")
	}

	envPath := filepath.Join(rootpath,".env")
	if err:= godotenv.Load(envPath);err != nil {
		log.Println("Failed Load .Env")
	}
}