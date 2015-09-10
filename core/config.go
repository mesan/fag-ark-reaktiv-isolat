package core

import (
	"encoding/json"
	"fmt"
	. "github.com/goarne/logging"
	"io/ioutil"
	"os"
)

type AppConfig struct {
	Server  ServerConfig
	Logging LogConfig
}

type ServerConfig struct {
	Port int64
	Root string
}

func (a *AppConfig) ReadConfig(configFile string) {
	fileContent, e := ioutil.ReadFile(configFile)

	if e != nil {
		fmt.Println("Fant ikke configfil.", e)
		os.Exit(1)
	}

	json.Unmarshal(fileContent, &a)

}
