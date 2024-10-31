package reader

import "NBodySim/internal/mathutils/vector"

type Edge struct {
	V1, V2 int
}

type Polygon struct {
	Vertexes []int
}

type PolygonObject struct {
	Vertexes []vector.Vector3d
	Edges    []Edge
	Polygons []Polygon
}
