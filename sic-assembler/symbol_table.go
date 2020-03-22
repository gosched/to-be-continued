package main

import (
	"regexp"
	"strconv"
	"strings"
)

// BKER's Hash Algorithm

// SymbolTable .
type SymbolTable struct {
	HashTable map[string]string // Symbol, HexadecimalLocation
}

// Init .
func (st *SymbolTable) Init() {
	st.HashTable = make(map[string]string)
}

// UpdateLocationCounter .
func (st *SymbolTable) UpdateLocationCounter(locationCounter *int, Type, OperandContent string) error {

	switch Type {
	case "OPCODE":
		*locationCounter = *locationCounter + 3
	case "WORD":
		*locationCounter = *locationCounter + 3
	case "RESW":
		tmp, err := strconv.Atoi(OperandContent)
		if err != nil {
			return err
		}
		*locationCounter = *locationCounter + 3*tmp
	case "RESB":
		tmp, err := strconv.Atoi(OperandContent)
		if err != nil {
			return err
		}
		*locationCounter = *locationCounter + tmp
		//Logger.Println("locationCounter", fmt.Sprintf("%x", *locationCounter))
	case "BYTEWithC":
		*locationCounter = *locationCounter + len(OperandContent)
		/*
			if strings.HasPrefix(Operand, "C") == true {
				fmt.Println("Operand", Operand, len(Operand))
				fmt.Println()
				*locationCounter = *locationCounter + len(Operand) - 3
				//tmp := strings.Split(Operand, "")
				//*locationCounter = *locationCounter + len(tmp) - 3
			}
		*/
		/*
			if strings.HasPrefix(OperandContent, "X") == true {
				tmp := strings.Split(OperandContent, "")
				*locationCounter = *locationCounter + (len(tmp)-3)/2
			} else {
				*locationCounter = *locationCounter + len(OperandContent)
			}
		*/
	case "BYTEWithX":
		*locationCounter = *locationCounter + (len(OperandContent)-3)/2
	}

	return nil
}

// InsertLabelAddressPair .
func (st *SymbolTable) InsertLabelAddressPair() {

}

// FindByteLen . error
func (st *SymbolTable) FindByteLen(Operand string) int {
	if strings.HasPrefix(Operand, "C") == true {
		tmp := strings.Split(Operand, "")
		return len(tmp) - 3
	}

	if strings.HasPrefix(Operand, "X") == true {
		return 1
	}

	return 0
}

// FindWithKey .
func (st *SymbolTable) FindWithKey(key string) (string, bool) {
	if value, ok := st.HashTable[key]; ok == true {
		return value, true
	}
	return "", false
}

// IsDuplicatedDefine .
func (st *SymbolTable) IsDuplicatedDefine(symbol string) bool {
	if _, exist := st.HashTable[symbol]; exist {
		return true
	}
	return false
}

// NotDuplicatedDefine .
func (st *SymbolTable) NotDuplicatedDefine(symbol string) bool {
	if _, exist := st.HashTable[symbol]; exist {
		return false
	}
	return true
}

// LabelFormatInValid .
func (st *SymbolTable) LabelFormatInValid(symbol string) bool {
	if m, _ := regexp.MatchString("^[A-Za-z0-9]+$", symbol); m {
		return false
	}
	return true
}

// LabelFormatValid .
func (st *SymbolTable) LabelFormatValid(symbol string) bool {
	if m, _ := regexp.MatchString("^[A-Za-z0-9]+$", symbol); m {
		return true
	}
	return false
}
