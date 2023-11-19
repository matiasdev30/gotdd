package calculate

const opSoma = "+"
const opSub = "-"
const opDiv = "/"
const opMul = "*"

func Calculate(n1 float64, n2 float64, op string) float64 {
	var result float64

	switch op {
	case opSoma:
		result = n1 + n2
	case opSub:
		result = n1 - n2
	case opDiv:
		result = n1 / n2
	case opMul:
		result = n1 * n2
	}

	return result
}