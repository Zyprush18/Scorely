package test

import (
	"testing"

	"github.com/Zyprush18/Scorely/repository/database"
	"github.com/stretchr/testify/assert"
)


func TestConnectD(t *testing.T)  {
	err_db := database.Connect()

	assert.NoError(t, err_db)
	assert.Nil(t,  err_db)
}