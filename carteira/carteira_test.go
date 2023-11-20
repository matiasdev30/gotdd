package carteira

import (
	"testing"
)

func TestCarteira(t *testing.T) {

	confirmarSaldo := func(t *testing.T, carteira Carteira, esperado BitCoin) {
		resultado := carteira.Saldo()

		if resultado != esperado {
			t.Errorf("resulatdo %s, esperado %s", resultado, esperado)
		}
	}

	t.Run("Fazendo deposito", func(t *testing.T) {
		carteira := Carteira{}

		carteira.Depositar(10)

		confirmarSaldo(t, carteira, BitCoin(10))
	})

	t.Run("Fazendo saque", func(t *testing.T) {
		saldoInicial := BitCoin(20)
		carteira := Carteira{saldo: saldoInicial}
		erro := carteira.Retirar(BitCoin(30))

		confirmarSaldo(t, carteira, BitCoin(10))

		if erro == nil {
			t.Error("Esperava um erro mais nenhum aconteceu")
		}else {
			t.Error(erro.Error())
		}
	})
}
