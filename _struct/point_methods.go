package _struct

import "math"

func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (p *Point) Scale(factor float64) *Point {
	p.X *= factor
	p.Y *= factor
	return p
}

func (p *Point3) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
}

func (p *Polar) Abs() float64 {
	x := p.R * math.Sin(p.Ceta) * math.Cos(p.Delta)
	y := p.R * math.Sin(p.Ceta) * math.Sin(p.Delta)
	z := p.R * math.Cos(p.Ceta)

	return math.Sqrt(x*x + y*y + z*z)
}
