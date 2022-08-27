package mysql

import (
	"github.com/migrog/licenser-api/security/domain"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func Connect() {

	dsn := viper.GetString(`database.mysql.connection`)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}
	Conn = db

	db.AutoMigrate(domain.User{})
}
