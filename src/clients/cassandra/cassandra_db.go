package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	//cluster *gocql.ClusterConfig
	session *gocql.Session
)

func init() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
	//fmt.Println("cassandra connection sucessfully")
	//defer session.Close()
}

func GetSession() *gocql.Session {
	return session
}