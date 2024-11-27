package main

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"NBodySim/internal/reader"
	"fmt"
	"image/color"
)

func main() {
	read, err := reader.NewObjReader("./models/20_icosahedron.obj")
	if err != nil {
		panic(err)
	}
	po, err := read.ReadPolygonObject()
	if err != nil {
		panic(err)
	}

	po.Normals = make([]vector.Vector3d, 0, len(po.Polygons))

	for _, face := range po.Polygons {
		v1 := po.Vertexes[face.Vertexes[0]]
		v2 := po.Vertexes[face.Vertexes[1]]
		v3 := po.Vertexes[face.Vertexes[2]]
		p := object.NewPolygon(&v1, &v2, &v3, color.White)
		normal := p.GetNormal()
		nv := normal.ToVector()
		nv.Normalize()
		reference := vector.MultiplyVectorScalar(&v1, -1)
		if reference.Dot(&nv) < 0 { // Outer normal
			nv.MultiplyScalar(-1)
		}
		po.Normals = append(po.Normals, nv)
	}

	// output

	for _, v := range po.Vertexes {
		fmt.Printf("v %.4f %.4f %.4f\n", v.X, v.Y, v.Z)
	}

	for _, n := range po.Normals {
		fmt.Printf("vn %.4f %.4f %.4f\n", n.X, n.Y, n.Z)
	}
	for i, face := range po.Polygons {
		fmt.Printf("f %d//%d %d//%d %d//%d\n",
			face.Vertexes[0]+1, i+1,
			face.Vertexes[1]+1, i+1,
			face.Vertexes[2]+1, i+1,
		)
	}

}
