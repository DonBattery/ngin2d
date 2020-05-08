package ngin2d

import "sort"

const MinimumDistance = 0.1

// Collision object created when an Elem (ElemA) has been moved by a Vector (OriginalVector)
// and its collision were checked against another Elem (ElemB) and they are colliding.
// Type is the type of the collision (needs to be reworked)
// Resolution is the reversed OriginalVector and its length is equal to the distance ElemA needs to travel in order to avoid collision with ElemB
type Collision struct {
	// Type       CollisionType
	Resolution *Vector
}

// func NewCollision(t CollisionType, v *Vector) *Collision {
func NewCollision(v *Vector) *Collision {
	return &Collision{
		// Type:       t,
		Resolution: v,
	}
}

// Collisions is a list of Collision
type Collisions struct {
	collisions []*Collision
}

// NewCollisions creates a new Collisions object and returnst its pointer
func NewCollisions() *Collisions {
	return &Collisions{}
}

// Add adds a Collision to the Collisions and sorts them by their resolution vector in descending order
func (cs *Collisions) Add(colls ...*Collision) {
	cs.collisions = append(cs.collisions, colls...)
	sort.Slice(cs.collisions, func(i, j int) bool {
		return cs.collisions[i].Resolution.Len() > cs.collisions[j].Resolution.Len()
	})
}

// Next takes the first Collision (with the longest resolution vector) from the Collisions
// and returns it along a boolean indicating if ther are more Collision in the list
// If Next called on an empty Collisions object, a nil pointer and false will be returned
func (cs *Collisions) Next() (*Collision, bool) {
	if cs.collisions != nil && len(cs.collisions) > 0 {
		coll := cs.collisions[0]
		cs.collisions = cs.collisions[1:]
		return coll, len(cs.collisions) > 0
	}
	return nil, false
}
