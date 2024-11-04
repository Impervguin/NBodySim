package object

type ObjectPool struct {
	pool map[int64]Object
}

type ObjectPoolVisitor interface {
	ObjectVisitor
	VisitObjectPool(pool *ObjectPool)
}

func NewObjectPool() *ObjectPool {
	return &ObjectPool{pool: make(map[int64]Object)}
}

func (op *ObjectPool) GetObject(id int64) (Object, bool) {
	obj, found := op.pool[id]
	return obj, found
}

func (op *ObjectPool) GetObjects() []Object {
	objects := make([]Object, 0, len(op.pool))
	for _, obj := range op.pool {
		objects = append(objects, obj)
	}
	return objects
}

func (op *ObjectPool) PutObject(obj Object) {
	op.pool[obj.GetId()] = obj
}

func (op *ObjectPool) RemoveObject(id int64) {
	delete(op.pool, id)
}

func (op *ObjectPool) Accept(visitor ObjectVisitor) {
	if vis, ok := visitor.(ObjectPoolVisitor); ok {
		vis.VisitObjectPool(op)
	} else {
		for _, obj := range op.pool {
			obj.Accept(visitor)
		}
	}
}

func (op *ObjectPool) GetCount() int {
	return len(op.pool)
}

func (op *ObjectPool) Transform(t TransformAction) {
	for _, obj := range op.pool {
		obj.Transform(t)

	}
}

func (op *ObjectPool) Clone() *ObjectPool {
	newPool := make(map[int64]Object, len(op.pool))
	for id, obj := range op.pool {
		newPool[id] = obj.Clone()
	}
	return &ObjectPool{pool: newPool}
}
