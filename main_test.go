package main

import "testing"

func TestOla(t *testing.T)  {
	resultado := Ola("Mani")
	esperado := "Olá, Mani"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}

