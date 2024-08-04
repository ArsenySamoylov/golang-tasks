package areacalc

const pi = 3.14159

type Shape interface {
	Area() float64
	Type() string
}

type Rectangle struct {
	w, h float64
	name string
}

func NewRectangle(w, h float64, name string) *Rectangle {
	return &Rectangle{w: w, h: h, name: name}
}

func (r Rectangle) Area() float64 {
	return r.w * r.h
}

func (r Rectangle) Type() string {
	return "rectangle"
}

type Circle struct {
	r    float64
	name string
}

func NewCircle(r float64, name string) *Circle {
	return &Circle{r: r, name: name}
}

func (c Circle) Area() float64 {
	return pi * c.r * c.r
}

func (c Circle) Type() string {
	return "circle"
}

func AreaCalculator(figures []Shape) (accum string, area float64) {
	for i, f := range figures {
		area += f.Area()
		accum += f.Type()

		if i < len(figures)-1 {
			accum += "-"
		}
	}

	return
}
