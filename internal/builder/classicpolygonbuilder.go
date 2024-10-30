package builder

import (
	"NBodySim/internal/object"
	"NBodySim/internal/reader"
	"NBodySim/internal/vectormath"
	"fmt"
	"image/color"

	"math/rand"
)

type ClassicPolygonBuilder struct {
	obj       object.PolygonObject
	reader    reader.PolygonObjectReader
	readerObj *reader.PolygonObject
	center    *vectormath.Vector3d
	polygons  []*object.Polygon
	vertices  []*vectormath.Vector3d
}

type ClassicPolygonFactory struct{}

func (f *ClassicPolygonFactory) GetBuilder() PolygonObjectBuilder {
	return &ClassicPolygonBuilder{}
}

func (b *ClassicPolygonBuilder) setReader(reader reader.PolygonObjectReader) error {
	b.reader = reader
	return nil
}

func (b *ClassicPolygonBuilder) buildVertices() error {
	obj, err := b.reader.ReadPolygonObject()
	if err != nil {
		return err
	}
	b.readerObj = obj
	b.vertices = make([]*vectormath.Vector3d, 0, len(obj.Vertexes))
	for _, vertex := range obj.Vertexes {
		b.vertices = append(b.vertices, vectormath.NewVector3d(vertex.X, vertex.Y, vertex.Z))
	}
	return nil
}

func (b *ClassicPolygonBuilder) buildPolygon() error {
	if b.readerObj == nil {
		return fmt.Errorf("vertices not built yet")
	}
	polygons := make([]*object.Polygon, 0, len(b.readerObj.Polygons))
	for _, polygon := range b.readerObj.Polygons {
		if len(polygon.Vertexes) > 3 {
			start := polygon.Vertexes[0]
			color := color.RGBA{R: uint8(rand.Intn(256)), G: uint8(rand.Intn(256)), B: uint8(rand.Intn(256)), A: 255}
			// color := color.RGBA{R: 255, B: 0, G: 0, A: 255}
			for i := range polygon.Vertexes[1 : len(polygon.Vertexes)-1] {
				opolygon := object.NewPolygon(
					b.vertices[polygon.Vertexes[start]],
					b.vertices[polygon.Vertexes[i]],
					b.vertices[polygon.Vertexes[i+1]],
					color,
				)
				polygons = append(polygons, opolygon)
			}
		} else if len(polygon.Vertexes) == 3 {
			opolygon := object.NewPolygon(
				b.vertices[polygon.Vertexes[0]],
				b.vertices[polygon.Vertexes[1]],
				b.vertices[polygon.Vertexes[2]],
				color.RGBA{R: uint8(rand.Intn(256)), G: uint8(rand.Intn(256)), B: uint8(rand.Intn(256)), A: 255},
				// color.RGBA{R: 255, B: 0, G: 0, A: 255},
			)
			polygons = append(polygons, opolygon)
		} else {
			return fmt.Errorf("invalid polygon with %d vertices", len(polygon.Vertexes))
		}
	}
	b.polygons = polygons
	return nil
}

func (b *ClassicPolygonBuilder) buildCenter() error {
	b.center = vectormath.NewVector3d(0, 0, 0)
	return nil
}

func (b *ClassicPolygonBuilder) getObject() (object.Object, error) {
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
