package fixture

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Fixture struct {
	Points []neo4j.Point3D
}


func (f Fixture) SynchronizeNeo4j() error {
	// get each point in Fixture and create a node in Neo4j
	for i, point := range f.Points {
		fmt.Println(i, point)
	}
	return nil
}