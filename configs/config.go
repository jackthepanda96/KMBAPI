package configs

import (
	"encoding/json"
	"io"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

type ProgramConfig struct {
	ServerPort    int
	DBPort        int
	DBHost        string
	DBUser        string
	DBPass        string
	DBName        string
	Secret        string
	RefreshSecret string
}

func InitConfig() *ProgramConfig {
	var res = new(ProgramConfig)
	res = loadConfig()

	if res == nil {
		logrus.Fatal("Config : Cannot start program, failed to load configuration")
		return nil
	}

	return res
}

func loadConfig() *ProgramConfig {
	var res = new(ProgramConfig)

	// err := godotenv.Load(".env")

	// if err != nil {
	// 	logrus.Error("Config : Cannot load config file,", err.Error())
	// 	return nil
	// }

	if val, found := os.LookupEnv("SERVER"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : invalid port value,", err.Error())
			return nil
		}
		res.ServerPort = port
	}

	if val, found := os.LookupEnv("DBPORT"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : invalid db port value,", err.Error())
			return nil
		}
		res.DBPort = port
	}

	if val, found := os.LookupEnv("DBHOST"); found {
		res.DBHost = val
	}

	if val, found := os.LookupEnv("DBUSER"); found {
		res.DBUser = val
	}

	if val, found := os.LookupEnv("DBPASS"); found {
		res.DBPass = val
	}

	if val, found := os.LookupEnv("DBNAME"); found {
		res.DBName = val
	}

	if val, found := os.LookupEnv("SECRET"); found {
		res.Secret = val
	}

	if val, found := os.LookupEnv("REFSECRET"); found {
		res.RefreshSecret = val
	}

	if val, found := os.LookupEnv("GOOCREDS"); found {
		jsonBytes, err := json.Marshal(val)
		if err != nil {
			panic(err)
		}

		// Open the JSON file for writing.
		file, err := os.Create("credentials.json")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// Write the JSON data to the file.
		_, err = io.WriteString(file, string(jsonBytes))
		if err != nil {
			panic(err)
		}
	}

	return res

}
