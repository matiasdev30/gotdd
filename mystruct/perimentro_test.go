package mystruct

import "testing"

func TestPerimentro(t *testing.T) {
	tetAreas := []struct{
		forma Forma
		esperado float64
	}{
		{Retangulo{largura: 12, altura: 6}, 72.0},
		{Circulo{raio: 10}, 314.1592653589793},
	}

	for _, tt := range tetAreas{
		resultado := tt.forma.Area()

		if resultado != tt.esperado{
			t.Errorf("resultado %.2f, esperado %.2f", resultado, tt.esperado)
		}
	}
}
