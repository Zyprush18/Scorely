package helper

import (
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

const (
	Success            = http.StatusOK
	Created            = http.StatusCreated
	BadRequest         = http.StatusBadRequest
	Notfound           = http.StatusNotFound
	Conflict           = http.StatusConflict
	Forbidden          = http.StatusForbidden
	UnprocessbleEntity = http.StatusUnprocessableEntity
	Unauthorized       = http.StatusUnauthorized
	MethodNotAllowed   = http.StatusMethodNotAllowed
	InternalServError  = http.StatusInternalServerError
)

const (
	Gets   = http.MethodGet
	Post   = http.MethodPost
	Put    = http.MethodPut
	Delete = http.MethodDelete
)

// custom type for middleware
type ctxKey string

const KeyUserID ctxKey = "id_teacher"
const KeyCodeRole ctxKey = "role_code"

// struct message
type Messages struct {
	Message    string `json:"message,omitempty"`
	Data       any    `json:"data,omitempty"`
	Token      string `json:"token,omitempty"`
	Errors     string `json:"error,omitempty"`
	Fields     any    `json:"field,omitempty"`
	Pagination *Pag   `json:"pagination,omitempty"`
}

type Pag struct {
	Page      int    `json:"page"`
	TotalData int    `json:"total_data"`
	PerPage   int64  `json:"perpage"`
	Totalpage int    `json:"total_page"`
	Prev      string `json:"prev"`
	Next      string `json:"next"`
}

// createdat and updatedat struct
type Models struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"default:null"`
}

type Loggers interface {
	Logfile(logs string)
}

type Logger struct {
	path string
}

func NewLogger(pathfile string) Loggers {
	return Logger{path: pathfile}
}

// added log
func (l Logger) Logfile(logs string) {
	file, err := os.OpenFile(l.path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
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
				return errors.New(v.Field() + " is " + v.Tag() + " " + v.Param())
			}
		}
	}

	return nil
}

func IsDuplicateEntryError(err error) bool {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		if mysqlErr.Number == 1062 {
			return true
		}
	}
	return false
}

func QueryParam(r *http.Request, perpages int) (page, perpage int, sort, search string, err error) {
	search = r.URL.Query().Get("search")
	p := r.URL.Query().Get("page")
	page = 1
	if p != "" {
		page, err = strconv.Atoi(p)
		if err != nil {
			return 0, 0, "", "", err
		}
	}

	sort = r.URL.Query().Get("sort")
	if sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	return page, perpages, sort, search, nil

}

func Paginations(page, perpage, count int) *Pag {
	var next int
	prev := page - 1
	if prev <= 0 {
		prev = 0
	}
	totalpage := int(math.Ceil(float64(count) / float64(perpage)))
	if page < totalpage {
		next = page + 1
	}

	return &Pag{
		Page:      page,
		TotalData: int(count),
		PerPage:   int64(perpage),
		Totalpage: totalpage,
		Prev:      fmt.Sprintf("/role?page=%d", prev),
		Next:      fmt.Sprintf("/role?page=%d", next),
	}
}

func HashingPassword(password string) string {
	passhash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		Loggers.Logfile(Logger{}, err.Error())
	}
	return string(passhash)
}

func DecryptPassword(passhash, passreq string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(passhash), []byte(passreq)); err != nil {
		return errors.New("invalid_pw")
	}

	return nil
}
