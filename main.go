package main


const espanhol = "espanhol"
const frances = "françes"
const prefixoFrances = "Bonjour"
const prefixoOlaEmPt = "Olá"
const prefixoOlaEspanhol = "Holla"

func Ola(name string, idioma string) string {

	if name == ""{
		name = "Mundo"
	}

	return prefixoSaudacao(idioma) + ", " + name
}

func prefixoSaudacao(idioma string)(prefixo string){

	switch idioma {
	case "espanhol":
		prefixo = prefixoOlaEspanhol
	case "françes":
		prefixo = prefixoFrances
	default :
		prefixo = prefixoOlaEmPt
	}

	return 
}

