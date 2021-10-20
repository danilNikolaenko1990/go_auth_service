package cmd

import (
	passwordCipher "auth-service/cipher"
	"auth-service/controllers"
	"auth-service/db"
	"auth-service/service"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
)

var HttpCmd = &cobra.Command{
	Use:   "http",
	Short: "HTTP Rest api",
	Run:   startHttp,
}

const (
	dbUserFlagName     = "db-user"
	dbPasswordFlagName = "db-password"
	dbHostFlagName     = "db-host"
	dbPortFlagName     = "db-port"
	dbNameFlagName     = "db-name"
	serverPortFlagName = "server-port"
)

var (
	serverPort string
	dbUser     string
	dbPassword string
	dbHost     string
	dbPort     string
	dbName     string
)

func init() {
	rootCmd.AddCommand(HttpCmd)

	HttpCmd.Flags().StringVar(&serverPort, serverPortFlagName, "8080", "")
	HttpCmd.MarkFlagRequired(serverPortFlagName)

	HttpCmd.Flags().StringVar(&dbUser, dbUserFlagName, "admin", "")
	HttpCmd.MarkFlagRequired(dbUserFlagName)

	HttpCmd.Flags().StringVar(&dbPassword, dbPasswordFlagName, "admin", "")
	HttpCmd.MarkFlagRequired(dbPasswordFlagName)

	HttpCmd.Flags().StringVar(&dbHost, dbHostFlagName, "localhost", "")
	HttpCmd.MarkFlagRequired(dbHostFlagName)

	HttpCmd.Flags().StringVar(&dbPort, dbPortFlagName, "5432", "")
	HttpCmd.MarkFlagRequired(dbHostFlagName)

	HttpCmd.Flags().StringVar(&dbName, dbNameFlagName, "postgres", "")
	HttpCmd.MarkFlagRequired(dbNameFlagName)
}

func startHttp(cmd *cobra.Command, args []string) {
	r := mux.NewRouter()
	userRepo := db.UserRepositoryFactory{}.CreateUserRepo(db.ConnectionSettings{
		DbUser:     dbUser,
		DbPassword: dbPassword,
		DbHost:     dbHost,
		DbPort:     dbPort,
		DbName:     dbName,
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

	if err := http.ListenAndServe(fmt.Sprintf(":%s", getPort()), r); err != nil {
		log.Fatal(err)
	}
}

func getPort() string {
	return serverPort
}
