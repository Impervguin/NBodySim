package shadowapproximator

import (
	"NBodySim/internal/object"
	"NBodySim/internal/transform"
	"NBodySim/internal/zmapper/approximator"
	"NBodySim/internal/zmapper/approximator/colorist"
	"NBodySim/internal/zmapper/shadowmapper"
)

type ShadowDiscreteApproximator interface {
	ApproximatePolygon(pol *object.Polygon, ch chan<- approximator.DiscreteFlatPoint) error
	VisitShadowMapper(sh *shadowmapper.ShadowMapper)
	ToShadowTransform(reverse transform.TransformAction)
}

type ShadowDiscreteApproximatorFabric interface {
	CreateShadowDiscreteApproximator() ShadowDiscreteApproximator
	GetColorist() colorist.Colorist
}
