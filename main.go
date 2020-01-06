package main

import (
	"gocassandra/model"
	"gocassandra/processor"
	"gocassandra/util"
	"log"
	"time"
)

func main() {
	cassandraConfig := util.ConfigReader("application.json")
	session := processor.
		ConnectionStarter(cassandraConfig.Contact_Point, cassandraConfig.Keyspace, cassandraConfig.PORT)
	defer session.Close()

	temperatureData := model.Temperature{
		Id:          "004",
		Timestamp:   time.Now().UTC().String(),
		Temperature: 98.6,
	}

	if v, err := util.Upsert(session, cassandraConfig, &temperatureData); err == nil {
		log.Println("=============" + v + "================")
	}

	resultSet := util.Query(session, cassandraConfig)
	for _, v := range resultSet {
		log.Println(v)
	}

	log.Println("===============CLEANING UP==================")
	var idDelete []string
	for _, v := range util.Query(session, cassandraConfig) {
		idDelete = append(idDelete, v.Id)
	}
	log.Println(util.Delete(session, cassandraConfig, idDelete))
}
