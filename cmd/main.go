package main

import (
	"log"

	"github.com/g0sy23/ik-app"
	"github.com/g0sy23/ik-app/internal/enterprise"
	"github.com/g0sy23/ik-app/internal/handler"
	"github.com/g0sy23/ik-app/internal/repository"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error on initializing configs - %s", err.Error())
	}

	repository := ik_repository.NewRepository()
	enterprise := ik_enterprise.NewEnterprise(repository)
	handlers	 := ik_handler.NewHandler(enterprise)
	server		 := ik_app.NewServer(handlers)

	if err := server.Run(viper.GetString("port")); err != nil {
		log.Fatalf("errorn on initializing server - %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}