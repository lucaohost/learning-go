package interfacetest

type Retangle struct {
	Width  float64
	Height float64
}

func (r Retangle) Area() float64 {
	return r.Height * r.Width
}
