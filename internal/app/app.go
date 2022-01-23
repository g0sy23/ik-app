package ik_app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/g0sy23/ik-app/internal/handler"
	"github.com/g0sy23/ik-app/internal/repository"
	"github.com/g0sy23/ik-app/internal/server"
	"github.com/g0sy23/ik-app/internal/services"
	"github.com/g0sy23/ik-app/pkg/database/postgres"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Run(configPath, configName string) {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(configPath, configName); err != nil {
		logrus.Fatalf("error on initializing config - '%s'", err.Error())
	}

	database, err := initDatabase()
	if err != nil {
		logrus.Fatalf("error on initializing database - '%s'", err.Error())
	}

	repository := ik_repository.New(database)
	services := ik_services.New(repository)
	handler := ik_handler.New(services)
	server := ik_server.New(handler)

	go func() {
		if err := server.Run(viper.GetString("port")); err != nil {
			logrus.Fatalf("error on initializing server - '%s'", err.Error())
		}
	}()

	logrus.Print("ik-app running")
	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	
	logrus.Print("ik-app shutting down")
	
	if err := server.Shutdown(); err != nil {
		logrus.Errorf("error on shutting down server - '%s'", err.Error())
	}
	
	if err := database.Close(); err != nil {
		logrus.Errorf("error on closing database - '%s'", err.Error())
	}
}

func initConfig(configPath, configName string) error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	return viper.ReadInConfig()
}

func initDatabase() (*postgresdb.PostgresDB, error) {
	return postgresdb.New(
		postgresdb.Config{
			DriverName: viper.GetString("db.driver"),
			Host:       viper.GetString("db.host"),
			Port:       viper.GetString("db.port"),
			Username:   viper.GetString("db.username"),
			Password:   viper.GetString("db.password"),
			DBName:     viper.GetString("db.name"),
			SSLMode:    viper.GetString("db.sslmode"),
		},
	)
}
