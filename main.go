package main

import (
	passwordCipher "auth-service/cipher"
	"auth-service/config"
	"auth-service/controllers"
	"auth-service/db"
	"auth-service/service"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	var c config.Config
	err := envconfig.Process("auth_service", &c)
	if err != nil {
		log.Fatal(err.Error())
	}
	startServer(c)
}

func startServer(config config.Config) {
	fmt.Println("%v", config)
	r := mux.NewRouter()
	userRepo := db.UserRepositoryFactory{}.CreateUserRepo(db.ConnectionSettings{
		DbUser:     config.DbUser,
		DbPassword: config.DbPassword,
		DbHost:     config.DbHost,
		DbPort:     config.DbPort,
		DbName:     config.DbName,
	})
	cipher := passwordCipher.BcryptCipher{}
	userCreator := service.UserCreator{
		UserRepo: userRepo,
		Cipher:   cipher,
	}
	loginChecker := service.LoginChecker{
		UserRepo: userRepo,
		Cipher:   cipher,
	}
	authController := controllers.NewAuthController(loginChecker, userCreator, validator.New())

	r.HandleFunc("/register", authController.Register).Methods("POST")
	r.HandleFunc("/login", authController.Login).Methods("POST")

	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), r); err != nil {
		log.Fatal(err)
	}
}
