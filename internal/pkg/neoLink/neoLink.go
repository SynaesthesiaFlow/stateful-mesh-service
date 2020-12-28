package neoLink

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Neo4j_connection struct {
	uri      string
	username string
	password string

	driver  neo4j.Driver
	session neo4j.Session
}

func (neo *Neo4j_connection) CreateNewConnection() {
	// NEO4J logic
	neo.uri = "bolt://localhost:7687"
	neo.username = "neo4j"
	neo.password = "password"

	var err error
	neo.driver, err = neo4j.NewDriver(neo.uri, neo4j.BasicAuth(neo.username, neo.password, ""))
	if err != nil {
		fmt.Println("an error HERE neo4j Driver")
	}
	//defer driver.Close()

	// Handle driver lifetime based on your application lifetime requirements  driver's lifetime is usually
	// bound by the application lifetime, which usually implies one driver instance per application

	neo.session = neo.driver.NewSession(neo4j.SessionConfig{})

	// ### Sessions
	// Sessions are shortlived, cheap to create and NOT thread safe. Typically create one or more sessions
	// per request in your web application. Make sure to call Close on the session when done.
	// For multidatabase support, set sessionConfig.DatabaseName to requested database
	// Session config will default to write mode, if only reads are to be used configure session for
	// read mode.
	//defer session.Close()
}


// ### Sessions
// Sessions are shortlived, cheap to create and NOT thread safe. Typically create one or more sessions
// per request in your web application. Make sure to call Close on the session when done.
// For multidatabase support, set sessionConfig.DatabaseName to requested database
// Session config will default to write mode, if only reads are to be used configure session for
// read mode.
//defer session.Close()

func (neo *Neo4j_connection) RunTestCypher(s string) error {
	result, err := neo.session.Run("CREATE (n:Item { id: $id, name: $name }) RETURN n.id, n.name", map[string]interface{}{
		"id":   21,
		"name": "Item 21",
	})
	if err != nil {
		return err
	}

	var record *neo4j.Record
	for result.NextRecord(&record) {
		fmt.Printf("Created Item with Id = '%d' and Name = '%s'\n", record.Values[0].(int64), record.Values[1].(string))
	}
	return result.Err()
}

func (neo *Neo4j_connection) CloseConnection() {
	neo.session.Close()
	neo.driver.Close()
}

