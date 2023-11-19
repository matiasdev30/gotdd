package main


const espanhol = "espanhol"
const frances = "françes"
const prefixoFrances = "Bonjour"
const prefixoOlaEmPt = "Olá"
const prefixoOlaEspanhol = "Holla"

func Ola(name string, idioma string) string {
	if name == "" {
		name = "Mundo"
	}

	if idioma == espanhol {
		return prefixoOlaEspanhol + ", " + name 
	}

	if(idioma == frances){
		return prefixoFrances + ", " + name
	}

	return prefixoOlaEmPt + ", " + name
}

