package main

import (
	"math"
	"strings"
)

// PositionalNotationConverter .
type PositionalNotationConverter struct {
	HexadecimalCharDecimalIntTable map[string]int
}

// Init .
func (PNC *PositionalNotationConverter) Init() {
	PNC.HexadecimalCharDecimalIntTable = make(map[string]int)
	PNC.HexadecimalCharDecimalIntTable["0"] = 0
	PNC.HexadecimalCharDecimalIntTable["1"] = 1
	PNC.HexadecimalCharDecimalIntTable["2"] = 2
	PNC.HexadecimalCharDecimalIntTable["3"] = 3
	PNC.HexadecimalCharDecimalIntTable["4"] = 4
	PNC.HexadecimalCharDecimalIntTable["5"] = 5
	PNC.HexadecimalCharDecimalIntTable["6"] = 6
	PNC.HexadecimalCharDecimalIntTable["7"] = 7
	PNC.HexadecimalCharDecimalIntTable["8"] = 8
	PNC.HexadecimalCharDecimalIntTable["9"] = 9
	PNC.HexadecimalCharDecimalIntTable["A"] = 10
	PNC.HexadecimalCharDecimalIntTable["B"] = 11
	PNC.HexadecimalCharDecimalIntTable["C"] = 12
	PNC.HexadecimalCharDecimalIntTable["D"] = 13
	PNC.HexadecimalCharDecimalIntTable["E"] = 14
	PNC.HexadecimalCharDecimalIntTable["F"] = 15
}

// ConvertHexStringToInt .
func (a *Assembler) ConvertHexStringToInt(s string) int {
	s = strings.ToUpper(s)
	maxExponential := len(s) - 1
	var result float64

	for i := range s {
		result += float64(a.PositionalNotationConverter.HexadecimalCharDecimalIntTable[string(s[i])]) * math.Pow(16, float64(maxExponential))
		maxExponential--
	}

	return int(result)
}
