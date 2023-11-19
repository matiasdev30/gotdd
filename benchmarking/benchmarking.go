package benchmarking

const qtdRepeticoes = 4

func Repetir(v string) string {
	str := ""

	for c := 0; c < qtdRepeticoes; c ++ {
		str += "a"
	}

	return str
}