package builder

import (
	"NBodySim/internal/object"
	"NBodySim/internal/reader"
	"path"
	"testing"
)

const PathToModels = "../../models/test/"

func TestClassicCube(t *testing.T) {
	filename := "cube.obj"
	file := path.Join(PathToModels, filename)
	reader, err := reader.NewObjReader(file)
	if err != nil {
		t.Error(err)
		return
	}

	factory := &ClassicPolygonFactory{}
	builder := factory.GetBuilder()
	builder.setReader(reader)

	err = builder.buildVertices()
	if err != nil {
		t.Error(err)
		return
	}
	err = builder.buildPolygon()
	if err != nil {
		t.Error(err)
		return
	}
	err = builder.buildCenter()
	if err != nil {
		t.Error(err)
		return
	}
	obj, err := builder.getObject()
	if err != nil {
		t.Error(err)
		return
	}

	polygonObject, ok := obj.(*object.PolygonObject)
	if !ok {
		t.Error("Expected PolygonObject, got", obj)
		return
	}
	if len(polygonObject.GetPolygons()) != 12 {
		t.Error("Expected 12 polygons, got", len(polygonObject.GetPolygons()))
	}
	if len(polygonObject.GetVertices()) != 8 {
		t.Error("Expected 8 vertices, got", len(polygonObject.GetVertices()))
	}
}

func TestClassicNonConvex(t *testing.T) {
	filename := "nonconvex.obj"
	file := path.Join(PathToModels, filename)
	reader, err := reader.NewObjReader(file)
	if err != nil {
		t.Error(err)
		return
	}

	factory := &ClassicPolygonFactory{}
	builder := factory.GetBuilder()
	builder.setReader(reader)

	err = builder.buildVertices()
	if err != nil {
		t.Error(err)
		return
	}
	err = builder.buildPolygon()
	if err == nil {
		t.Error("Expected non-convex polygon to fail")
	}
}

func TestClassicConvex(t *testing.T) {
	filename := "convex.obj"
	file := path.Join(PathToModels, filename)
	reader, err := reader.NewObjReader(file)
	if err != nil {
		t.Error(err)
		return
	}

	factory := &ClassicPolygonFactory{}
	builder := factory.GetBuilder()
	builder.setReader(reader)

	err = builder.buildVertices()
	if err != nil {
		t.Error(err)
		return
	}
	err = builder.buildPolygon()
	if err != nil {
		t.Error(err)
		return
	}
	err = builder.buildCenter()
	if err != nil {
		t.Error(err)
		return
	}
	obj, err := builder.getObject()
	if err != nil {
		t.Error(err)
		return
	}
	polygonObject, ok := obj.(*object.PolygonObject)
	if !ok {
		t.Error("Expected PolygonObject, got", obj)
		return
	}
	if len(polygonObject.GetPolygons()) != 3 {
		t.Error("Expected 3 polygon, got", len(polygonObject.GetPolygons()))
	}
}
