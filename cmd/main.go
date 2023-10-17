package main

import (
	"github.com/leykorn-orlan/todo-app"
	"github.com/leykorn-orlan/todo-app/pkg/handler"
	"github.com/leykorn-orlan/todo-app/pkg/repository"
	"github.com/leykorn-orlan/todo-app/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initialiation configs : %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while runnig http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
