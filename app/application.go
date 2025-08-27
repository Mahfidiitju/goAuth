package app

import (
	dbConfig "AuthInGo/config/db"
	config "AuthInGo/config/env"
	"AuthInGo/controller"
	repo "AuthInGo/db/repositories"
	"AuthInGo/router"
	"AuthInGo/service"
	"fmt"
	"net/http"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
}

func NewConfig() Config {
	port := config.GetString("PORT", "8080")
	return Config{
		Port: port,
	}
}

func NewApplication(config Config) *Application {
	return &Application{
		Config: config,
	}
}

func (a *Application) Start() {
	db, err := dbConfig.SetupDB()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return

	}
	ur := repo.NewUserRepository(db)
	us := service.NewUserService(ur)
	uc := controller.NewUserController(us)
	uRouter := router.NewUserRouter(uc)
	server := &http.Server{
		Addr:    a.Config.Port,
		Handler: router.SetupRouter(uRouter),
	}
	fmt.Println("Starting server on port", a.Config.Port)
	server.ListenAndServe()
}
