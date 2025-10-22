package circle

import "math"

func (c *Circle) CalculateArea() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
