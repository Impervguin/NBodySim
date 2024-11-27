package reader

import (
	"NBodySim/internal/mathutils/vector"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ObjReader struct {
	filename string
}

func NewObjReader(filename string) (*ObjReader, error) {
	return &ObjReader{filename: filename}, nil
}

func (or *ObjReader) ReadPolygonObject() (*PolygonObject, error) {
	f, err := os.Open(or.filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	vertices := make([]vector.Vector3d, 0)
	edges := make([]Edge, 0)
	polygons := make([]Polygon, 0)
	normals := make([]vector.Vector3d, 0)

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		str := strings.TrimSpace(scan.Text())

		if len(str) != 0 && str[0] == '#' {
			continue
		}
		parts := strings.Split(str, " ")
		if len(parts) == 0 {
			continue
		}
		command := parts[0]
		switch command {
		case "v":
			if len(parts) < 4 {
				return nil, fmt.Errorf("invalid vertex line: %s", str)
			}
			x, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				return nil, err
			}
			y, err := strconv.ParseFloat(parts[2], 64)
			if err != nil {
				return nil, err
			}
			z, err := strconv.ParseFloat(parts[3], 64)
			if err != nil {
				return nil, err
			}
			vertices = append(vertices, *vector.NewVector3d(x, y, z))
		case "vn":
			if len(parts) < 4 {
				return nil, fmt.Errorf("invalid normal line: %s", str)
			}
			x, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				return nil, err
			}
			y, err := strconv.ParseFloat(parts[2], 64)
			if err != nil {
				return nil, err
			}
			z, err := strconv.ParseFloat(parts[3], 64)
			if err != nil {
				return nil, err
			}
			normals = append(normals, *vector.NewVector3d(x, y, z))
		case "f":
			if len(parts) < 4 {
				return nil, fmt.Errorf("invalid face line: %s", str)
			}
			vertexIndices := make([]int, 0)
			normalIndices := make(map[int]int)
			for i, indexStr := range parts[1:] {
				vParts := strings.Split(indexStr, "/")
				if len(vParts) == 0 {
					return nil, fmt.Errorf("invalid face index: %s", indexStr)
				}
				vIndex, err := strconv.Atoi(vParts[0])
				if err != nil {
					return nil, err
				}
				if vIndex < 1 || vIndex > len(vertices) {
					return nil, fmt.Errorf("invalid vertex index: %d", vIndex)
				}
				vertexIndices = append(vertexIndices, vIndex-1)

				if len(vParts) >= 3 {
					nIndex, err := strconv.Atoi(vParts[2])
					if err != nil {
						return nil, err
					}
					if nIndex < 1 || nIndex > len(normals) {
						return nil, fmt.Errorf("invalid normal index: %d", nIndex)
					}
					normalIndices[i] = nIndex - 1
				}
			}
			polygon := Polygon{vertexIndices, normalIndices}
			polygons = append(polygons, polygon)
			for i := 0; i < len(vertexIndices)-1; i++ {
				edges = append(edges, Edge{vertexIndices[i], vertexIndices[i+1]})
			}
			edges = append(edges, Edge{vertexIndices[len(vertexIndices)-1], vertexIndices[0]})
		}
	}
	if err := scan.Err(); err != nil {
		return nil, err
	}
	po := &PolygonObject{
		Vertexes: vertices,
		Edges:    edges,
		Polygons: polygons,
		Normals:  normals,
	}
	return po, nil
}
