package calculate

import (
	"testing"

)

func TestCaculate(t *testing.T) {

	mensagemDeVerificacao := func (resultado float64, esperado float64) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%f' esperado '%f'", resultado, esperado)
		}
	}

	t.Run("Fazendo a adição de dois valores", func(t *testing.T) {
		resultado := Calculate(2, 4, "+")
		var esperado float64 = 6

		mensagemDeVerificacao(resultado, esperado)
	})

	t.Run("Fazendo a sutração de dois valores", func(t *testing.T){
		resultado := Calculate(7, 8, "-")
		var esperado float64 = -1

		mensagemDeVerificacao(resultado, esperado)
	})

	t.Run("Fazendo a divisão de dois valores", func(t *testing.T) {
		resultado := Calculate(7, 8, "/")
		var esperado float64 = 0.875

		mensagemDeVerificacao(resultado, esperado)
	})

	t.Run("Fazendo a multiplicação de dois valores", func(t *testing.T){
		resulatdo := Calculate(7, 8, "*")
		var esperado float64 = 56

		mensagemDeVerificacao(resulatdo, esperado)
	})

}
