// Copyright 2011 The draw2d Authors. All rights reserved.
// created: 27/05/2011 by Laurent Le Goff

package geometry

import (
	"math"
)

// 2d vector
type Vector struct {
	X, Y float64
}

var (
	// Zero Vector
	ZeroVector = Vector{0, 0}
)

// Check if v1 is Equals to v2
func (v1 Vector) Equals(v2 Vector) bool {
	return v1.X == v2.X && v1.Y == v2.Y
}

// Check if v1 is Nearly Equals to v2 (against an epsilon defined in NearlyZero constant)
func (v1 Vector) NearlyEquals(v2 Vector) bool {
	return math.Abs(v1.X-v2.X) < NearlyZero && math.Abs(v1.Y-v2.Y) < NearlyZero
}

// Add two vectors
func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{v1.X + v2.Y, v1.Y + v2.Y}
}

// Substract two vector (v1 - v2)
func (v1 Vector) Sub(v2 Vector) Vector {
	return Vector{v1.X - v2.X, v1.Y - v2.Y}
}

// Opposite vector
func (v Vector) Opposite() Vector {
	return Vector{-v.X, -v.Y}
}

// float64 multiplication
func (v Vector) Mult(s float64) Vector {
	return Vector{v.X * s, v.Y * s}
}

// Vector dot product
func (v1 Vector) Dot(v2 Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

// Length Square of v
func (v Vector) Length() float64 {
	return math.Hypot(v.X, v.Y)
}

// Length Square of v
func (v Vector) LengthSquare() float64 {
	return v.Dot(v)
}

// Perpendicular vector. (90 degree rotation)
func (v Vector) Normal() Vector {
	return Vector{-v.Y, v.X}
}

// Returns the vector projection of v1 onto v2.
func (v1 Vector) Projection(v2 Vector) Vector {
	return v2.Mult(v1.Dot(v2) / v2.Dot(v2))
}

// Rotate v1 by v2. Scaling will occur if v1 is not a unit vector.
func (v1 Vector) Rotate(v2 Vector) Vector {
	return Vector{v1.X*v2.X - v1.Y*v2.Y, v1.Y*v2.X + v1.X*v2.Y}
}

// Inverse of Rotate().
func (v1 Vector) UnRotate(v2 Vector) Vector {
	return Vector{v1.X*v2.X + v1.Y*v2.Y, v1.Y*v2.X - v1.X*v2.Y}
}

// Return the vector between v1 and v2
func (v1 Vector) Center(v2 Vector) Vector {
	return Vector{(v1.X + v2.X) / 2, (v1.Y + v2.Y) / 2}
}

// Return the vector between v1 and v2 and interpolated by t (0, 1) 
// v*(1-t)+v2*t
func (v1 Vector) Lerp(v2 Vector, t float64) Vector {
	return v1.Mult(1 - t).Add(v2.Mult(t))
}

// Return the unit Vector of v
func (v Vector) Normalize() Vector {
	if v.X == 0 && v.Y == 0 {
		return ZeroVector
	}
	return v.Mult(1 / v.Length())
}

// Create a vector parallel to v that have d length
func (v Vector) SetLength(d float64) Vector {
	if v.X == 0 && v.Y == 0 {
		return ZeroVector
	}
	return v.Mult(d / v.Length())
}

// Create a vector parallel to v that have d length vector length is greater than d
func (v Vector) Clamp(d float64) Vector {
	lensq := v.LengthSquare()
	if lensq <= d*d {
		return v

	}
	return v.Mult(d / math.Sqrt(lensq))
}

// Returns the distance between v1 and v2.
func (v1 Vector) Distance(v2 Vector) float64 {
	return v1.Sub(v2).Length()
}

func (v1 Vector) DistanceSquare(v2 Vector) float64 {
	return v1.Sub(v2).LengthSquare()
}

// Returns the unit length vector for the given angle (in radians).
func AngleToVector(a float64) Vector {
	return Vector{math.Cos(a), math.Sin(a)}
}

// Returns the angular direction v is pointing in (in radians).
func (v Vector) Angle() float64 {
	return math.Atan2(v.Y, v.X)
}
