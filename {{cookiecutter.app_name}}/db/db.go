package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/{{cookiecutter.use_db}}"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB ...
var DB *gorm.DB

// Connect ...
func Connect(dbURI string) {
	var err error

	var l logger.Interface
	l = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			// LogLevel:      logger.Info,
			LogLevel: logger.Silent,
			Colorful: true,
		})

	DB, err = gorm.Open({{cookiecutter.use_db}}.Open(dbURI), &gorm.Config{
		Logger: l,
	})

	if err != nil {
		fmt.Println(err.Error())
	}
}