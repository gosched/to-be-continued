package main

// StatementRecord .
type StatementRecord struct {
	LineNumber           string
	Location             string
	Label                string
	Opcode               string
	Operand              string
	ObjectCode           string
	IsPureCommentORBlank bool
}
