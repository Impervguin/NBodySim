package reader

type PolygonObjectReader interface {
	ReadPolygonObject() (*PolygonObject, error)
}
