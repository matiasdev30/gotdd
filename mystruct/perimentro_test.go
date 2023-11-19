package mystruct

import "testing"

func TestPerimentro(t *testing.T) {
	resultado := Perimetro(10.0, 10.0)
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("resultado -> '%2.f' esperado -> '%2.f'", resultado, esperado)
	}
}