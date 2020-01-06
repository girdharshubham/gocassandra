package util

import (
	"github.com/gocql/gocql"
	"gocassandra/model"
	"log"
)

func Upsert(session *gocql.Session, config *Config, temp *model.Temperature) (string, error) {
	if err := session.Query(
		"INSERT INTO "+config.Keyspace+"."+config.Table+"(id, timestamp, temperature) VALUES (?,?,?)",
		temp.Id, temp.Timestamp, temp.Temperature).
		Exec(); err != nil {
		return "Something Went Wrong!", err
	}

	return "Upsert Successful", nil
}

func Query(session *gocql.Session, config *Config) []model.Temperature {
	var tableData []model.Temperature
	interMap := map[string]interface{}{}
	iter := session.Query("SELECT * FROM data").Iter()
	for iter.MapScan(interMap) {
		tableData = append(tableData, model.Temperature{
			Id:          interMap["id"].(string),
			Timestamp:   interMap["timestamp"].(string),
			Temperature: interMap["temperature"].(float32),
		})
		interMap = map[string]interface{}{}
	}

	return tableData
}

func Delete(session *gocql.Session, config *Config, idArray []string) error {
	for _, v := range idArray {
		if err := session.Query("Delete FROM "+config.Keyspace+"."+config.Table+" WHERE id = ?", v).Exec(); err != nil {
			log.Fatal("Fatal Error")
			return err
		}
	}

	return nil
}
