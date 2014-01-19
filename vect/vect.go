package vect

import (
	"math"
)

// Origin is a zero value for V.
var Origin = V{0, 0}

// FClamp returns val clamped between min and max.
func FClamp(val, min, max float64) float64 {
	if val < min {
		return min
	} else if val > max {
		return max
	}
	return val
}

// A V represents vector of 2 coordinates.
type V struct{ X, Y float64 }

// New returns a new vector.
func New(x, y float64) V {
	return V{x, y}
}

// Equal returns true if v equals a.
func (v V) Equal(a V) bool {
	return v == a
}

// Equal returns true if a equals b.
func Equal(a, b V) bool {
	return a == b
}

// Add returns the sum a+b.
func (a V) Add(b V) V {
	return Add(a, b)
}

// Add returns the sum of two vectors.
func Add(a, b V) V {
	return V{a.X + b.X, a.Y + b.Y}
}

// Sub returns the difference a-b.
func (a V) Sub(b V) V {
	return Sub(a, b)
}

// Sub returns the difference of two vectors.
func Sub(a, b V) V {
	return V{a.X - b.X, a.Y - b.Y}
}

// Negate returns v, negated.
func (v V) Negate() V {
	return Negate(v)
}

// Negate returns negated vector.
func Negate(v V) V {
	return V{-v.X, -v.Y}
}

// Mult returns the product a*s.
func (a V) Mult(s float64) V {
	return Mult(a, s)
}

// Mult returns the product of vector v and scalar s.
func Mult(v V, s float64) V {
	return V{v.X * s, v.Y * s}
}

// Dot returns dot product for v with a.
func (v V) Dot(a V) float64 {
	return Dot(v, a)
}

// Dot returns dot product between two vectors.
func Dot(a, b V) float64 {
	return (a.X * b.X) + (a.Y * b.Y)
}

// Cross returns cross product between two vectors.
func Cross(a, b V) float64 {
	return (a.X * b.Y) - (a.Y * b.X)
}

// CrossVF returns cross product between a vector and a float.
// Equivalent to
//   V{s*v.Y, -s*v.Y}
func CrossVF(v V, s float64) V {
	return V{s * v.Y, -s * v.X}
}

// CrossFV returns cross product between a float and a vector.
// Not the same as CrossVF, equivalent to
//   V{-s*v.Y, s*v.X}
func CrossFV(s float64, v V) V {
	return V{-s * v.Y, s * v.X}
}

// Perp returns v rotated by 90°.
func Perp(v V) V {
	return V{-v.Y, v.X}
}

// Rperp returns v rotated by -90°.
func Rperp(v V) V {
	return V{v.Y, -v.X}
}

// Project returns the vector projection of v onto a.
func (v V) Project(a V) V {
	return Project(v, a)
}

// Project returns the vector projection of a onto b.
func Project(a, b V) V {
	return Mult(b, Dot(a, b)/Dot(b, b))
}

// ForAngle returns the unit length vector for the angle in radians.
func ForAngle(angle float64) V {
	return V{math.Cos(angle), math.Sin(angle)}
}

// ToAngle returns the angular direction v is pointing in (in radians).
func (v V) ToAngle() float64 {
	return math.Atan2(v.X, v.Y)
}

// ToAngle is the same as V.ToAngle()
func ToAngle(v V) float64 {
	return v.ToAngle()
}

// FromComplex returns a new vector with X and Y set respectively
// to real and imaginary part of a complex.
func FromComplex(c complex128) V {
	return V{real(c), imag(c)}
}

// ToComplex returns a new complex number X + Yi.
func (v V) ToComplex() complex128 {
	return complex(v.X, v.Y)
}

// ToComplex is the same as V.ToComplex()
func ToComplex(v V) complex128 {
	return v.ToComplex()
}

// Rotate returns v rotated by a.
func (v V) Rotate(a V) V {
	return Rotate(v, a)
}

// Rotate returns a rotated by b.
func Rotate(a, b V) V {
	return V{a.X*b.X - a.Y*b.Y, a.X*b.Y + a.Y*b.X}
}

// Unrotate is a reverse of V.Rotate()
func (v V) Unrotate(a V) V {
	return Unrotate(v, a)
}

// Unrotate is a reverse of Rotate()
func Unrotate(a, b V) V {
	return V{a.X*b.X + a.Y*b.Y, a.Y*b.X - a.X*b.Y}
}

// LengthSquared returns square root of vector's length.
func (v V) LengthSquared() float64 {
	// length of a vector is a distance to origin
	return DistanceSquared(v, V{})
}

// LengthSquared returns the squared length of the vector.
func LengthSquared(v V) float64 {
	return DistanceSquared(v, V{})
}

// Length returns vector length.
func (v V) Length() float64 {
	//length of a vector is a distance to origin
	return Distance(v, V{})
}

// Length returns the length of the vector.
func Length(v V) float64 {
	return Distance(v, V{})
}

// Lerp returns linear interpolation between a and b.
func Lerp(a, b V, t float64) V {
	return V{
		a.X + (a.X-b.X)*t,
		a.Y + (a.Y-b.Y)*t,
	}
}

// Normalize returns v normalized to the length of 1.
func (v V) Normalize() V {
	return Normalize(v)
}

// Normalize returns a vector of v normalized to the length of 1.
func Normalize(v V) V {
	f := 1.0 / (Length(v) + math.SmallestNonzeroFloat64)
	return Mult(v, f)
}

// DistanceSquared returns the distance from a for v, squared.
func (v V) DistanceSquared(a V) float64 {
	return DistanceSquared(v, a)
}

// DistSquared returns the squared distance between vectors.
func DistanceSquared(a, b V) float64 {
	return (a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y)
}

// Distance returns the distance from a for v.
func (v V) Distance(a V) float64 {
	return math.Sqrt(DistanceSquared(v, a))
}

// Distance returns the distance between two vectors.
func Distance(a, b V) float64 {
	return math.Sqrt(DistanceSquared(a, b))
}

// Clamp returns v clamped to length ln.
func (v V) Clamp(ln float64) V {
	return Clamp(v, ln)
}

// Clamp returns v clamped to length ln.
func Clamp(v V, ln float64) V {
	if Dot(v, v) > ln*ln {
		return Mult(Normalize(v), ln)
	}
	return v
}

// Min returns a vector with minimal X and Y from two vectors.
func Min(a, b V) V {
	return V{math.Min(a.X, b.X), math.Min(a.Y, b.Y)}
}

// Max returns a vector with maximal X and Y from two vectors.
func Max(a, b V) V {
	return V{math.Max(a.X, b.X), math.Max(a.Y, b.Y)}
}
