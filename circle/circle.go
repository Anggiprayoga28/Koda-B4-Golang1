package circle

type Circle struct {
	Radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{
		Radius: radius,
	}
}

func (c *Circle) IsValid() bool {
	return c.Radius > 0
}
