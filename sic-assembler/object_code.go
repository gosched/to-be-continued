package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

// ObjectCodeContainer .
type ObjectCodeContainer struct {
	RecordMaxLen     int
	Header           *Header
	Text             *Text
	End              *End
	ObjectCodeRecord string
}

// Init .
func (ofd *ObjectCodeContainer) Init() {
	ofd.Header.LineBeginSymbol = "H"
	ofd.Text.LineBeginSymbol = "T"
	ofd.End.LineBeginSymbol = "E"
}

// GetHeaderData .
func (ofd *ObjectCodeContainer) GetHeaderData() (headerData string) {
	headerData += ofd.Header.LineBeginSymbol
	headerData += FillSuffixWithBlank(ofd.Header.ProgramName, ofd.RecordMaxLen)
	headerData += FillPrefixWithZero(ofd.Header.StartingAddress, ofd.RecordMaxLen)
	headerData += FillPrefixWithZero(ofd.Header.ProgramLength, ofd.RecordMaxLen)
	headerData += "\n"
	return headerData
}

// GetTextData .
func (ofd *ObjectCodeContainer) GetTextData(a *Assembler) (textData string) {

	list := list.New()

	for _, v := range a.IntermediateFileContainer {
		if v.Label+v.Location+v.ObjectCode+v.Opcode+v.Operand != "" && v.Opcode != "START" {
			if v.Opcode == "BYTE" {
				if strings.HasPrefix(v.Operand, "C") == true {
					tmp := strings.Split(v.Operand, "")
					tmp2 := strings.Split(v.ObjectCode, "")
					LocationCounter := a.ConvertHexStringToInt(v.Location) - 1
					for i := 0; i < len(tmp)-1; i++ {
						if i >= 2 {
							t := &StatementRecord{}
							LocationCounter++
							t.Location = fmt.Sprintf("%x", LocationCounter)
							t.Opcode = v.Opcode
							t.ObjectCode = tmp2[i*2-4] + tmp2[i*2-3]
							list.PushBack(t)
						}
					}
				}
				if strings.HasPrefix(v.Operand, "X") == true {
					tmp := strings.Split(v.ObjectCode, "")
					LocationCounter := a.ConvertHexStringToInt(v.Location) - 1
					for i := 0; i < len(tmp)-1; i++ {
						if i%2 == 0 {
							t := &StatementRecord{}
							LocationCounter++
							t.Location = fmt.Sprintf("%x", LocationCounter)
							t.Opcode = v.Opcode
							t.Operand = tmp[i]
							t.ObjectCode = tmp[i] + tmp[i+1]
							list.PushBack(t)
						}
					}
				}
			} else {
				t := &StatementRecord{}
				t.Label = v.Label
				t.Location = v.Location
				t.ObjectCode = v.ObjectCode
				t.Opcode = v.Opcode
				t.Operand = v.Operand
				list.PushBack(t)
			}
		}
	}

	rowSlice := []*StatementRecord{}
	byteLimit := 0

	for e := list.Front(); e != nil; e = e.Next() {

		if e.Value.(*StatementRecord).Opcode == "END" {
			break
		}

		// fmt.Println(e.Value.(*StatementRecord))

		rowSlice = append(rowSlice, e.Value.(*StatementRecord))
		tmp := ""
		for _, v := range rowSlice {
			tmp += v.ObjectCode
		}

		byteLimit = (strings.Count(tmp, "") - 1) / 2
		if byteLimit > 30 {
			backStatementRecord := rowSlice[len(rowSlice)-1]
			tmp := ""
			for i := 0; i < len(rowSlice)-1; i++ {
				tmp += rowSlice[i].ObjectCode
			}

			realLen := FillPrefixWithZero(strconv.FormatInt(int64((strings.Count(tmp, "")-1)/2), 16), 2)
			textData += ofd.Text.LineBeginSymbol + FillPrefixWithZero(rowSlice[0].Location, 6) + realLen + tmp + "\n"

			byteLimit = 0
			rowSlice = []*StatementRecord{}
			rowSlice = append(rowSlice, backStatementRecord)
		}

		if byteLimit == 30 {
			tmp := ""
			for _, v := range rowSlice {
				tmp += v.ObjectCode
			}

			realLen := FillPrefixWithZero(strconv.FormatInt(int64((strings.Count(tmp, "")-1)/2), 16), 2)
			textData += ofd.Text.LineBeginSymbol + FillPrefixWithZero(rowSlice[0].Location, 6) + realLen + tmp + "\n"

			byteLimit = 0
			rowSlice = []*StatementRecord{}
		}

		if byteLimit < 30 {
			if e.Value.(*StatementRecord).Opcode == "RESW" || e.Value.(*StatementRecord).Opcode == "RESB" {

				if len(rowSlice) == 1 {
					byteLimit = 0
					rowSlice = []*StatementRecord{}
					continue
				}

				tmp := ""
				for _, v := range rowSlice {
					tmp += v.ObjectCode
				}

				realLen := FillPrefixWithZero(strconv.FormatInt(int64((strings.Count(tmp, "")-1)/2), 16), 2)
				textData += ofd.Text.LineBeginSymbol + FillPrefixWithZero(rowSlice[0].Location, 6) + realLen + tmp + "\n"

				byteLimit = 0
				rowSlice = []*StatementRecord{}
				continue
			}
			if e.Next() == nil {

				tmp := ""
				for _, v := range rowSlice {
					tmp += v.ObjectCode
				}
				realLen := FillPrefixWithZero(strconv.FormatInt(int64((strings.Count(tmp, "")-1)/2), 16), 2)
				textData += ofd.Text.LineBeginSymbol + FillPrefixWithZero(rowSlice[0].Location, 6) + realLen + tmp + "\n"
				byteLimit = 0
				break
			}

		}
	}
	if len(rowSlice) != 0 {
		tmp := ""
		for _, v := range rowSlice {
			tmp += v.ObjectCode
		}
		realLen := FillPrefixWithZero(strconv.FormatInt(int64((strings.Count(tmp, "")-1)/2), 16), 2)
		textData += ofd.Text.LineBeginSymbol + FillPrefixWithZero(rowSlice[0].Location, 6) + realLen + tmp + "\n"

		rowSlice = []*StatementRecord{}
	}

	textData = strings.ToUpper(textData)
	/*
		rowSlice := []*StatementRecord{}
		byteCounter := 0

		for e := list.Front(); e != nil; e = e.Next() {

			if e.Value.(*StatementRecord).Opcode == "END" {
				break
			}

			rowSlice = append(rowSlice, e.Value.(*StatementRecord))

			if byteCounter < 27 {
				if e.Value.(*StatementRecord).Opcode == "RESW" || e.Value.(*StatementRecord).Opcode == "RESB" {

					if len(rowSlice) == 1 {
						byteCounter = 0
						rowSlice = []*StatementRecord{}
						continue
					}

					fmt.Println("!")
					for _, v := range rowSlice {
						fmt.Println(v)
					}

					tmp := ""
					for _, v := range rowSlice {
						tmp += v.ObjectCode
					}
					realLen := FillPrefixWithZero(strconv.FormatInt(int64((strings.Count(tmp, "")-1)/2), 16), 2)
					textData += "T" + FillPrefixWithZero(rowSlice[0].Location, 6) + realLen + tmp + "\n"

					byteCounter = 0
					rowSlice = []*StatementRecord{}
					continue
				}
				if e.Next() == nil {

					fmt.Println("!")
					for _, v := range rowSlice {
						fmt.Println(v)
					}

					tmp := ""
					for _, v := range rowSlice {
						tmp += v.ObjectCode
					}
					realLen := FillPrefixWithZero(strconv.FormatInt(int64((strings.Count(tmp, "")-1)/2), 16), 2)
					textData += "T" + FillPrefixWithZero(rowSlice[0].Location, 6) + realLen + tmp + "\n"
					byteCounter = 0
					break
				}
			}

			if byteCounter == 27 {

				fmt.Println("!")
				for _, v := range rowSlice {
					fmt.Println(v)
				}

				tmp := ""
				for _, v := range rowSlice {
					tmp += v.ObjectCode
				}
				realLen := FillPrefixWithZero(strconv.FormatInt(int64((strings.Count(tmp, "")-1)/2), 16), 2)
				textData += "T" + FillPrefixWithZero(rowSlice[0].Location, 6) + realLen + tmp + "\n"

				byteCounter = 0
				rowSlice = []*StatementRecord{}
			}

			if e.Value.(*StatementRecord).Opcode == "BYTE" {
				if strings.HasPrefix(e.Value.(*StatementRecord).Operand, "C") == true {
					byteCounter++
				}
				if strings.HasPrefix(e.Value.(*StatementRecord).Operand, "X") == true {

					byteCounter++
				}
			} else {
				if e.Value.(*StatementRecord).Opcode != "RESW" && e.Value.(*StatementRecord).Opcode != "RESB" {
					byteCounter += 3
				}
			}
		}

		if len(rowSlice) != 0 {

			fmt.Println("!")
			for _, v := range rowSlice {
				fmt.Println(v)
			}

			tmp := ""
			for _, v := range rowSlice {
				tmp += v.ObjectCode
			}
			realLen := FillPrefixWithZero(strconv.FormatInt(int64((strings.Count(tmp, "")-1)/2), 16), 2)
			textData += "T" + FillPrefixWithZero(rowSlice[0].Location, 6) + realLen + tmp + "\n"

			rowSlice = []*StatementRecord{}
		}

		textData = strings.ToUpper(textData)
	*/

	/*
		list := list.New()
		rowSlice := []*StatementRecord{}

		for _, v := range IntermediateFileContainer {
			if v.Label+v.Location+v.ObjectCode+v.Opcode+v.Operand != "" && v.Opcode != "START" {
				t := &StatementRecord{}
				t.Label = v.Label
				t.Location = v.Location
				t.ObjectCode = v.ObjectCode
				t.Opcode = v.Opcode
				t.Operand = v.Operand
				list.PushBack(t)
			}
		}

		length, realLen := 0, ""

		for e := list.Front(); e != nil; e = e.Next() {
			length++

			if e.Value.(*StatementRecord).Opcode == "END" {
				break
			}

			rowSlice = append(rowSlice, e.Value.(*StatementRecord))

			if length < 10 {
				if e.Value.(*StatementRecord).Opcode == "RESW" || e.Value.(*StatementRecord).Opcode == "RESB" {
					length--
					realLen = FillPrefixWithZero(strconv.FormatInt(int64(length*3), 16), 2)

					if len(rowSlice) == 1 {
						length = 0
						rowSlice = []*StatementRecord{}
						continue
					}

					textData += ("T" + FillPrefixWithZero(rowSlice[0].Location, 6) + realLen)
					for _, v := range rowSlice {
						textData += v.ObjectCode
					}
					textData += "\n"

					length = 0
					rowSlice = []*StatementRecord{}
					continue
				}

				if e.Next() == nil {
					realLen = FillPrefixWithZero(strconv.FormatInt(int64(length*3), 16), 2)
					textData += ("T" + FillPrefixWithZero(rowSlice[0].Location, 6) + realLen)
					for _, v := range rowSlice {
						textData += v.ObjectCode
					}
					textData += "\n"
					break
				}
			}

			if length == 10 {
				realLen = FillPrefixWithZero(strconv.FormatInt(int64(length*3), 16), 2)
				textData += ("T" + FillPrefixWithZero(rowSlice[0].Location, 6) + realLen)
				for _, v := range rowSlice {
					textData += v.ObjectCode
				}
				textData += "\n"

				length = 0
				rowSlice = []*StatementRecord{}
			}
		}

		if len(rowSlice) != 0 {
			realLen = FillPrefixWithZero(strconv.FormatInt(int64(length*3), 16), 2)
			textData += ("T" + FillPrefixWithZero(rowSlice[0].Location, 6) + realLen)
			for _, v := range rowSlice {
				textData += v.ObjectCode
			}
			textData += "\n"

			rowSlice = []*StatementRecord{}
		}

		textData = strings.ToUpper(textData)
	*/

	return textData
}

// GetEndData .
func (ofd *ObjectCodeContainer) GetEndData(s *SymbolTable) (endData string) {
	endData += ofd.End.LineBeginSymbol + FillPrefixWithZero(s.HashTable[ofd.End.FirstCanRunInstructionSymbol], ofd.RecordMaxLen) + "\n"
	return endData
}

// Header .
type Header struct {
	LineBeginSymbol string
	ProgramName     string
	StartingAddress string
	ProgramLength   string // hex, byte
}

// Text .
type Text struct {
	LineBeginSymbol string
	ObjectCode      string
}

// End .
type End struct {
	LineBeginSymbol              string
	FirstCanRunInstructionSymbol string
}

// FillPrefixWithZero .
func FillPrefixWithZero(sourceStr string, maxLength int) (newStr string) {
	for i := 1; i <= maxLength-len(sourceStr); i++ {
		newStr += "0"
	}
	newStr += sourceStr

	return newStr
}

// FillSuffixWithZero .
func FillSuffixWithZero(sourceStr string, maxLength int) (newStr string) {
	for i := 1; i <= maxLength-len(sourceStr); i++ {
		newStr += "0"
	}
	return newStr
}

// FillSuffixWithBlank .
func FillSuffixWithBlank(sourceStr string, maxLength int) (newStr string) {
	newStr += sourceStr
	for i := 1; i <= maxLength-len(sourceStr)%maxLength; i++ {
		newStr += " "
	}
	return newStr
}
