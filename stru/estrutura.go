package stru

import "math"

//Shape é o formato
type Shape interface {
	Area() float64
}

//Retangulo é o tipo que modela um retângulo
type Retangulo struct {
	larg float64
	alt  float64
}

//Circulo é o tipo que modela um círculo
type Circulo struct {
	raio float64
}

//Perimetro calcula o perímetro de um retângulo
func (r Retangulo) Perimetro() float64 {
	return (r.alt + r.larg) * 2
}

//Area calcula a area de um retângulo
func (r Retangulo) Area() float64 {
	return r.alt * r.larg
}

//Area calcula a area de um retângulo
func (c Circulo) Area() float64 {
	return c.raio * c.raio * math.Pi
}
