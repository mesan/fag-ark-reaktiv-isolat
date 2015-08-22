package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type AppConfig struct {
	Server  ServerConfig
	Logging LogConfig
}

type LogConfig struct {
	Filename         string
	Size             int64
	MaxNumberOfFiles int
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
