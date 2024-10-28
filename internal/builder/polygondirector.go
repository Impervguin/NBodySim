package builder

import (
	"NBodySim/internal/object"
	"NBodySim/internal/reader"
)

type PolygonObjectDirector struct {
	factory PolygonObjectFactory
	reader  reader.PolygonObjectReader
}

type PolygonObjectFactory interface {
	GetBuilder() PolygonObjectBuilder
}

type PolygonObjectBuilder interface {
	Builder
	setReader(reader reader.PolygonObjectReader) error
	buildVertices() error
	buildPolygon() error
	buildCenter() error
}

func NewPolygonObjectDirector(factory PolygonObjectFactory, reader reader.PolygonObjectReader) *PolygonObjectDirector {
	return &PolygonObjectDirector{factory: factory, reader: reader}
}

func (d *PolygonObjectDirector) Construct() (object.Object, error) {
	builder := d.factory.GetBuilder()
	if err := builder.setReader(d.reader); err != nil {
		return nil, err
	}
	if err := builder.buildVertices(); err != nil {
		return nil, err
	}
	if err := builder.buildPolygon(); err != nil {
		return nil, err
	}
	if err := builder.buildCenter(); err != nil {
		return nil, err
	}
	return builder.getObject()
}
