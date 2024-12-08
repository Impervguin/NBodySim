package nbody

type NBodyVisitor interface {
	VisitBody(body Body)
	VisitNBody(body *NBody)
}
