package builder

import (
	"NBodySim/internal/mathutils/normal"
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"NBodySim/internal/reader"
	"fmt"
)

type InnerNormalPolygonBuilder struct {
	obj       object.PolygonObject
	reader    reader.PolygonObjectReader
	readerObj *reader.PolygonObject
	center    *vector.Vector3d
	polygons  []*object.Polygon
	vertices  []*vector.Vector3d
}

type InnerNormalBuilderFactory struct{}

func (f *InnerNormalBuilderFactory) GetBuilder() PolygonObjectBuilder {
	return &InnerNormalPolygonBuilder{}
}

func (b *InnerNormalPolygonBuilder) setReader(reader reader.PolygonObjectReader) error {
	b.reader = reader
	return nil
}

func (b *InnerNormalPolygonBuilder) buildVertices() error {
	obj, err := b.reader.ReadPolygonObject()
	if err != nil {
		return err
	}
	b.readerObj = obj
	b.vertices = make([]*vector.Vector3d, 0, len(obj.Vertexes))
	for _, vertex := range obj.Vertexes {
		b.vertices = append(b.vertices, vector.NewVector3d(vertex.X, vertex.Y, vertex.Z))
	}
	return nil
}

func (b *InnerNormalPolygonBuilder) buildPolygon() error {
	if b.readerObj == nil {
		return fmt.Errorf("vertices not built yet")
	}
	polygons := make([]*object.Polygon, 0, len(b.readerObj.Polygons))
	for pind, polygon := range b.readerObj.Polygons {

		pverts := make([]*vector.Vector3d, 0, len(polygon.Vertexes))
		for _, index := range polygon.Vertexes {
			pverts = append(pverts, b.vertices[index])
		}
		if !CheckConvexPolygon(pverts) {
			return fmt.Errorf("non-convex polygon detected")
		}

		if len(polygon.Normals) == 0 {
			return fmt.Errorf("missing normals for polygon %v", pind)
		}

		// All normals must be equal
		normInd := 0
		for _, n := range polygon.Normals {
			if normInd == 0 {
				normInd = n
			}
			if n != normInd {
				return fmt.Errorf("inconsistent normals for polygon %v", pind)
			}
		}

		innerNormal := b.readerObj.Normals[normInd]

		if len(polygon.Vertexes) > 3 {
			start := polygon.Vertexes[0]
			color := DefaultObjectColor
			for i := 1; i < len(polygon.Vertexes)-1; i++ {
				n := normal.NewNormal(*b.vertices[polygon.Vertexes[start]], *vector.AddVectors(b.vertices[polygon.Vertexes[start]], &innerNormal))
				opolygon, err := object.NewPolygonInnerNormal(
					b.vertices[polygon.Vertexes[start]],
					b.vertices[polygon.Vertexes[i]],
					b.vertices[polygon.Vertexes[i+1]],
					n,
					color,
				)
				if err != nil {
					return err
				}

				polygons = append(polygons, opolygon)
			}
		} else if len(polygon.Vertexes) == 3 {
			n := normal.NewNormal(*b.vertices[polygon.Vertexes[0]], *vector.AddVectors(b.vertices[polygon.Vertexes[0]], &innerNormal))
			opolygon, err := object.NewPolygonInnerNormal(
				b.vertices[polygon.Vertexes[0]],
				b.vertices[polygon.Vertexes[1]],
				b.vertices[polygon.Vertexes[2]],
				n,
				DefaultObjectColor,
			)
			if err != nil {
				return err
			}
			polygons = append(polygons, opolygon)
		} else {
			return fmt.Errorf("invalid polygon with %d vertices", len(polygon.Vertexes))
		}
	}
	b.polygons = polygons
	return nil
}

func (b *InnerNormalPolygonBuilder) buildCenter() error {
	b.center = vector.NewVector3d(0, 0, 0)
	return nil
}

func (b *InnerNormalPolygonBuilder) getObject() (object.Object, error) {
	if b.readerObj == nil || b.vertices == nil {
		return nil, fmt.Errorf("vertices not built yet")
	}
	if b.center == nil {
		return nil, fmt.Errorf("center not built yet")
	}
	if b.polygons == nil {
		return nil, fmt.Errorf("polygons not built yet")
	}

	return object.NewPolygonObject(b.vertices, b.polygons, *b.center), nil
}
