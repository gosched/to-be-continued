package main

import (
	"fmt"
	"os"
)

// Assembler .
type Assembler struct {
	OpcodeTable *OpcodeTable

	SymbolTable *SymbolTable

	SourceCodeData *SourceCodeData

	IntermediateFileContainer []*StatementRecord

	ObjectCodeContainer *ObjectCodeContainer

	PositionalNotationConverter *PositionalNotationConverter

	WordEqualHowManyBytes int
}

// NewAssembler .
func NewAssembler() *Assembler {
	a := &Assembler{

		PositionalNotationConverter: &PositionalNotationConverter{},

		OpcodeTable: &OpcodeTable{},
		SymbolTable: &SymbolTable{},

		IntermediateFileContainer: []*StatementRecord{},

		ObjectCodeContainer: &ObjectCodeContainer{
			Header: &Header{},
			Text:   &Text{},
			End:    &End{},
		},

		SourceCodeData: &SourceCodeData{},
	}
	return a
}

// InitAssembler .
func (a *Assembler) InitAssembler(OpcodeTableDefineDataPath, SICORSIXE string) (err error) {

	a.PositionalNotationConverter.Init()

	a.OpcodeTable.Init()
	a.SymbolTable.Init()
	a.ObjectCodeContainer.Init()
	a.SourceCodeData.Init()

	if SICORSIXE == "sic" || SICORSIXE == "SIC" {
		a.WordEqualHowManyBytes = 3
		a.ObjectCodeContainer.RecordMaxLen = 6
	}

	err = a.OpcodeTable.Load(OpcodeTableDefineDataPath)

	return err
}

// LoadSourceCode .
func (a *Assembler) LoadSourceCode(sourceCodePath string) error {

	err := a.SourceCodeData.Load(sourceCodePath)
	if err != nil {
		return err
	}

	return nil
}

// CreateRecordsData .
func (a *Assembler) CreateRecordsData() error {

	err := a.Part1()
	if err != nil {
		Logger.Fatalln("SetSymbolTableAndPartObjectData", err)
		return err
	}

	err = a.Part2()
	if err != nil {
		fmt.Println(err)
	}

	a.Part3()

	return nil
}

// ExportRecords .
func (a *Assembler) ExportRecords(ouputPath string) error {

	file, err := os.OpenFile(ouputPath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	file.WriteString(a.ObjectCodeContainer.ObjectCodeRecord)

	return nil
}

// ConflictWithMnemonic .
func (a *Assembler) ConflictWithMnemonic(symbol string) bool {
	if _, exist := a.OpcodeTable.BKDRHashTable[a.OpcodeTable.BKDRHash(symbol)]; exist {
		return true
	}
	switch symbol {
	case "START":
		return true
	case "END":
		return true
	case "WORD":
		return true
	case "RESW":
		return true
	case "RESB":
		return true
	case "BYTE":
		return true
	}
	return false
}
