package db

import (
	"time"

	"gopkg.in/mgo.v2"
)

// Connection defines connection detailes
type Connection struct {
	session *mgo.Session
}

// NewConnection establishes a new connection to a particular mongo database
func NewConnection(host string, dbName string) (conn *Connection) {

	info := &mgo.DialInfo{

		// Address of the database host
		Addrs: []string{host},

		// Timeout when a failure happens while connecting to the database
		Timeout: 60 * time.Second,

		// Database name
		Database: dbName,

		// Database credentials
		Username: "will be added here",
		Password: "will be added here",
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	conn = &Connection{session}
	return conn
}

// Use retrieves a connection to a mongo collection always use the default db
// since we use just one
func (conn *Connection) Use(docName string) (collection *mgo.Collection) {
	return conn.session.DB("").C(docName)
}

// Close handles closing a database connection session
func (conn *Connection) Close() {
	conn.session.Close()
	return
}
