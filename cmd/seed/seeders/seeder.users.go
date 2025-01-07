package seeders

import (
	"github.com/sirupsen/logrus"
	"github.com/snykk/grow-shop/internal/constants"
	"github.com/snykk/grow-shop/internal/datasources/records"
	"github.com/snykk/grow-shop/pkg/helpers"
	"github.com/snykk/grow-shop/pkg/logger"
)

var pass string
var UserData []records.Users

func init() {
	var err error
	pass, err = helpers.GenerateHash("12345")
	if err != nil {
		logger.Panic(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySeeder})
	}

	UserData = []records.Users{
		{
			UserName: "patrick star 7",
			Email:    "patrick@gmail.com",
			UserPassword: pass,
		},
		{
			UserName: "john doe",
			Email:    "johndoe@gmail.com",
			UserPassword: pass,
		},
	}
}
