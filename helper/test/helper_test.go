package test

import (
	"testing"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/stretchr/testify/assert"
)


func TestValidate(t *testing.T) {
	t.Run("Validation_success", func(t *testing.T) {
		roleReq := &request.Roles{
			NameRole: "Admin",
		}
		err := helper.ValidateForm(roleReq)

		assert.NoError(t, err)
		assert.Nil(t, err)
	})

	t.Run("Validation_Failed", func(t *testing.T) {
		roleReq := &request.Roles{
			NameRole: "",
		}
		err := helper.ValidateForm(roleReq)

		assert.Error(t, err)
		assert.NotNil(t, err)
	})
}