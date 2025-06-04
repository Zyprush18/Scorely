package service

import (
	"fmt"
	"log"
	"os"
	"time"
)

// createdat and updatedat struct
type Models struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

// added log
func Logfile(logs string) {
	file, err := os.OpenFile("./log/app.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed Open file: %v", err)
	}

	defer file.Close()

	if _,err := file.WriteString(logs); err != nil {
		log.Fatalln("Failed to Add log")
	}
	fmt.Println("Success write log")
}