package formas

import (
	//"formas"
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida %f é diferente da área esperada %f",
				areaRecebida,
				areaEsperada,
			)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{float64(100000)}
		areaEsperada := float64(math.Pi * math.Pow(circ.Raio, 2))
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida %f é diferente da área esperada %f",
				areaRecebida,
				areaEsperada,
			)
		}
	})
}
