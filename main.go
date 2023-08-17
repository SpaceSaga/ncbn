package main

import (
	"fmt"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"ncbn/infrastucture"
	"ncbn/types"
	"ncbn/webapi"
	"net/http"
)

func main() {
	config, err := GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbService, err := InitializeServices(config)
	if err != nil {
		log.Fatal(err)
	}

	api := &webapi.WebAPI{
		DBService: dbService,
	}

	port := ":8080"

	println("Starting server..")
	println("Listening on [" + port + "]")
	http.HandleFunc(
		"/api/endpoint",
		api.HandlePostRequest,
	)

	// Add the Swagger middleware
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	//http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe(port, nil))
}

func InitializeServices(config *types.Config) (*infrastucture.DatabaseService, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s",
		config.DBAddress,
		config.DBUser,
		config.DBPassword,
		config.DBName,
	)

	databaseService := infrastucture.NewDatabaseServiceWithDriverName(connString, config.DBDriverName)
	return databaseService, nil
}

func GetConfig() (*types.Config, error) {
	viper.SetConfigFile("app-settings.json")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %s", err)
	}

	var config types.Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %s", err)
	}

	return &config, nil
}
