package main

import (
	"log"
	"os"

	"code.google.com/p/gcfg"
)

type Config struct {
	LogInfo struct {
		LogDir string
	}

	OracleDB struct {
		HostName   string
		PortNumber string
		UserName   string
		Password   string
		SID        string
	}
}

func GetConfig(confName string) (cfg Config, err error) {
	if _, err := os.Stat(confName); os.IsNotExist(err) {
		log.Println(err.Error())

		return cfg, err
	}

	err = gcfg.ReadFileInto(&cfg, confName)
	if err != nil {
		log.Fatal(err.Error())
	}

	return cfg, err
}
