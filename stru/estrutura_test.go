package stru

import (
	"math"
	"testing"
)

func TestEstrutura(t *testing.T) {

	areaTest := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Retangulo", shape: Retangulo{alt: 2, larg: 2}, want: 4},
		{name: "Circulo", shape: Circulo{raio: 4}, want: 2 * 2 * math.Pi},
	}

	for _, a := range areaTest {
		t.Run("Testando calculo da area", func(t *testing.T) {
			got := a.shape.Area()
			want := a.want
			if got != want {
				t.Errorf("Esperava %v e obteve %v (%T)", want, got, a.shape)
			}
		})
	}

}
