package main

import (
  "github.com/g0sy23/ik-app/internal/enterprise"
  "github.com/g0sy23/ik-app/internal/handler"
  "github.com/g0sy23/ik-app/internal/repository"
  "github.com/g0sy23/ik-app/internal/server"
  "github.com/sirupsen/logrus"
  "github.com/spf13/viper"
)

func main() {
  logrus.SetFormatter(new(logrus.JSONFormatter))
  if err := initConfig(); err != nil {
    logrus.Fatalf("error on initializing config - '%s'", err.Error())
  }

  database, err := ik_repository.NewPostgresDB(
    ik_repository.Config{
      Host:     viper.GetString("db.host"),
      Port:     viper.GetString("db.port"),
      Username: viper.GetString("db.username"),
      Password: viper.GetString("db.password"),
      DBName:   viper.GetString("db.name"),
      SSLMode:  viper.GetString("db.sslmode"),
    },
  )
  if err != nil {
    logrus.Fatalf("error on initializing postgres - '%s'", err.Error())
  }

  repository := ik_repository.NewRepository(database)
  enterprise := ik_enterprise.NewEnterprise(repository)
  handler    := ik_handler.NewHandler(enterprise)
  server     := ik_server.NewServer(handler)
  
  if err := server.Run(viper.GetString("port")); err != nil {
    logrus.Fatalf("error on initializing server - '%s'", err.Error())
  }
}

func initConfig() error {
  viper.AddConfigPath("configs")
  viper.SetConfigName("config")
  return viper.ReadInConfig()
}