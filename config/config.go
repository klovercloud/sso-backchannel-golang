package config

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var ClientId string
var ClientSecret string
var AuthorizeURI string
var RedirectURI string
var RunMode string
var ServerPort int

func InitEnVars() error {
	//checking runMode
	RunMode = os.Getenv("RUN_MODE")
	if RunMode == "" {
		RunMode = DEVELOP
	}
	var err error
	log.Println("RUN MODE:", RunMode)

	//loading envArs from .env file is runMode != PRODUCTION
	if RunMode != PRODUCTION {
		err = godotenv.Load()
		if err != nil {
			log.Println("ERROR: ", err.Error())
			return err
		}
	}
	var boolVal bool
	ClientId, boolVal = os.LookupEnv("CLIENT_ID")
	if boolVal == false {
		return errors.New("CLIENT_ID not fount in EnVars")
	}

	ClientSecret, boolVal = os.LookupEnv("CLIENT_SECRET")
	if boolVal == false {
		return errors.New("CLIENT_SECRET not fount in EnVars")
	}

	var serverPortStr string
	serverPortStr, boolVal = os.LookupEnv("SERVER_PORT")
	if boolVal == false {
		return errors.New("SERVER_PORT not fount in EnVars")
	}
	ServerPort, err = strconv.Atoi(serverPortStr)
	if err != nil {
		return err
	}

	AuthorizeURI, boolVal = os.LookupEnv("AUTHORIZE_URI")
	if boolVal == false {
		return errors.New("AUTHORIZE_URI not fount in EnVars")
	}

	RedirectURI, boolVal = os.LookupEnv("REDIRECT_URI")
	if boolVal == false {
		return errors.New("REDIRECT_URI not fount in EnVars")
	}

	return nil
}
