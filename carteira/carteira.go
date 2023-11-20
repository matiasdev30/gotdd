package carteira

import (
	"errors"
	"fmt"
)

type BitCoin int

type Carteira struct {
	saldo BitCoin
}

func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (c *Carteira) Depositar(qtd BitCoin) {
	c.saldo += qtd
}

func (c Carteira) Saldo() BitCoin {
	return c.saldo
}

func (c * Carteira) Retirar(qtd BitCoin) error {

	if c.Saldo() < qtd{
		return errors.New("Saldo insuficeinte")
	}

	c.saldo -= qtd
	return nil
}