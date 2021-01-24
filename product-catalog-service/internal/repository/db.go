package repository

import (
	"time"

	"gopkg.in/mgo.v2"
)

// DBConnection defines connection detailes
type DBConnection struct {
	session *mgo.Session
}

// NewConnection establishes a new connection to a particular mongo database
func NewConnection(host string, dbName string) (conn *DBConnection) {

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

	conn = &DBConnection{session}
	return conn
}

// Use retrieves a connection to a mongo collection
func (conn *DBConnection) Use(dbName, docName string) (collection *mgo.Collection) {
	return conn.session.DB(dbName).C(docName)
}

// Close handles closing a database connection session
func (conn *DBConnection) Close() {
	conn.session.Close()
	return
}
