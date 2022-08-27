package enviroment

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var (
	MongoConnectionString = ""
	ProductDatabase       = ""
	ProductCollection     = ""
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
	ProductDatabase = viper.GetString("mongo.database")
	ProductCollection = viper.GetString("mongo.productCollection")
}
