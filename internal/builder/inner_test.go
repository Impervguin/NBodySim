package builder

import (
	"NBodySim/internal/object"
	"NBodySim/internal/reader"
	"path"
	"testing"
)

func TestInnerCube(t *testing.T) {
	filename := "cube.obj"
	file := path.Join(PathToModels, filename)
	reader, err := reader.NewObjReader(file)
	if err != nil {
		t.Error(err)
		return
	}

	factory := &InnerNormalBuilderFactory{}
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
	for _, p := range polygonObject.GetPolygons() {
		if !p.NormalIsInner() {
			t.Error("Expected inner normal for polygon")
		}
		if !p.GetNormal().NormalIsInner {
			t.Error("Expected inner normal for polygon's normal")
		}
	}
}

func TestInnerNonConvex(t *testing.T) {
	filename := "nonconvex.obj"
	file := path.Join(PathToModels, filename)
	reader, err := reader.NewObjReader(file)
	if err != nil {
		t.Error(err)
		return
	}

	factory := &InnerNormalBuilderFactory{}
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

func TestInnerConvex(t *testing.T) {
	filename := "convex.obj"
	file := path.Join(PathToModels, filename)
	reader, err := reader.NewObjReader(file)
	if err != nil {
		t.Error(err)
		return
	}

	// 	v -1 0 0
	// v -1 1 0
	// v 1 1 0
	// pol := object.NewPolygon(vector.NewVector3d(-1, 0, 0), vector.NewVector3d(-1, 1, 0), vector.NewVector3d(1, 0, 0), color.Black)
	// t.Error(pol.GetNormal().ToVector())

	factory := &InnerNormalBuilderFactory{}
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

func TestPartNormals(t *testing.T) {
	filename := "partnormals.obj"
	file := path.Join(PathToModels, filename)
	reader, err := reader.NewObjReader(file)
	if err != nil {
		t.Error(err)
		return
	}

	factory := &InnerNormalBuilderFactory{}
	builder := factory.GetBuilder()
	builder.setReader(reader)

	err = builder.buildVertices()
	if err != nil {
		t.Error(err)
		return
	}
	err = builder.buildPolygon()
	if err == nil {
		t.Error("expected error")
		return
	}
}
