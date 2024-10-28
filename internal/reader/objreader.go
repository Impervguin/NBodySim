package reader

import (
	"NBodySim/internal/vectormath"
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
	vertices := make([]vectormath.Vector3d, 0)
	edges := make([]Edge, 0)
	polygons := make([]Polygon, 0)

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
			vertices = append(vertices, *vectormath.NewVector3d(x, y, z))
		case "f":
			if len(parts) < 4 {
				return nil, fmt.Errorf("invalid face line: %s", str)
			}
			faceIndices := make([]int, 0)
			for _, indexStr := range parts[1:] {
				vParts := strings.Split(indexStr, "/")
				if len(vParts) == 0 {
					return nil, fmt.Errorf("invalid face index: %s", indexStr)
				}
				vIndex, err := strconv.Atoi(vParts[0])
				if err != nil {
					return nil, err
				}
				faceIndices = append(faceIndices, vIndex-1)
			}
			polygon := Polygon{faceIndices}
			polygons = append(polygons, polygon)
			for i := 0; i < len(faceIndices)-1; i++ {
				edges = append(edges, Edge{faceIndices[i], faceIndices[i+1]})
			}
			edges = append(edges, Edge{faceIndices[len(faceIndices)-1], faceIndices[0]})
		}
	}
	if err := scan.Err(); err != nil {
		return nil, err
	}
	po := &PolygonObject{
		Vertexes: vertices,
		Edges:    edges,
		Polygons: polygons,
	}
	return po, nil
}
