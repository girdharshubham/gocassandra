package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Contact_Point string
	PORT          int
	Keyspace      string
	Table         string
}

func ConfigReader(fileName string) (cassandraDetails *Config) {
	jsonFile, _ := os.Open(fileName)
	defer jsonFile.Close()
	jsonByte, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(jsonByte, &cassandraDetails)
	return
}
