package balancear

var closeExpected = ""
var pS, pE, sS, sE, cS, cE int
var runes = []rune(closeExpected)

//Balance testa o fechamento correto de {[()]}
func Balance(str string) bool {

	for i := 0; i < len(str); i++ {
		runes = []rune(closeExpected)
		if str[i] == '{' {
			closeExpected += "}"
		}
		if str[i] == '[' {
			closeExpected += "]"
		}
		if str[i] == '(' {
			closeExpected += ")"
		}

		if str[i] == '}' || str[i] == ']' || str[i] == ')' {
			if fechamento(str[i]) {

			} else {
				return false
			}
		}
	}

	if len(str) == 0 || len(closeExpected) > 0 {
		return false
	}
	return true
}

func fechamento(caracter byte) bool {
	if len(closeExpected) > 0 && closeExpected[len(closeExpected)-1] == caracter {
		closeExpected = string(runes[:len(closeExpected)-1])
	} else {
		return false
	}
	return true
}
