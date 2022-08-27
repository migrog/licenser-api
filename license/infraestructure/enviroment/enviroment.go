package enviroment

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var (
	MongoConnectionString = ""
	LicenseDatabase       = ""
	LicenseCollection     = ""
	UserBaseURL           = ""
	UserRoute             = ""
	ProductBaseURL        = ""
	ProductRoute          = ""
)

func NewSettings(env string) {
	var err error

	viper.SetConfigName(fmt.Sprintf("setting.%s", env))
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	MongoConnectionString = viper.GetString("mongo.connectionString")
	LicenseDatabase = viper.GetString("mongo.database")
	LicenseCollection = viper.GetString("mongo.licenseCollection")

	UserBaseURL = viper.GetString("userApi.baseURL")
	UserRoute = viper.GetString("userApi.users")

	ProductBaseURL = viper.GetString("productApi.baseURL")
	ProductRoute = viper.GetString("productApi.products")
}
