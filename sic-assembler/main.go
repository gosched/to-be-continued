package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	// Logger .
	Logger *log.Logger
)

func main() {

	// go run *.go
	// go run *.go -input source-2.txt
	// go run *.go -input source-err.txt

	inputFilePath := flag.String("input", "source.txt", "source code file path")

	flag.Parse()

	now := time.Now()

	logFile, err := os.OpenFile("./log/"+"errors.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("open log file err", err)
	}
	defer logFile.Close()

	Logger = log.New(os.Stdout, "Info:", log.Ldate|log.Ltime|log.Lshortfile)

	a := NewAssembler()

	err = a.InitAssembler("./opcode-table/"+"opcode-table.txt", "sic")
	if err != nil {
		Logger.Fatalln("InitAssembler", err)
	}

	fmt.Println("inputFilePath:", *inputFilePath)
	err = a.LoadSourceCode("./input/" + *inputFilePath)
	if err != nil {
		Logger.Fatalln("LoadSourceCode", err)
	}

	err = a.CreateRecordsData()
	if err != nil {
		Logger.Fatalln("CreateRecordsData", err)
	}

	err = a.ExportRecords("./result/" + "result.txt")
	if err != nil {
		Logger.Fatalln("ExportRecords", err)
	}

	if errFlag == false {
		fmt.Println(a.ObjectCodeContainer.ObjectCodeRecord)
		fmt.Println("time cost", time.Since(now))
	} else {
		/*
			for _, v := range a.IntermediateData {
				fmt.Printf("%-10s %-10s %-15s %-15s %-15s %-15s\n", v.LineNumber, v.Location, v.Label, v.Opcode, v.Operand, v.ObjectCode)
			}
			fmt.Println()
		*/
	}

}
