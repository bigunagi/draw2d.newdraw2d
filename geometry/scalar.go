package geometry

import (
	"math"
)

// float64
type Scalar float64

const (
	NearlyZero Scalar = 1e-12
)

func Scalars(floats ...Scalar) []Scalar{
	return floats
}

func (s Scalar) Abs() Scalar {
	return Scalar(math.Abs(float64(s)))
}

/*
// float32
type Scalar float32

const (
	Epsilon Scalar = 1e-12
)

func Scalars(floats ...Scalar) []Scalar{
	return floats
}

func (s Scalar) Abs() Scalar {
	return Scalar(math.Abs(float64(s)))
}
*/
/*
// int32
type Scalar int32

const (
	NearlyZero Scalar = 1e-12
)

func Scalars(floats ...Scalar) []Scalar{
	return floats
}

func (s Scalar) Abs() Scalar {
	return Scalar(math.Abs(float64(s)))
}
*/