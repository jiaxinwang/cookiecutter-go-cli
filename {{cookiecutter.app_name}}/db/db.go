package db

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/'{{cookiecutter.use_db}}'"
)

// DB ...
var DB *gorm.DB

// InitDB ...
func InitDB(dbURI string) {
	var err error

	if DB, err = gorm.Open("postgres", dbURI); err != nil {
		logrus.Panic(err)
	}
}
