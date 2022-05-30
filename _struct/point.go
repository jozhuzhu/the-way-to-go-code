package _struct

import "math"

type Point struct {
	X float64
	Y float64
}

type Point3 struct {
	X, Y, Z float64
}

type Polar struct {
	R     float64
	Delta float64
	Ceta  float64
}

func Abs(p *Point) float64 {
	return math.Sqrt(math.Pow(p.X, 2) + math.Pow(p.Y, 2))
}

func Scale2D(p *Point, factor float64) *Point {
	p.X = factor * p.X
	p.Y = factor * p.Y

	return p
}

func Scale3D(p *Polar, factor float64) *Polar {
	x := p.R * math.Sin(p.Ceta) * math.Cos(p.Delta)
	y := p.R * math.Sin(p.Ceta) * math.Sin(p.Delta)
	z := p.R * math.Cos(p.Ceta)

	x *= factor
	y *= factor
	z *= factor

	p.R = math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2) + math.Pow(z, 2))

	return p
}
