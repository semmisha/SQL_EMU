package config

import (
	"github.com/sirupsen/logrus"
	"os"
)

func GetConfig(logger *logrus.Logger) map[string]string {

	config := make(map[string]string)
	config["url"] = os.Getenv("URL_MSSQL")
	config["port"] = os.Getenv("PORT_MSSQL")
	config["db"] = os.Getenv("DB_MSSQL")
	config["username"] = os.Getenv("USERNAME_MSSQL")
	config["password"] = os.Getenv("PASSWORD_MSSQL")

	for i, _ := range config {

		if len(config[i]) == 0 {
			logger.Fatalf("Config parameter: %v is empty", i)
		}

	}

	return config

}
