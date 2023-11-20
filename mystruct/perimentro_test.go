package mystruct

import "testing"


//Table drive test
func TestPerimentro(t *testing.T) {
	tetAreas := []struct{
		forma Forma
		esperado float64
	}{
		{forma: Retangulo{largura: 12, altura: 6}, esperado: 72.0},
		{forma: Circulo{raio: 10}, esperado: 314.1592653589793},
		{forma: Triangulo{base: 12, altura: 6}, esperado: 36.0},
	}

	for _, tt := range tetAreas{
		resultado := tt.forma.Area()

		if resultado != tt.esperado{
			t.Errorf("%#v resultado %.2f, esperado %.2f", tt.forma, resultado, tt.esperado)
		}
	}
}
