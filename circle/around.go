package circle

import "math"

func (c *Circle) CalculateCircumference() float64 {
	return 2 * math.Pi * c.Radius
}
