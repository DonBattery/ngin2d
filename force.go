package ngin2d

type Force struct {
	v *Vector
}

func NewForce(v *Vector) *Force {
	return &Force{
		v: v,
	}
}
