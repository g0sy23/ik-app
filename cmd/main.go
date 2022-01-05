package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/g0sy23/ik-app/internal/enterprise"
	"github.com/g0sy23/ik-app/internal/handler"
	"github.com/g0sy23/ik-app/internal/repository"
	"github.com/g0sy23/ik-app/internal/server"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error on initializing config - '%s'", err.Error())
	}

	database, err := initDatabase()
	if err != nil {
		logrus.Fatalf("error on initializing database - '%s'", err.Error())
	}

	repository := ik_repository.NewRepository(database)
	enterprise := ik_enterprise.NewEnterprise(repository)
	handler := ik_handler.NewHandler(enterprise)
	server := ik_server.NewServer(handler)

	go func() {
		if err = server.Run(viper.GetString("port")); err != nil {
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

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func initDatabase() (*sqlx.DB, error) {
	return ik_repository.NewPostgresDB(
		ik_repository.Config{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			Password: viper.GetString("db.password"),
			DBName:   viper.GetString("db.name"),
			SSLMode:  viper.GetString("db.sslmode"),
		},
	)
}
