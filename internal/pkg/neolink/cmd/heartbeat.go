package main

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)


func neo4j_simple_test() error {
	// NEO4J logic
	uri := "bolt://localhost:7687"
	username := "neo4j"
	password := "password"

	//message, err := neo.HelloWorld(uri, username, password)
	//
	//if err != nil {
	//    fmt.Println("oops an error")
	//}
	//
	//fmt.Println(message)
	// Neo4j 4.0, defaults to no TLS therefore use bolt:// or neo4j://
	// Neo4j 3.5, defaults to self-signed cetificates, TLS on, therefore use bolt+ssc:// or neo4j+ssc://
	dbUri := uri
	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth(username, password, ""))

	// Handle driver lifetime based on your application lifetime requirements  driver's lifetime is usually
	// bound by the application lifetime, which usually implies one driver instance per application
	defer driver.Close()

	// Sessions are shortlived, cheap to create and NOT thread safe. Typically create one or more sessions
	// per request in your web application. Make sure to call Close on the session when done.
	// For multidatabase support, set sessionConfig.DatabaseName to requested database
	// Session config will default to write mode, if only reads are to be used configure session for
	// read mode.
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	result, err := session.Run("CREATE (n:Item { id: $id, name: $name }) RETURN n.id, n.name", map[string]interface{}{
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


func main() {
	err := neo4j_simple_test()
	if err != nil {
		fmt.Println("dam son")
	}
}
