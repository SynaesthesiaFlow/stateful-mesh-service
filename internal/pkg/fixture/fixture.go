package fixture

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

type Fixture struct {
	Points [12]neo4j.Point3D
}

