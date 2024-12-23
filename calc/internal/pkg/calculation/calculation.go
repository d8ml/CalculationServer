package calculation

import (
	errors "github.com/d8ml/calculation_server/calc/internal/pkg"
	"math"
	"strconv"
)

func checkSym(ch uint8) bool {
	if ch == ' ' {
		return true
	}

	if ch >= '0' && ch <= '9' {
		return true
	}

	if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
		return true
	}

	if ch == '(' || ch == ')' {
		return true
	}

	return false
}

func isNum(ch uint8) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}

func Calculate(input string) (float64, error) {
	calcSeq := make([]float64, 0)
	opStack := make([]int, 0)
	curS := ""

	for i := 0; i < len(input); i++ {
		if !checkSym(input[i]) {
			return 0, errors.InvalidExpression
		}

		ch := input[i]
		switch ch {
		case '+':
			if len(curS) > 0 {
				num, _ := strconv.ParseFloat(curS, 64)
				calcSeq = append(calcSeq, num)
				curS = ""
			}
			for len(opStack) > 0 && opStack[len(opStack)-1] <= 4 {
				calcSeq = append(calcSeq, float64(2*1e9+opStack[len(opStack)-1]))
				opStack = opStack[:len(opStack)-1]
			}
			opStack = append(opStack, 1)
		case '-':
			if len(curS) > 0 {
				num, _ := strconv.ParseFloat(curS, 64)
				calcSeq = append(calcSeq, num)
				curS = ""
			}
			if i == len(input)-1 {
				return 0, errors.InvalidExpression
			}

			j := i - 1
			for j >= 0 && input[j] == ' ' {
				j--
			}
			if j >= 0 && (isNum(input[j]) || input[j] == ')') {
				for len(opStack) > 0 && opStack[len(opStack)-1] <= 4 {
					calcSeq = append(calcSeq, float64(2*1e9+opStack[len(opStack)-1]))
					opStack = opStack[:len(opStack)-1]
				}
				opStack = append(opStack, 2)
			} else {
				if isNum(input[i+1]) {
					curS += "-"
				} else {
					for len(opStack) > 0 && opStack[len(opStack)-1] <= 4 {
						calcSeq = append(calcSeq, float64(2*1e9+opStack[len(opStack)-1]))
						opStack = opStack[:len(opStack)-1]
					}
					opStack = append(opStack, 2)
				}
			}
		case '*':
			if len(curS) > 0 {
				num, _ := strconv.ParseFloat(curS, 64)
				calcSeq = append(calcSeq, num)
				curS = ""
			}
			for len(opStack) > 0 && (opStack[len(opStack)-1] == 3 || opStack[len(opStack)-1] == 4) {
				calcSeq = append(calcSeq, float64(2*1e9+opStack[len(opStack)-1]))
				opStack = opStack[:len(opStack)-1]
			}
			opStack = append(opStack, 3)
		case '/':
			if len(curS) > 0 {
				num, _ := strconv.ParseFloat(curS, 64)
				calcSeq = append(calcSeq, num)
				curS = ""
			}
			for len(opStack) > 0 && (opStack[len(opStack)-1] == 3 || opStack[len(opStack)-1] == 4) {
				calcSeq = append(calcSeq, float64(2*1e9+opStack[len(opStack)-1]))
				opStack = opStack[:len(opStack)-1]
			}
			opStack = append(opStack, 4)
		case '(':
			if len(curS) > 0 {
				num, _ := strconv.ParseFloat(curS, 64)
				calcSeq = append(calcSeq, num)
				curS = ""
			}
			opStack = append(opStack, 5)
		case ')':
			if len(curS) > 0 {
				num, _ := strconv.ParseFloat(curS, 64)
				calcSeq = append(calcSeq, num)
				curS = ""
			}
			for len(opStack) > 0 && opStack[len(opStack)-1] < 5 {
				calcSeq = append(calcSeq, float64(2*1e9+opStack[len(opStack)-1]))
				opStack = opStack[:len(opStack)-1]
			}
			if len(opStack) == 0 || len(opStack) > 0 && opStack[len(opStack)-1] != 5 {
				return 0, errors.InvalidExpression
			}
			opStack = opStack[:len(opStack)-1]
		case ' ':
			if len(curS) > 0 {
				num, _ := strconv.ParseFloat(curS, 64)
				calcSeq = append(calcSeq, num)
				curS = ""
			}
		default:
			curS += string(ch)
		}
	}

	if len(curS) > 0 {
		num, _ := strconv.ParseFloat(curS, 64)
		calcSeq = append(calcSeq, num)
		curS = ""
	}

	for len(opStack) > 0 {
		if opStack[len(opStack)-1] == 5 {
			return 0, errors.InvalidExpression
		}
		calcSeq = append(calcSeq, float64(2*1e9+opStack[len(opStack)-1]))
		opStack = opStack[:len(opStack)-1]
	}

	process := make([]float64, 200)
	processP := -1
	for i := 0; i < len(calcSeq); i++ {
		if calcSeq[i] <= 2*1e9 {
			processP++
			process[processP] = calcSeq[i]
		} else {
			if processP < 1 {
				return 0, errors.InvalidExpression
			} else {
				a := process[processP]
				processP--
				b := process[processP]

				switch calcSeq[i] {
				case 2*1e9 + 1:
					process[processP] = a + b
				case 2*1e9 + 2:
					process[processP] = b - a
				case 2*1e9 + 3:
					process[processP] = a * b
				case 2*1e9 + 4:
					if math.Abs(a) < 1e-9 {
						return 0, errors.InvalidExpression
					}
					process[processP] = b / a
				}
			}
		}
	}

	if processP != 0 {
		return 0, errors.InvalidExpression
	} else {
		return process[0], nil
	}
}
