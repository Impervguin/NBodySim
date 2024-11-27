package reader

import "NBodySim/internal/mathutils/vector"

type Edge struct {
	V1, V2 int
}

type Polygon struct {
	Vertexes []int
	Normals  map[int]int // map from vertex index to its normals
}

type PolygonObject struct {
	Vertexes []vector.Vector3d
	Edges    []Edge
	Polygons []Polygon
	Normals  []vector.Vector3d
}
