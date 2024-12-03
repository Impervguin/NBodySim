package cutter

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"NBodySim/internal/transform"
	"testing"
)

func TestCutCubeSeen(t *testing.T) {
	cube, err := GetTestCube()
	if err != nil {
		t.Error(err)
		return
	}
	camera := object.NewCamera(*vector.NewVector3d(0, 0, 0), *vector.NewVector3d(0, 0, 1), *vector.NewVector3d(0, 1, 0), 1, 1, 1)
	cutter := NewSimpleCamCutter(camera)
	cube.Transform(transform.NewMoveAction(vector.NewVector3d(0, 0, 2)))
	cutted, _ := cube.Clone().(*object.PolygonObject)
	cutter.VisitPolygonObject(cutted)
	if len(cutted.GetPolygons()) != 12 {
		t.Errorf("Cube shouldn't change, got %v, wanted %v", cutted.GetPolygons(), cube.GetPolygons())
		return
	}
	cutPolygons, cubePolygons := cutted.GetPolygons(), cube.GetPolygons()
	for i := 0; i < 12; i++ {
		if !object.PolygonEqual(cubePolygons[i], cutPolygons[i]) {
			t.Errorf("Cut polygons should match cube polygons, got %v, wanted %v", cutted.GetPolygons(), cube.GetPolygons())
			return
		}
	}
}

func TestCutCubeNotSeen(t *testing.T) {
	cube, err := GetTestCube()
	if err != nil {
		t.Error(err)
		return
	}
	camera := object.NewCamera(*vector.NewVector3d(0, 0, 0), *vector.NewVector3d(0, 0, 1), *vector.NewVector3d(0, 1, 0), 1, 1, 1)
	cutter := NewSimpleCamCutter(camera)
	cube.Transform(transform.NewMoveAction(vector.NewVector3d(0, 0, -2)))
	cutted, _ := cube.Clone().(*object.PolygonObject)
	cutter.VisitPolygonObject(cutted)
	if len(cutted.GetPolygons()) != 0 {
		t.Errorf("Cube should change, got %v, wanted %v", cutted.GetPolygons(), cube.GetPolygons())
		return
	}
}

func TestCutCubePartSeen(t *testing.T) {
	cube, err := GetTestCube()
	if err != nil {
		t.Error(err)
		return
	}
	camera := object.NewCamera(*vector.NewVector3d(0, 0, 0), *vector.NewVector3d(0, 0, 1), *vector.NewVector3d(0, 1, 0), 1, 1, 1)
	cutter := NewSimpleCamCutter(camera)
	cube.Transform(transform.NewMoveAction(vector.NewVector3d(0, 0, 1)))
	cutted, _ := cube.Clone().(*object.PolygonObject)
	cutter.VisitPolygonObject(cutted)
	if len(cutted.GetPolygons()) == 12 {
		t.Errorf("Cube should change, got %v, wanted %v", cutted.GetPolygons(), cube.GetPolygons())
		return
	}
	for _, v := range cutted.GetVertices() {
		if v.Z < 1 {
			t.Errorf("Cut polygons shouldn't contain vertices behind the camera, got %v", cutted.GetVertices())
			return
		}
	}
}
