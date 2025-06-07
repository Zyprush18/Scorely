package utils

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

const (
	Success          = http.StatusOK
	Created          = http.StatusCreated
	BadRequest       = http.StatusBadRequest
	MethodNotAllowed = http.StatusMethodNotAllowed
)

const (
	Gets = http.MethodGet
	Post = http.MethodPost
)

type Messages struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"error,omitempty"`
}

// createdat and updatedat struct
type Models struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"default:null"`
}

// added log
func Logfile(logs string) {
	file, err := os.OpenFile("./log/app.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed Open file: %v", err)
	}

	defer file.Close()

	if _, err := file.WriteString(logs + "\n"); err != nil {
		log.Fatalln("Failed to Add log")
	}
	fmt.Println("Success write log")
}

func ValidateForm(data interface{}) error {
	validate := validator.New()

	// ambil field berdasarkan format json
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.Split(field.Tag.Get("json"), ",")[0]
		if name == "-" {
			return ""
		}

		return name
	})

	if err := validate.Struct(data); err != nil {
		var validateerr validator.ValidationErrors
		if errors.As(err, &validateerr) {
			for _, v := range validateerr {
				return errors.New(v.Field() + " is " + v.Tag())
			}
		}
	}

	return nil
}
