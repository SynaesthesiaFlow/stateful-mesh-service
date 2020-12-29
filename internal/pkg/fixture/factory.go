package fixture

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func CreateStrip(spacing float64, numLEDs int) (result Fixture) {
	points := make([]neo4j.Point3D, numLEDs, numLEDs*2)
	for i := 0; i < numLEDs; i++ {
		points[i] = neo4j.Point3D{float64(i)*0.1, float64(i)/0.9, float64(i), 0}
		fmt.Println(points[i])
	}
	return  Fixture{points}
}