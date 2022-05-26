package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	if len(input) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	cleanExpression := getCleanExpression(input)
	oneOperand, twoOperand := getFirstAndSecondOperands(cleanExpression)

	if _, err = checkOperandsForValid(oneOperand, twoOperand); err != nil {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	oneNumber, err := strconv.Atoi(oneOperand)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	twoNumber, err := strconv.Atoi(twoOperand)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	sumNumber := oneNumber + twoNumber
	output = strconv.Itoa(sumNumber)

	return output, nil
}

// checkOperandsForValid Checking operands for validity.
func checkOperandsForValid(operandOne, operandTwo string) (string, error) {
	if len(operandOne) == 0 || len(operandTwo) == 0 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	for i := range operandOne {
		if i > 0 {
			if operandOne[i] == 45 || operandOne[i] == 43 {
				return "", fmt.Errorf("%w", errorNotTwoOperands)
			}
		}
	}

	return "", nil
}

// getFirstAndSecondOperands Get 2 operands according to the task condition.
func getFirstAndSecondOperands(expression string) (oneOperand string, twoOperand string) {
	for i := range expression {
		if expression[i] == 45 || expression[i] == 43 {
			oneOperand = expression[:i]
			// If the expression is with 3 operands, then in twoOperand it will write the last operand with a sign.
			twoOperand = expression[i:]
		}
	}

	return oneOperand, twoOperand
}

// getCleanExpression Get a clean expression that includes letters, numbers, + and -. Excluding all other characters.
func getCleanExpression(input string) string {
	runes := []rune(input)
	expression := ""
	for i := range runes {
		if !unicode.IsNumber(runes[i]) && !unicode.IsLetter(runes[i]) && runes[i] != 45 && runes[i] != 43 {
			continue
		}
		expression += string(runes[i])
	}

	return expression
}
