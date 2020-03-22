package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var errFlag bool

// go doc regexp/syntax

// Part1 .
func (a *Assembler) Part1() error {

	LocationCounter, StartLocation := 0, 0
	findSTART, findEND := false, false
	for lineNumber, statementInLine := range a.SourceCodeData.Statements {
		realLineNumber := lineNumber + 1

		if statementInLine == "" {
			t := &StatementRecord{
				LineNumber:           strconv.Itoa(realLineNumber),
				IsPureCommentORBlank: true,
			}
			a.IntermediateFileContainer = append(a.IntermediateFileContainer, t)
		} else {
			//沒有找到.的話, len(slice) == 1, slice[0] = 原始字串
			statementSplitWithCommentSymbol := strings.Split(strings.TrimSpace(statementInLine), ".")
			statementWithoutComment := strings.TrimSpace(statementSplitWithCommentSymbol[0])

			if statementWithoutComment == "" {
				// pure comment
				t := &StatementRecord{
					LineNumber:           strconv.Itoa(realLineNumber),
					IsPureCommentORBlank: true,
				}
				a.IntermediateFileContainer = append(a.IntermediateFileContainer, t)
			} else {
				// without comment statment
				t := &StatementRecord{
					Location:   strconv.FormatInt(int64(LocationCounter), 16),
					LineNumber: strconv.Itoa(realLineNumber),
				}

				wordsInLine := strings.Fields(statementWithoutComment)
				wordsInLineLength := len(wordsInLine)

				if strings.Count(statementWithoutComment, ",") > 1 {
					errFlag = true
					Logger.Printf("line %-3s syntax/format invalid\n", strconv.Itoa(realLineNumber))
				}
				if strings.Count(statementWithoutComment, ",") == 1 {
					if !strings.HasSuffix(statementWithoutComment, "X") {
						errFlag = true
						Logger.Printf("line %-3s syntax/format invalid\n", strconv.Itoa(realLineNumber))
					} else {
						partOperandSlice := strings.Split(statementWithoutComment, ",")
						if partOperandSlice[0] == "" {
							errFlag = true
							Logger.Printf("line %-3s syntax/format invalid\n", strconv.Itoa(realLineNumber))
						}

						if wordsInLineLength == 2 {
							// STCH BUFFER,X
							if _, exist := a.OpcodeTable.BKDRHashTable[a.OpcodeTable.BKDRHash(wordsInLine[0])]; !exist {
								errFlag = true
								Logger.Printf("line %-3s mnemonic/opcode invalid\n", strconv.Itoa(realLineNumber))
							} else {
								t.Opcode = wordsInLine[0]
								partOperand := strings.Trim(wordsInLine[1], ",X")
								if conflict := a.ConflictWithMnemonic(partOperand); conflict {
									errFlag = true
									Logger.Printf("line %-3s namespace conflict with mnemonic\n", strconv.Itoa(realLineNumber))
								} else {
									t.Operand = partOperand + ",X"
									err := a.SymbolTable.UpdateLocationCounter(&LocationCounter, "OPCODE", "")
									if err != nil {
										return err
									}
									// Logger.Println("index type 0", t.Label, t.Opcode, t.Operand)
								}
							}
						}
						if wordsInLineLength >= 3 {
							//partOperandSlice := strings.Split(statementWithoutComment, ",")
							if strings.TrimSpace(strings.Join(partOperandSlice[1:], "")) == "X" {
								if wordsInLineLength == 3 {
									if strings.HasSuffix(statementWithoutComment, ", X") {
										// STCH BUFFER, X
										if _, exist := a.OpcodeTable.BKDRHashTable[a.OpcodeTable.BKDRHash(wordsInLine[0])]; !exist {
											errFlag = true
											Logger.Printf("line %-3s mnemonic/opcode invalid\n", strconv.Itoa(realLineNumber))
										} else {
											t.Opcode = wordsInLine[0]
											partOperand := strings.Trim(wordsInLine[1], ",")
											if conflict := a.ConflictWithMnemonic(partOperand); conflict {
												errFlag = true
												Logger.Printf("line %-3s namespace conflict with mnemonic\n", strconv.Itoa(realLineNumber))
											} else {
												t.Operand = partOperand + ",X"
												err := a.SymbolTable.UpdateLocationCounter(&LocationCounter, "OPCODE", "")
												if err != nil {
													return err
												}
												// Logger.Println("index type 1", t.Label, t.Opcode, t.Operand)
											}
										}
									} else if strings.HasSuffix(statementWithoutComment, " ,X") {
										// STCH BUFFER ,X
										if _, exist := a.OpcodeTable.BKDRHashTable[a.OpcodeTable.BKDRHash(wordsInLine[0])]; !exist {
											errFlag = true
											Logger.Printf("line %-3s mnemonic/opcode invalid\n", strconv.Itoa(realLineNumber))
										} else {
											t.Opcode = wordsInLine[0]
											if conflict := a.ConflictWithMnemonic(wordsInLine[1]); conflict {
												errFlag = true
												Logger.Printf("line %-3s namespace conflict with mnemonic\n", strconv.Itoa(realLineNumber))
											} else {
												t.Operand = strings.TrimSpace(wordsInLine[1]) + ",X"
												err := a.SymbolTable.UpdateLocationCounter(&LocationCounter, "OPCODE", "")
												if err != nil {
													return err
												}
												// Logger.Println("index type 2", t.Label, t.Opcode, t.Operand)
											}
										}
									} else {
										// Apple STCH BUFFER,X
										labelIsValid := true
										if a.SymbolTable.LabelFormatInValid(wordsInLine[0]) {
											errFlag = true
											Logger.Printf("line %-3s symbol/label format invalid\n", strconv.Itoa(realLineNumber))
											labelIsValid = false
										}
										if a.SymbolTable.IsDuplicatedDefine(wordsInLine[0]) {
											errFlag = true
											Logger.Printf("line %-3s duplicated symbol/label\n", strconv.Itoa(realLineNumber))
											labelIsValid = false
										}
										if conflict := a.ConflictWithMnemonic(wordsInLine[0]); conflict {
											errFlag = true
											Logger.Printf("line %-3s namespace conflict with mnemonic\n", strconv.Itoa(realLineNumber))
											labelIsValid = false
										}
										if labelIsValid == true {
											t.Label = wordsInLine[0]
											a.SymbolTable.HashTable[t.Label] = fmt.Sprintf("%x", LocationCounter)
											if _, exist := a.OpcodeTable.BKDRHashTable[a.OpcodeTable.BKDRHash(wordsInLine[1])]; !exist {
												errFlag = true
												Logger.Printf("line %-3s mnemonic/opcode invalid\n", strconv.Itoa(realLineNumber))
											} else {
												t.Opcode = wordsInLine[1]
												partOperandSlice := strings.Split(wordsInLine[2], ",")
												if conflict := a.ConflictWithMnemonic(partOperandSlice[0]); conflict {
													errFlag = true
													Logger.Printf("line %-3s namespace conflict with mnemonic\n", strconv.Itoa(realLineNumber))
												} else {
													t.Operand = wordsInLine[2]
													err := a.SymbolTable.UpdateLocationCounter(&LocationCounter, "OPCODE", "")
													if err != nil {
														return err
													}
													// Logger.Println("index type 3", t.Label, t.Opcode, t.Operand)
												}
											}
										}
									}
								}
								if wordsInLineLength == 4 {
									if strings.Contains(statementWithoutComment, " , X") {
										// STCH	BUFFER , X
										if _, exist := a.OpcodeTable.BKDRHashTable[a.OpcodeTable.BKDRHash(wordsInLine[0])]; !exist {
											errFlag = true
											Logger.Printf("line %-3s mnemonic/opcode invalid\n", strconv.Itoa(realLineNumber))
										} else {
											t.Opcode = wordsInLine[0]
											if conflict := a.ConflictWithMnemonic(wordsInLine[1]); conflict {
												errFlag = true
												Logger.Printf("line %-3s namespace conflict with mnemonic\n", strconv.Itoa(realLineNumber))
											} else {
												t.Operand = wordsInLine[1] + ",X"
												// Logger.Println("index type 4", t.Label, t.Opcode, t.Operand)
											}
										}
									}
									// Apple STCH	BUFFER, X
									if strings.HasSuffix(statementWithoutComment, ", X") {
										labelIsValid := true
										if a.SymbolTable.LabelFormatInValid(wordsInLine[0]) {
											errFlag = true
											Logger.Printf("line %-3s symbol/label format invalid\n", strconv.Itoa(realLineNumber))
											labelIsValid = false
										}
										if a.SymbolTable.IsDuplicatedDefine(wordsInLine[0]) {
											errFlag = true
											Logger.Printf("line %-3s duplicated symbol/label\n", strconv.Itoa(realLineNumber))
											labelIsValid = false
										}
										if conflict := a.ConflictWithMnemonic(wordsInLine[0]); conflict {
											errFlag = true
											Logger.Printf("line %-3s namespace conflict with mnemonic\n", strconv.Itoa(realLineNumber))
											labelIsValid = false
										}
										if labelIsValid == true {
											t.Label = wordsInLine[0]
											a.SymbolTable.HashTable[t.Label] = fmt.Sprintf("%x", LocationCounter)
											if _, exist := a.OpcodeTable.BKDRHashTable[a.OpcodeTable.BKDRHash(wordsInLine[1])]; !exist {
												errFlag = true
												Logger.Printf("line %-3s mnemonic/opcode invalid\n", strconv.Itoa(realLineNumber))
											} else {
												t.Opcode = wordsInLine[1]
												partOperand := strings.Trim(wordsInLine[2], ",")
												if conflict := a.ConflictWithMnemonic(partOperand); conflict {
													errFlag = true
													Logger.Printf("line %-3s namespace conflict with mnemonic\n", strconv.Itoa(realLineNumber))
												} else {
													t.Operand = partOperand + ",X"
													err := a.SymbolTable.UpdateLocationCounter(&LocationCounter, "OPCODE", "")
													if err != nil {
														return err
													}
													// Logger.Println("index type 4", t.Label, t.Opcode, t.Operand)
												}
											}
										}
									}
									// Apple STCH	BUFFER ,X
									if strings.HasSuffix(statementWithoutComment, " ,X") {
										labelIsValid := true
										if a.SymbolTable.LabelFormatInValid(wordsInLine[0]) {
											errFlag = true
											Logger.Printf("line %-3s symbol/label format invalid\n", strconv.Itoa(realLineNumber))
											labelIsValid = false
										}
										if a.SymbolTable.IsDuplicatedDefine(wordsInLine[0]) {
											Logger.Printf("line %-3s duplicated symbol/label\n", strconv.Itoa(realLineNumber))
											errFlag = true
											labelIsValid = false
										}
										if conflict := a.ConflictWithMnemonic(wordsInLine[0]); conflict {
											errFlag = true
											Logger.Printf("line %-3s namespace conflict with mnemonic\n", strconv.Itoa(realLineNumber))
											labelIsValid = false
										}
										if labelIsValid == true {
											t.Label = wordsInLine[0]
											a.SymbolTable.HashTable[t.Label] = fmt.Sprintf("%x", LocationCounter)
											if _, exist := a.OpcodeTable.BKDRHashTable[a.OpcodeTable.BKDRHash(wordsInLine[1])]; !exist {
												errFlag = true
												Logger.Printf("line %-3s mnemonic/opcode invalid\n", strconv.Itoa(realLineNumber))
											} else {
												t.Opcode = wordsInLine[1]
												t.Operand = strings.TrimSpace(wordsInLine[2]) + ",X"
												err := a.SymbolTable.UpdateLocationCounter(&LocationCounter, "OPCODE", "")
												if err != nil {
													return err
												}
												// Logger.Println("index type 4", t.Label, t.Opcode, t.Operand)
											}
										}
									}

								}
								if wordsInLineLength == 5 {
									// Apple STCH	BUFFER , X
									labelIsValid := true
									if a.SymbolTable.LabelFormatInValid(wordsInLine[0]) {
										errFlag = true
										Logger.Printf("line %-3s symbol/label format invalid\n", strconv.Itoa(realLineNumber))
										labelIsValid = false
									}
									if a.SymbolTable.IsDuplicatedDefine(wordsInLine[0]) {
										errFlag = true
										Logger.Printf("line %-3s duplicated symbol/label\n", strconv.Itoa(realLineNumber))
										labelIsValid = false
									}
									if conflict := a.ConflictWithMnemonic(wordsInLine[0]); conflict {
										errFlag = true
										Logger.Printf("line %-3s namespace conflict with mnemonic\n", strconv.Itoa(realLineNumber))
										labelIsValid = false
									}
									if labelIsValid == true {
										t.Label = wordsInLine[0]
										a.SymbolTable.HashTable[t.Label] = fmt.Sprintf("%x", LocationCounter)
										if _, exist := a.OpcodeTable.BKDRHashTable[a.OpcodeTable.BKDRHash(wordsInLine[1])]; !exist {
											errFlag = true
											Logger.Printf("line %-3s mnemonic/opcode invalid\n", strconv.Itoa(realLineNumber))
										} else {
											t.Opcode = wordsInLine[1]
											if conflict := a.ConflictWithMnemonic(wordsInLine[2]); conflict {
												errFlag = true
												Logger.Printf("line %-3s namespace conflict with mnemonic\n", strconv.Itoa(realLineNumber))
											} else {
												t.Operand = wordsInLine[2] + ",X"
												err := a.SymbolTable.UpdateLocationCounter(&LocationCounter, "OPCODE", "")
												if err != nil {
													return err
												}
												// Logger.Println("index type 5", t.Label, t.Opcode, t.Operand)
											}
										}
									}
								}
							} else {
								errFlag = true
								Logger.Printf("line %-3s operand invalid\n", strconv.Itoa(realLineNumber))
							}
						}
					}
				}

				if !strings.HasSuffix(statementWithoutComment, "X") {
					switch wordsInLineLength {
					case 0:

					case 1:
						if wordsInLine[0] == "RSUB" {
							t.Opcode = wordsInLine[0]
							err := a.SymbolTable.UpdateLocationCounter(&LocationCounter, "OPCODE", "")
							if err != nil {
								return err
							}
						} else {
							errFlag = true
							Logger.Printf("line %-3s mnemonic/opcode invalid\n", strconv.Itoa(realLineNumber))
							errFlag = true
						}

					case 2:
						// maybe correct END + firstCanRunSymbol
						if wordsInLine[0] == "END" {
							findEND = true
							t.Opcode = wordsInLine[0]
							if conflict := a.ConflictWithMnemonic(wordsInLine[1]); conflict {
								errFlag = true
								Logger.Printf("line %-3s namespace conflict with mnemonic\n", strconv.Itoa(realLineNumber))
								errFlag = true
							}
							_, exist := a.SymbolTable.FindWithKey(wordsInLine[1])
							if exist == true {
								t.Operand = wordsInLine[1]
								a.IntermediateFileContainer = append(a.IntermediateFileContainer, t)
								a.ObjectCodeContainer.End.FirstCanRunInstructionSymbol = t.Operand
							} else {
								errFlag = true
								Logger.Printf("line %-3s first can run instruction location invalid\n", strconv.Itoa(realLineNumber))
								errFlag = true
							}
							goto end
						}

						if wordsInLine[0] == "RSUB" {
							errFlag = true
							Logger.Printf("line %-3s rsub format err\n", strconv.Itoa(realLineNumber))
						}

						// maybe correct label + rsub
						if wordsInLine[1] == "RSUB" {
							labelIsValid := true
							if a.SymbolTable.LabelFormatInValid(wordsInLine[0]) {
								errFlag = true
								Logger.Printf("line %-3s symbol/label format invalid\n", strconv.Itoa(realLineNumber))
								labelIsValid = false
							}
							if a.SymbolTable.IsDuplicatedDefine(wordsInLine[0]) {
								errFlag = true
								Logger.Printf("line %-3s duplicated symbol/label\n", strconv.Itoa(realLineNumber))
								labelIsValid = false
							}
							if conflict := a.ConflictWithMnemonic(wordsInLine[0]); conflict {
								errFlag = true
								Logger.Printf("line %-3s namespace conflict with mnemonic\n", strconv.Itoa(realLineNumber))
								labelIsValid = false
							}
							if labelIsValid == true {
								t.Label = wordsInLine[0]
								t.Opcode = wordsInLine[1]
								a.SymbolTable.HashTable[t.Label] = fmt.Sprintf("%x", LocationCounter)
								err := a.SymbolTable.UpdateLocationCounter(&LocationCounter, "OPCODE", "")
								if err != nil {
									return err
								}
							}
						} else {
							// maybe correct opcode + operand
							if _, exist := a.OpcodeTable.BKDRHashTable[a.OpcodeTable.BKDRHash(wordsInLine[0])]; !exist {
								errFlag = true
								Logger.Printf("line %-3s mnemonic/opcode invalid\n", strconv.Itoa(realLineNumber))
							} else {
								if conflict := a.ConflictWithMnemonic(wordsInLine[1]); conflict {
									errFlag = true
									Logger.Printf("line %-3s namespace conflict with mnemonic\n", strconv.Itoa(realLineNumber))
								} else {
									t.Opcode = wordsInLine[0]
									t.Operand = wordsInLine[1]
									err := a.SymbolTable.UpdateLocationCounter(&LocationCounter, "OPCODE", "")
									if err != nil {
										return err
									}
								}
							}
						}
					case 3:
						// label opcode operand
						labelIsValid := true
						if a.SymbolTable.LabelFormatInValid(wordsInLine[0]) {
							errFlag = true
							Logger.Printf("line %-3s symbol/label format invalid\n", strconv.Itoa(realLineNumber))
							labelIsValid = false
						}
						if a.SymbolTable.IsDuplicatedDefine(wordsInLine[0]) {
							errFlag = true
							Logger.Printf("line %-3s duplicated symbol/label\n", strconv.Itoa(realLineNumber))
							labelIsValid = false
						}
						if conflict := a.ConflictWithMnemonic(wordsInLine[0]); conflict {
							errFlag = true
							Logger.Printf("line %-3s namespace conflict with mnemonic\n", strconv.Itoa(realLineNumber))
							labelIsValid = false
						}
						if labelIsValid == true {
							t.Label = wordsInLine[0]
							a.SymbolTable.HashTable[t.Label] = fmt.Sprintf("%x", LocationCounter)
							if _, exist := a.OpcodeTable.BKDRHashTable[a.OpcodeTable.BKDRHash(wordsInLine[1])]; !exist {
								switch wordsInLine[1] {
								case "WORD":
									t.Opcode = wordsInLine[1]
									if m, _ := regexp.MatchString("^[0-9]+$", wordsInLine[2]); m {
										if conflict := a.ConflictWithMnemonic(wordsInLine[2]); conflict {
											errFlag = true
											Logger.Printf("line %-3s namespace conflict with mnemonic\n", strconv.Itoa(realLineNumber))
										} else {
											t.Operand = wordsInLine[2]
											err := a.SymbolTable.UpdateLocationCounter(&LocationCounter, "WORD", "")
											if err != nil {
												return err
											}
										}
									} else {
										errFlag = true
										Logger.Printf("line %-3s operand not base 10\n", strconv.Itoa(realLineNumber))
										errFlag = true
									}
								case "RESW":
									t.Opcode = wordsInLine[1]
									if m, _ := regexp.MatchString("^[0-9]+$", wordsInLine[2]); m {
										if conflict := a.ConflictWithMnemonic(wordsInLine[2]); conflict {
											errFlag = true
											Logger.Printf("line %-3s namespace conflict with mnemonic\n", strconv.Itoa(realLineNumber))
										} else {
											t.Operand = wordsInLine[2]
											err := a.SymbolTable.UpdateLocationCounter(&LocationCounter, "RESW", wordsInLine[2])
											if err != nil {
												return err
											}
										}
									} else {
										errFlag = true
										Logger.Printf("line %-3s operand not base 10\n", strconv.Itoa(realLineNumber))
										errFlag = true
									}
								case "RESB":
									t.Opcode = wordsInLine[1]
									if m, _ := regexp.MatchString("[^0-9]", wordsInLine[2]); m {
										errFlag = true
										Logger.Printf("line %-3s operand invalid\n", strconv.Itoa(realLineNumber))
									} else if conflict := a.ConflictWithMnemonic(wordsInLine[2]); conflict {
										errFlag = true
										Logger.Printf("line %-3s namespace conflict with mnemonic\n", strconv.Itoa(realLineNumber))
									} else {
										t.Operand = wordsInLine[2]
										err := a.SymbolTable.UpdateLocationCounter(&LocationCounter, "RESB", wordsInLine[2])
										if err != nil {
											return err
										}
									}

								case "BYTE":
									t.Opcode = wordsInLine[1]
									if strings.HasSuffix(statementWithoutComment, "'") {
										if strings.HasPrefix(wordsInLine[2], "C'") {
											operandContentPrefixIndex := strings.Index(statementInLine, "'") + 1
											operandContentSuffixIndex := strings.LastIndex(statementInLine, "'") - 1
											if operandContentSuffixIndex-operandContentPrefixIndex > -1 {
												operandInSlice := strings.Split(statementWithoutComment, "")
												operandContent := strings.Join(operandInSlice[operandContentPrefixIndex:operandContentSuffixIndex+1], "")
												t.Operand = "C'" + operandContent + "'"
												err := a.SymbolTable.UpdateLocationCounter(&LocationCounter, "BYTEWithC", operandContent)
												if err != nil {
													return err
												}
											} else {
												errFlag = true
												Logger.Printf("line %-3s operand length invalid\n", strconv.Itoa(realLineNumber))
											}
										} else if strings.HasPrefix(wordsInLine[2], "X'") {
											operandContentPrefixIndex := strings.Index(statementInLine, "'") + 1
											operandContentSuffixIndex := strings.LastIndex(statementInLine, "'") - 1
											if operandContentSuffixIndex-operandContentPrefixIndex >= 1 {
												operandInSlice := strings.Split(statementWithoutComment, "")
												operandContent := strings.Join(operandInSlice[operandContentPrefixIndex:operandContentSuffixIndex+1], "")
												ok := true
												if m, _ := regexp.MatchString("[^0-9A-Fa-f]", operandContent); m {
													errFlag = true
													ok = false
													Logger.Printf("line %-3s operand invalid\n", strconv.Itoa(realLineNumber))
												}
												if len(operandContent)%2 != 0 {
													errFlag = true
													ok = false
													Logger.Printf("line %-3s operand invalid\n", strconv.Itoa(realLineNumber))
												}
												if ok {
													t.Operand = wordsInLine[2]
													err := a.SymbolTable.UpdateLocationCounter(&LocationCounter, "BYTEWithX", wordsInLine[2])
													if err != nil {
														return err
													}
												}
											} else {
												Logger.Printf("line %-3s operand length invalid\n", strconv.Itoa(realLineNumber))
											}
										} else {
											errFlag = true
											Logger.Printf("line %-3s operand invalid\n", strconv.Itoa(realLineNumber))
										}
									} else {
										errFlag = true
										Logger.Printf("line %-3s operand invalid\n", strconv.Itoa(realLineNumber))
									}

								case "START":
									findSTART = true
									t.Location = wordsInLine[2]
									t.Opcode = wordsInLine[1]
									if m, _ := regexp.MatchString("[^0-9A-Fa-f]", wordsInLine[2]); m {
										errFlag = true
										Logger.Printf("line %-3s operand not base 16\n", strconv.Itoa(realLineNumber))
										errFlag = true
									} else {
										if conflict := a.ConflictWithMnemonic(wordsInLine[2]); conflict {
											errFlag = true
											Logger.Printf("line %-3s namespace conflict with mnemonic\n", strconv.Itoa(realLineNumber))
											errFlag = true
										}
										t.Operand = wordsInLine[2]
									}
									a.ObjectCodeContainer.Header.ProgramName = wordsInLine[0]
									a.ObjectCodeContainer.Header.StartingAddress = wordsInLine[2]
									StartLocation = a.ConvertHexStringToInt(wordsInLine[2])
									LocationCounter = a.ConvertHexStringToInt(wordsInLine[2])

								case "END":
									findEND = true
									t.Opcode = wordsInLine[1]
									_, exist := a.SymbolTable.FindWithKey(wordsInLine[2])
									if exist == true {
										if conflict := a.ConflictWithMnemonic(wordsInLine[2]); conflict {
											errFlag = true
											Logger.Printf("line %-3s namespace conflict with mnemonic\n", strconv.Itoa(realLineNumber))
											errFlag = true
										}
										t.Operand = wordsInLine[2]
										a.IntermediateFileContainer = append(a.IntermediateFileContainer, t)
										a.ObjectCodeContainer.End.FirstCanRunInstructionSymbol = wordsInLine[1]
									} else {
										errFlag = true
										Logger.Printf("line %-3s first can run instruction location invalid\n", strconv.Itoa(realLineNumber))
										errFlag = true
									}
									goto end

								default:
									errFlag = true
									Logger.Printf("line %-3s mnemonic/opcode invalid\n", strconv.Itoa(realLineNumber))
								}

							} else {
								t.Opcode = wordsInLine[1]
								if conflict := a.ConflictWithMnemonic(wordsInLine[2]); conflict {
									errFlag = true
									Logger.Printf("line %-3s namespace conflict with mnemonic\n", strconv.Itoa(realLineNumber))
								} else {
									t.Operand = wordsInLine[2]

									err := a.SymbolTable.UpdateLocationCounter(&LocationCounter, "OPCODE", "")
									if err != nil {
										return err
									}
								}
							}
						}

					default:
						// 除了 BYTE 還有沒有別種可能？

						labelIsValid := true
						if a.SymbolTable.LabelFormatInValid(wordsInLine[0]) {
							errFlag = true
							Logger.Printf("line %-3s symbol/label format invalid\n", strconv.Itoa(realLineNumber))
							labelIsValid = false
						}
						if a.SymbolTable.IsDuplicatedDefine(wordsInLine[0]) {
							errFlag = true
							Logger.Printf("line %-3s duplicated symbol/label\n", strconv.Itoa(realLineNumber))
							labelIsValid = false
						}
						if conflict := a.ConflictWithMnemonic(wordsInLine[0]); conflict {
							errFlag = true
							Logger.Printf("line %-3s namespace conflict with mnemonic\n", strconv.Itoa(realLineNumber))
							labelIsValid = false
						}

						if labelIsValid == true {
							t.Label = wordsInLine[0]
							a.SymbolTable.HashTable[t.Label] = fmt.Sprintf("%x", LocationCounter)
						}

						if wordsInLine[1] == "BYTE" {
							t.Opcode = wordsInLine[1]
							if strings.HasSuffix(statementWithoutComment, "'") {
								if strings.HasPrefix(wordsInLine[2], "C'") {
									operandContentPrefixIndex := strings.Index(statementInLine, "'") + 1
									operandContentSuffixIndex := strings.LastIndex(statementInLine, "'") - 1
									if operandContentSuffixIndex-operandContentPrefixIndex > -1 {
										operandInSlice := strings.Split(statementWithoutComment, "")
										operandContent := strings.Join(operandInSlice[operandContentPrefixIndex:operandContentSuffixIndex+1], "")
										t.Operand = "C'" + operandContent + "'"
										err := a.SymbolTable.UpdateLocationCounter(&LocationCounter, "BYTEWithC", operandContent)
										if err != nil {
											return err
										}
									} else {
										errFlag = true
										Logger.Printf("line %-3s operand length invalid\n", strconv.Itoa(realLineNumber))
									}
								} else if strings.HasPrefix(wordsInLine[2], "X'") {
									tmpOperand := ""
									for i := 2; i < wordsInLineLength; i++ {
										tmpOperand += strings.TrimSpace(wordsInLine[i])
									}
									operandContentPrefixIndex := strings.Index(statementInLine, "'") + 1
									operandContentSuffixIndex := strings.LastIndex(statementInLine, "'") - 1
									if operandContentSuffixIndex-operandContentPrefixIndex > -1 {
										operandInSlice := strings.Split(statementWithoutComment, "")
										operandContent := strings.Join(operandInSlice[operandContentPrefixIndex:operandContentSuffixIndex+1], "")
										if m, _ := regexp.MatchString("[^0-9A-Fa-f]", operandContent); m {
											errFlag = true
											Logger.Printf("line %-3s operand invalid\n", strconv.Itoa(realLineNumber))
										} else if operandContentSuffixIndex-operandContentPrefixIndex > -1 {
											if len(operandContent)%2 != 0 {
												errFlag = true
												Logger.Printf("line %-3s operand invalid\n", strconv.Itoa(realLineNumber))
											} else {
												t.Operand = tmpOperand
												err := a.SymbolTable.UpdateLocationCounter(&LocationCounter, "BYTEWithX", wordsInLine[2])
												if err != nil {
													return err
												}
											}
										}
									}
								} else {
									errFlag = true
									Logger.Printf("line %-3s operand invalid\n", strconv.Itoa(realLineNumber))
								}
							}
						} else {
							delete(a.SymbolTable.HashTable, wordsInLine[0])
							t.Label = ""
							errFlag = true
							Logger.Printf("line %-3s format invalid\n", strconv.Itoa(realLineNumber))
						}
					}

				}

				a.IntermediateFileContainer = append(a.IntermediateFileContainer, t)

				// delete(a.SymbolTable.HashTable, "key")

			}
		}
	}
end:
	if findSTART == false {
		errFlag = true
		Logger.Printf("cannot find START mnemonic\n")
	}
	if findEND == false {
		errFlag = true
		Logger.Printf("cannot find END mnemonic\n")
	}

	//fmt.Println("ProgramLength LocationCounter", LocationCounter)
	//fmt.Println("ProgramLength LocationCounter FormatInt", strconv.FormatInt(int64(LocationCounter), 16))
	//fmt.Println("ProgramLength fmt.Sprintf", fmt.Sprintf("%X", LocationCounter))

	a.ObjectCodeContainer.Header.ProgramLength = strings.ToUpper(strconv.FormatInt(int64(LocationCounter-StartLocation), 16))

	return nil
}

// Part2 .
func (a *Assembler) Part2() error {

	for _, statement := range a.IntermediateFileContainer {

		statement.ObjectCode, _ = a.OpcodeTable.GetValueTypeString(statement.Opcode)

		if statement.Opcode == "RSUB" {
			statement.ObjectCode += "0000"
		} else {
			location, ok := a.SymbolTable.HashTable[statement.Operand]
			if ok == true {
				statement.ObjectCode += location
			} else {
				if statement.Opcode == "BYTE" {
					b1 := strings.HasPrefix(statement.Operand, "C'")
					if b1 == true {
						tmp := strings.Replace(statement.Operand, "C'", "", 1)
						tmpSlice := strings.Split(tmp, "")
						operandContent := ""
						for i := 0; i < len(tmp)-1; i++ {
							operandContent += tmpSlice[i]
						}
						for _, r := range operandContent {
							statement.ObjectCode += strings.ToUpper(strconv.FormatInt(int64(r), 16))
						}
					}

					b2 := strings.HasPrefix(statement.Operand, "X'")
					if b2 == true {
						tmp := strings.Replace(statement.Operand, "X'", "", 1)
						tmp = strings.Replace(tmp, "'", "", 1)
						statement.ObjectCode = tmp
					}

				} else if statement.Opcode == "WORD" {
					i, err := strconv.Atoi(statement.Operand)
					if err == nil {
						statement.ObjectCode = FillPrefixWithZero(strings.ToUpper(strconv.FormatInt(int64(i), 16)), 6)
					}
				} else {
					if strings.HasSuffix(statement.Operand, ",X") {
						slice := strings.Split(statement.Operand, ",")
						partOPCode, err := a.OpcodeTable.GetValueTypeString(statement.Opcode)
						partOPCode2, ok2 := a.SymbolTable.HashTable[slice[0]]
						if err == nil && ok2 == true {
							d, _ := strconv.ParseInt("0x"+partOPCode+"0000", 0, 64)
							d2, _ := strconv.ParseInt("0x"+partOPCode2, 0, 64)
							d3, _ := strconv.ParseInt("0x8000", 0, 64)
							statement.ObjectCode = strings.ToUpper(strconv.FormatInt(d+d2+d3, 16))
						}
					}
				}
			}
		}
	}

	for line, statement := range a.IntermediateFileContainer {
		realLineNumber := line + 1
		if strings.HasSuffix(statement.Operand, ",X") {
			slice := strings.Split(statement.Operand, ",")
			_, ok := a.SymbolTable.HashTable[slice[0]]
			if !ok {
				errFlag = true
				Logger.Printf("line %-3s undefined symbol/label\n", strconv.Itoa(realLineNumber))
			}
		} else {
			switch statement.Opcode {
			case "START":
			case "END":
			case "WORD":
			case "RESW":
			case "RESB":
			case "BYTE":
			case "RSUB":
			default:
				if statement.Operand != "" {
					_, ok := a.SymbolTable.HashTable[statement.Operand]
					if !ok {
						if statement.IsPureCommentORBlank == false {
							errFlag = true
							Logger.Printf("line %-3s undefined symbol/label", strconv.Itoa(realLineNumber))
						}
					}
				}
			}
		}
	}

	fmt.Println()

	return nil
}

// Part3 .
func (a *Assembler) Part3() {
	headerData := a.ObjectCodeContainer.GetHeaderData()
	textData := a.ObjectCodeContainer.GetTextData(a)
	endData := a.ObjectCodeContainer.GetEndData(a.SymbolTable)

	a.ObjectCodeContainer.ObjectCodeRecord = headerData + textData + endData
}
