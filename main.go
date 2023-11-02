package main

import (
	"fmt"
	"strconv"
	"unicode"
)

type TokenType struct {
	ordinal int
	id      string
	word    *string
}

func NewTokenType(ordinal int, id string, word *string) TokenType {
	return TokenType{
		ordinal: ordinal,
		id:      id,
		word:    word,
	}
}

var NumberToken TokenType
var ConstantToken TokenType
var AddToken TokenType
var SubToken TokenType
var MulToken TokenType
var DivToken TokenType

var TokenTypeValues []TokenType

func initTypes() {
	NumberToken = NewTokenType(0, "num", nil)

	ConstantToken = NewTokenType(1, "const", nil)

	addWord := "+"
	addPointer := &addWord
	AddToken = NewTokenType(2, "add", addPointer)

	subWord := "-"
	subPointer := &subWord
	SubToken = NewTokenType(3, "sub", subPointer)

	mulWord := "*"
	mulPointer := &mulWord
	MulToken = NewTokenType(4, "mul", mulPointer)

	divWord := "/"
	divPointer := &divWord
	DivToken = NewTokenType(5, "div", divPointer)

	TokenTypeValues = []TokenType{
		NumberToken, ConstantToken,
		AddToken, SubToken, MulToken, DivToken,
	}
}

type Token struct {
	tokenType TokenType
	value     any
}

func (t Token) ToString() string {
	return fmt.Sprintf("Token{tokenType: %s, value: %v}", t.tokenType.id, t.value)
}

func lexString(str string) []Token {
	var tokens []Token
	runes := []rune(str)
	position := 0
	for position < len(runes) {
		switch {
		case unicode.IsDigit(runes[position]):
			numString := ""
			for position < len(runes) && (unicode.IsDigit(runes[position]) || runes[position] == '.') {
				numString += string(runes[position])
				position++
			}
			position--
			number, err := strconv.ParseFloat(numString, 64)
			if err != nil {
				fmt.Printf("Error parsing '%s' to a number\n%s\n", numString, err)
			} else {
				tokens = append(tokens, Token{
					tokenType: NumberToken,
					value:     number,
				})
			}
			break
		case unicode.IsLetter(runes[position]):
			constName := ""
			for position < len(runes) && (unicode.IsLetter(runes[position])) {
				constName += string(runes[position])
				position++
			}
			position--
			tokens = append(tokens, Token{
				tokenType: ConstantToken,
				value:     constName,
			})
			break
		default:
			for _, tokenType := range TokenTypeValues {
				if tokenType.word != nil {
					if string(runes[position]) == *tokenType.word {
						tokens = append(tokens, Token{
							tokenType: tokenType,
							value:     *tokenType.word,
						})
					}
				}
			}
			break
		}
		position++
	}
	return tokens
}

func main() {
	initTypes()

	for {
		fmt.Print("> ")

		var expr string
		_, err := fmt.Scanln(&expr)
		if err != nil {
			fmt.Print("Error: ", err)
			break
		}

		tokens := lexString(expr)

		fmt.Print("[")
		for i, value := range tokens {
			fmt.Print(value.ToString())
			if i < len(tokens)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println("]")

		node := parseTokens(tokens)

		fmt.Println(node.ToString())

		evaluated := evalTree(node, map[string]any{})

		fmt.Println(evaluated)
	}
}
