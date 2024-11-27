package normal

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/transform"
)

// start con be moved only in linear space, not in perspective space
type Normal struct {
	Start, End    vector.Vector3d
	NormalIsInner bool
}

func NewNormal(start, end vector.Vector3d) *Normal {
	return &Normal{Start: start, End: end}
}

func (pn *Normal) ToVector() vector.Vector3d {
	return *vector.SubtractVectors(&pn.End, &pn.Start)
}

func (pn *Normal) Transform(transform transform.TransformAction) {
	transform.ApplyToVector(&pn.Start)
	transform.ApplyToVector(&pn.End)
}

func (pn *Normal) Copy() *Normal {
	return &Normal{
		Start:         *pn.Start.Copy(),
		End:           *pn.End.Copy(),
		NormalIsInner: pn.NormalIsInner,
	}
}
