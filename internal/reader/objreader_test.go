package reader

import (
	"path"
	"testing"
)

const PathToModels = "../../models/test/"

func TestCube(t *testing.T) {
	filename := "cube.obj"
	file := path.Join(PathToModels, filename)
	reader, err := NewObjReader(file)
	if err != nil {
		t.Error(err)
		return
	}
	po, err := reader.ReadPolygonObject()
	if err != nil {
		t.Error(err)
		return
	}
	if len(po.Polygons) != 12 {
		t.Error("Expected 12 polygons, got", len(po.Polygons))
	}
	if len(po.Edges) != 36 {
		t.Error("Expected 36 edges, got", len(po.Edges))
	}
	if len(po.Vertexes) != 8 {
		t.Error("Expected 8 vertexes, got", len(po.Vertexes))
	}
	if len(po.Normals) != 6 {
		t.Error("Expected 8 normals, got", len(po.Normals))
	}
}

func TestIndexOutOfRange(t *testing.T) {
	filename := "indexoutofrange.obj"
	file := path.Join(PathToModels, filename)
	reader, err := NewObjReader(file)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = reader.ReadPolygonObject()
	if err == nil {
		t.Error("Expected error, got nil ")
		return
	}
}

func TestInvalidFile(t *testing.T) {
	filename := "invalid.obj"
	file := path.Join(PathToModels, filename)
	reader, err := NewObjReader(file)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = reader.ReadPolygonObject()
	if err == nil {
		t.Error("Expected error, got nil ")
		return
	}
}

func TestNullIndex(t *testing.T) {
	filename := "nullindex.obj"
	file := path.Join(PathToModels, filename)
	reader, err := NewObjReader(file)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = reader.ReadPolygonObject()
	if err == nil {
		t.Error("Expected error, got nil ")
		return
	}
}

func TestNoNormals(t *testing.T) {
	filename := "nonormals.obj"
	file := path.Join(PathToModels, filename)
	reader, err := NewObjReader(file)
	if err != nil {
		t.Error(err)
		return
	}
	po, err := reader.ReadPolygonObject()
	if err != nil {
		t.Error(err)
		return
	}
	if len(po.Normals) != 0 {
		t.Error("Expected 0 normals, got", len(po.Normals))
	}
}
