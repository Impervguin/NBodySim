package object

type ObjectWithId struct {
	id int64
}

func (b *ObjectWithId) GetId() int64 {
	return b.id
}

var globalId int64 = 0

func getNextId() int64 {
	globalId++
	return globalId
}
