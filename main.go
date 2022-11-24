package main

import (
	"GoMicroservice/database"
	"GoMicroservice/router"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type config struct {
	Port             string `mapstructure:"port"`
	ConnectionString string `mapstructure:"connection_string"`
}

var AppConfig *config

func main() {
	// load configurations from config.json
	LoadAppConfig()

	// connect to database and run migration
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	// initialize mux router
	muxRouter := mux.NewRouter().StrictSlash(true)

	// register routes
	router.RegisterUserRoutes(muxRouter)

	// start go server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), muxRouter))

}

func LoadAppConfig() {
	log.Println("Loading Server Configurations")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
