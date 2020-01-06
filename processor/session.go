package processor

import (
	"github.com/gocql/gocql"
	"log"
)

func ConnectionStarter(contactPoint, keyspace string, port int) (session *gocql.Session) {
	clusterConfig := gocql.NewCluster(contactPoint)
	clusterConfig.Keyspace = keyspace
	clusterConfig.Port = port
	clusterConfig.Consistency = gocql.One
	session, _ = clusterConfig.CreateSession()

	log.Println("Connection Established with the following Properties:")
	log.Println(*clusterConfig)
	return
}
