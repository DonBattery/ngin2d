package ngin2d

type GameObject interface {
	GetXY() (float64, float64)
	SetXY(float64, float64)
	GetMass() float64
	SetMass(float64)
	GetVelocity() *Vector
	SetVelocity(*Vector)
	GetAcceleration() *Vector
	SetAcceleration(*Vector)
	ApplyForce(*Force)
	Move()
	MoveBy(*Vector)

	GetRect() *Rect
	IsColliding(GameObject) bool
	WouldCollide(*Vector, GameObject) *Collision
	GetResolution(*Vector, GameObject) *Vector
}

// BasicObject almost implements all requirement of Elem except the collision methods
type BasicObject struct {
	mass         float64
	position     *Vector
	velocity     *Vector
	acceleration *Vector
}

func (bo *BasicObject) GetMass() float64 {
	return bo.mass
}

func (bo *BasicObject) SetMass(m float64) {
	bo.mass = m
}

func (bo *BasicObject) GetVelocity() *Vector {
	return bo.velocity
}

func (bo *BasicObject) SetVelocity(v *Vector) {
	bo.velocity = v
}

func (bo *BasicObject) GetAcceleration() *Vector {
	return bo.acceleration
}

func (bo *BasicObject) SetAcceleration(v *Vector) {
	bo.acceleration = v
}

func (bo *BasicObject) ApplyForce(f *Force) {
	bo.acceleration.Add(f.v)
}

func (bo *BasicObject) Move() {
	bo.velocity.Add(bo.acceleration)
	bo.position.Add(bo.velocity)
}

func (bo *BasicObject) MoveBy(v *Vector) {
	bo.position.Add(v)
}
