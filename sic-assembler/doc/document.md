# SIC

## Function

| Assembler function | OpcodeTable | SymbolTable |
| - | - | - |
| load source code | load mnemonics and object code | store symbol and location |
| detect and report syntax errors | store mnemonics and object code | storage the location of the symbol |
| convert to object code | find the opcode code of the opcode | find the location of the symbol |
| export object program | remove mnemonics and object code | remove symbol and location |

---

## Main data structure

| OpcodeTable | SymbolTable | SourceCodeData |
| - | - | - |
| hash table | hash table | Statements (string slice) in struct |

---

| IntermediateFileContainer | ObjectCodeContainer |
| - | - |
| StatementRecords (struct slice) in struct | ObjectCodeRecord (string) in struct |

---

| StatementRecord | | | | | |
| - | - | - | - | - | - |
| LineNumber (string) | Location (string) | Label (string) | Opcode (string) | Operand (string) | ObjectCode (string) |

---

## Project Structure

* unix executable file (assembler)
* golang source code (*.go)
* doc
  * document
  * pseudo-code.txt
* input
  * source.txt
* log
  * errors.log
* opcode-table
  * opcode-table.txt
* result
  * result.txt

---