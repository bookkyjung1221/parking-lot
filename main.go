package main

import (
	"bufio"
	"os"
	"parkingLot/constants"
	"parkingLot/processor"
	"strings"
)

func main() {

	parkingServiceProcessor := new(processor.Processor)

	if len(os.Args) > 1 && os.Args[1] != "" {
		FileProcessor(os.Args[1])
	} else {
		for {
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			input = strings.TrimRight(input, "\r\n")
			if input == constants.Exit {
				break
			}
			parkingServiceProcessor.Execute(input)
		}
	}
}

func FileProcessor(fileName string) error {

	parkingServiceProcessor := new(processor.Processor)
	cmdFile, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer cmdFile.Close()

	cmdScanner := bufio.NewScanner(cmdFile)
	for cmdScanner.Scan() {
		cmdString := cmdScanner.Text()
		parkingServiceProcessor.Execute(cmdString)
	}

	if err := cmdScanner.Err(); err != nil {
		return err
	}
	return nil
}
