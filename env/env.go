package env

import (
	"os"
	
	"github.com/irisnet/iris-api-server/modules/logger"
	"github.com/irisnet/iris-api-server/utils/constants"
)

var (
	ENV        string
	DbHost     string
	DbPort     string
	ServerPort string
)

func init()  {
	env, found := os.LookupEnv(constants.ENV_NAME_ENV)
	if !found {
		ENV = constants.ENV_DEV
		logger.Error.Printf("Environment variable %v is not set, default set to %v\n",
			constants.ENV_NAME_ENV, ENV)
	} else {
		ENV = env
		logger.Info.Printf("Environment has been set to %v\n", ENV)
	}
	
	dbHost, found := os.LookupEnv(constants.ENV_NAME_DB_HOST)
	if found {
		DbHost = dbHost
	}
	
	dbPort, found := os.LookupEnv(constants.ENV_NAME_DB_PORT)
	if found {
		DbPort = dbPort
	}
	
	serverPort, found := os.LookupEnv(constants.ENV_NAME_SERVER_PORT)
	if found {
		ServerPort = serverPort
	}
	
	
	
}