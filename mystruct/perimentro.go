package mystruct

import "math"

type Retangulo struct {
	largura float64
	altura  float64
}

type Triangulo struct {
    base float64
    altura float64
}

type Circulo struct {
	raio float64
}

type Forma interface {
    Area() float64
}

func (t Triangulo) Area() float64 {
    return (t.base * t.altura) * 0.5
}

func (r Retangulo) Area() float64 {
	return r.altura * r.largura
}

func (c Circulo) Area() float64 {
	return math.Pi * c.raio * c.raio 
}