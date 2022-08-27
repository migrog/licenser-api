package database

import "github.com/migrog/licenser-api/security/repository/mysql"

func Connect(driver string) {
	switch driver {
	case "mysql":
		mysql.Connect()
	default:
		mysql.Connect()
	}
}
