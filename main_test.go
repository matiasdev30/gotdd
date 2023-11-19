package main

import "testing"

func TestOla(t *testing.T)  {

	verificarTest := func(resultado string, esperado string){
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s' esperado '%s'", resultado, esperado)
		} 
	}

	t.Run("Diz olá quando uma string for passada", func(t *testing.T){
		resultado := Ola("Mani", "")
		esperado := "Olá, Mani"

		verificarTest(resultado, esperado)
	})

	t.Run("Diz olá sem passar nenhum parametro", func(t *testing.T){
		resultado := Ola("", "")
		esperado := "Olá, Mundo"

		verificarTest(resultado, esperado)
	})

	t.Run("em espanhol", func(t * testing.T){
		resultado := Ola("Eloide", "espanhol")
		esperado := "Holla, Eloide"

		verificarTest(resultado, esperado)
	})

	t.Run("em frances", func(t * testing.T){
		resultado := Ola("Lima", "françes")
		esperado := "Bonjour, Lima"

		verificarTest(resultado, esperado)
	})
}

