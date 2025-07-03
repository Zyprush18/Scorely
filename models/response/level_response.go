package response
import "github.com/Zyprush18/Scorely/helper"

type Levels struct {
	IdLevel uint   `json:"id_level"`
	Level   string `json:"level"`
	helper.Models
}