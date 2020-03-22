package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// BKERHash is an algorithm invented by Brian Kernighan Dennis Ritchie

//  opcode(string) -> opcode(base 16)

// OpcodeTable .
type OpcodeTable struct {
	BKDRHashTable map[int32]string // input data collision uncheck
}

// BKDRHash .
func (ot *OpcodeTable) BKDRHash(opcodeString string) int32 {
	var seed, hash int32 = 131, 0

	for _, ch := range opcodeString {
		hash = hash*seed + ch
	}
	//fmt.Println(hash & 0x7FFFFFFF)
	return hash & 0x7FFFFFFF
}

// Init .
func (ot *OpcodeTable) Init() {
	ot.BKDRHashTable = make(map[int32]string)
}

// Load .
func (ot *OpcodeTable) Load(tablePath string) error {

	file, err := os.OpenFile(tablePath, os.O_RDWR, 0660)
	defer file.Close()

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txtRow := scanner.Text()
		tmpSlice := strings.Split(txtRow, " ")
		ot.BKDRHashTable[ot.BKDRHash(tmpSlice[0])] = tmpSlice[1]
	}

	return nil
}

// GetValueTypeString .
func (ot *OpcodeTable) GetValueTypeString(key string) (string, error) {
	tmpKey := ot.BKDRHash(key)
	value := ot.BKDRHashTable[tmpKey]
	if value == "" {
		err := errors.New("key not exist")
		return "", err
	}
	return value, nil
}

// GetValueTypeInt .
func (ot *OpcodeTable) GetValueTypeInt(key string) (int32, error) {
	tmpKey := ot.BKDRHash(key)
	value := ot.BKDRHashTable[tmpKey]
	if value == "" {
		err := errors.New("key not exist")
		return 0, err
	}
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return int32(result), nil
}

// NotConflictWithMnemonic .
func (ot *OpcodeTable) NotConflictWithMnemonic(symbol string) bool {
	if _, exist := ot.BKDRHashTable[ot.BKDRHash(symbol)]; exist {
		return false
	}
	switch symbol {
	case "START":
		return false
	case "END":
		return false
	case "WORD":
		return false
	case "RESW":
		return false
	case "RESB":
		return false
	case "BYTE":
		return false
	}
	return true
}

// Remove .
func (ot *OpcodeTable) Remove(key string) {
	realKey := ot.BKDRHash(key)
	if _, ok := ot.BKDRHashTable[realKey]; ok {
		delete(ot.BKDRHashTable, realKey)
	}
}
