package main

import (
	"bufio"
	"os"
)

// SourceCodeData .
type SourceCodeData struct {
	Statements []string
}

// Init .
func (SCD *SourceCodeData) Init() {
	SCD.Statements = []string{}
}

// Load .
func (SCD *SourceCodeData) Load(sourceCodePath string) error {
	testFile, err := os.OpenFile(sourceCodePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer testFile.Close()

	testFileScanner := bufio.NewScanner(testFile)
	for testFileScanner.Scan() {
		line := testFileScanner.Text()
		SCD.Statements = append(SCD.Statements, line)

		/*
			if line != "" {
				words := strings.Fields(line)
				isComment := strings.Contains(words[0], ".")
				if words[0] != "." && isComment == false {
					tmp := strings.Split(line, ".")
					SCD.Statements = append(SCD.Statements, tmp[0])
					SCD.CountLinesOfCode++
				}
			}
		*/
	}

	return nil
}
