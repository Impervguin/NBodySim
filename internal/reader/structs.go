package reader

import "NBodySim/internal/vectormath"

type Edge struct {
	V1, V2 int
}

type Polygon struct {
	Vertexes []int
}

type PolygonObject struct {
	Vertexes []vectormath.Vector3d
	Edges    []Edge
	Polygons []Polygon
}
