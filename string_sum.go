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

const (
	Empty      = ""
	PlusAscii  = 43
	MinusAscii = 45
)

func StringSum(input string) (output string, err error) {
	if len(input) == 0 {
		return Empty, fmt.Errorf("%w", errorEmptyInput)
	}

	cleanExpression := getCleanExpression(input)
	oneOperand, twoOperand := getFirstAndSecondOperands(cleanExpression)

	if _, err = checkOperandsForValid(oneOperand, twoOperand); err != nil {
		return Empty, fmt.Errorf("%w", errorNotTwoOperands)
	}

	oneNumber, err := strconv.Atoi(oneOperand)
	if err != nil {
		return Empty, fmt.Errorf("%w", err)
	}

	twoNumber, err := strconv.Atoi(twoOperand)
	if err != nil {
		return Empty, fmt.Errorf("%w", err)
	}

	sumNumber := oneNumber + twoNumber
	output = strconv.Itoa(sumNumber)

	return output, nil
}

// checkOperandsForValid Checking operands for validity.
func checkOperandsForValid(operandOne, operandTwo string) (string, error) {
	if len(operandOne) == 0 || len(operandTwo) == 0 {
		return Empty, fmt.Errorf("%w", errorNotTwoOperands)
	}

	for i, val := range operandOne {
		if i > 0 {
			if val == MinusAscii || val == PlusAscii {
				return Empty, fmt.Errorf("%w", errorNotTwoOperands)
			}
		}
	}
	return Empty, nil
}

// getFirstAndSecondOperands Get 2 operands according to the task condition.
func getFirstAndSecondOperands(expression string) (oneOperand string, twoOperand string) {
	for i, val := range expression {
		if val == PlusAscii || val == MinusAscii {
			oneOperand = expression[:i]
			// If the expression is with 3 operands, then in twoOperand it will write the last operand with a sign.
			twoOperand = expression[i:]
		}
	}
	return
}

// getCleanExpression Get a clean expression that includes letters, numbers, + and -. Excluding all other characters.
func getCleanExpression(input string) string {
	expression := Empty
	for _, val := range input {
		if unicode.IsNumber(val) || unicode.IsLetter(val) || val == PlusAscii || val == MinusAscii {
			expression += string(val)
		}
	}
	return expression
}
