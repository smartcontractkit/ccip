package csv

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/dione"
)

func PrepareCsvFile(filePath string, headers []string) {
	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)
	var data [][]string
	data = append(data, headers)
	err = w.WriteAll(data)
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func AppendToFile(filePath string, records []dione.NodeWallet, chainName string, ENV dione.Environment) {
	var keys [][]string
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatalf("failed to open file %s: %s", filePath, err)
	}
	defer f.Close()
	for _, record := range records {
		row := []string{string(ENV), chainName, strconv.FormatUint(record.ChainID, 10), record.Address}
		keys = append(keys, row)
	}
	w := csv.NewWriter(f)
	w.WriteAll(keys)
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
