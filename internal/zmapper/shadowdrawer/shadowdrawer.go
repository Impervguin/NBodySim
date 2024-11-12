package shadowdrawer

import (
	"NBodySim/internal/zmapper/objectdrawer"
	"NBodySim/internal/zmapper/shadowmapper"
)

type ShadowObjectDrawer interface {
	objectdrawer.ObjectDrawer
	VisitShadowMapper(mapper *shadowmapper.ShadowMapper)
}

type ShadowObjectDrawerFabric interface {
	CreateShadowObjectDrawer() ShadowObjectDrawer
}
