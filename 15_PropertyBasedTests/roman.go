package roman

import "strings"

type RomanNumerals []RomanNumeral

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var allRomanNumerals = RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{Value: 10, Symbol: "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func (r RomanNumerals) ValueOf(symbols ...byte) uint16 {
	// allows us to look up an integer value for a given roman numeral

	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

func ConvertToRoman(arabic uint16) string {
	// A Builder is used to efficiently build a string using Write methods. It minimizes memory copying.

	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) uint16 {
	total := uint16(0)

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		// look ahead to the next symbol if we can
		if couldBeSubtractive(i, symbol, roman) {
			nextSymbol := roman[i+1]

			if value := allRomanNumerals.ValueOf(symbol, nextSymbol); value != 0 {
				total += value
				i++ // move past this character too for the next loop
			} else {
				total += allRomanNumerals.ValueOf(symbol)
			}
		} else {
			total += allRomanNumerals.ValueOf(symbol)
		}
	}
	return total
}

func couldBeSubtractive(index int, currentSymbol uint8, roman string) bool {
	isSubtractiveSymbol := currentSymbol == 'I' || currentSymbol == 'X' || currentSymbol == 'C'
	return index+1 < len(roman) && isSubtractiveSymbol
}
