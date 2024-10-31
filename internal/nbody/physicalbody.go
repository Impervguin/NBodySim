package nbody

import "NBodySim/internal/mathutils/vector"

type PhysBody struct {
	Id       int64
	Position vector.Vector3d
	Velocity vector.Vector3d
	Mass     float64
}

func NewPhysBody(id int64, position, velocity vector.Vector3d, mass float64) *PhysBody {
	return &PhysBody{
		Id:       id,
		Position: position,
		Velocity: velocity,
		Mass:     mass,
	}
}

func FromBody(body Body) *PhysBody {
	return NewPhysBody(body.GetId(), body.GetPosition(), body.GetVelocity(), body.GetMass())
}

func (pb *PhysBody) GetId() int64 {
	return pb.Id
}

func (pb *PhysBody) Clone() *PhysBody {
	return NewPhysBody(pb.Id, pb.Position, pb.Velocity, pb.Mass)
}

func (pb *PhysBody) GetPosition() vector.Vector3d {
	return pb.Position
}

func (pb *PhysBody) GetVelocity() vector.Vector3d {
	return pb.Velocity
}

func (pb *PhysBody) GetMass() float64 {
	return pb.Mass
}

func (pb *PhysBody) Update(other Body) {
	pb.Position = other.GetPosition()
	pb.Velocity = other.GetVelocity()
	pb.Mass = other.GetMass()
}

func (pb *PhysBody) SetPosition(position vector.Vector3d) {
	pb.Position = position
}

func (pb *PhysBody) SetVelocity(velocity vector.Vector3d) {
	pb.Velocity = velocity
}

func (pb *PhysBody) SetMass(mass float64) {
	pb.Mass = mass
}
