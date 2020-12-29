package main

import (
	"fmt"
	"github.com/SynaesthesiaFlow/stateful-mesh-service/internal/pkg/fixture"
	"github.com/nargetdev/neostate"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"

)

func main() {
	neostate.Hello()

	f := fixture.Fixture{[]neo4j.Point3D{{1, 2, 3, 0}, {1134, 2, 3, 0}}}
	f.Points[0].X = 0.3
	fmt.Println(f)

	autogen := fixture.CreateStrip(0.5, 10)
	fmt.Println(autogen)

	autogen.SynchronizeNeo4j()

	neo := neostate.Neo4j_connection{
		//"bolt://localhost:7687",
		// "neo4j",
		//"password",
		//nil,
		//nil,
	}
	neo.CreateNewConnection()

	// create some cyphers, run them
	cypherString := "CREATE (\n" +
		"  p:Pixel\n" +
		"  {\n" +
		"    translation: point({x: $tx, y: $ty, z: $tz}),\n" +
		"    rotation:    point({ x: 2.3, y: 4.5, z: 2 }),\n" +
		"    color: [255, 255, 255]\n" +
		"  }\n" +
		") RETURN p.translation.x, p.translation.y, p.translation.z"

	cypherVars := map[string]interface{}{
		"tx": 4.314,
		"ty": 1.141234,
		"tz": 4.1234,
	}

	//cypherString := "CREATE (n:Item { id: $id, name: $name }) RETURN n.id, n.name"
	//
	//cypherVars := map[string]interface{}{
	//	"id":   21,
	//	"name": "Item 21",
	//}

	err := neo.RunTestCypher(cypherString, cypherVars)
	if err != nil {
		fmt.Println("error happened")
	}

	neo.CloseConnection()
}
