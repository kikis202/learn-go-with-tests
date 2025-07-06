package main

import "strings"

type RomanNumerals struct {
	Value  uint16
	Symbol string
}

var allRomanNumerals = []RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(val uint16) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for val >= numeral.Value {
			result.WriteString(numeral.Symbol)
			val -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(val string) uint16 {
	var result uint16

	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(val, numeral.Symbol) {
			result += numeral.Value
			val = strings.TrimPrefix(val, numeral.Symbol)
		}
	}

	return result
}
