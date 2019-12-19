package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init(){
	// connect to the cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	// TODO
}

func GetSession()*gocql.Session{
	return session
}