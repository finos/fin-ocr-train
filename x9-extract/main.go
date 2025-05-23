// Copyright (c) 2024 Capital One
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/moov-io/imagecashletter"
)

const (
	maxReaderBufSize = 64 * 1024 * 1024
)

type CheckInfo struct {
	Id                     string `json:"id"`
	FileName               string `json:"fileName"`
	FileSeqNo              int    `json:"fileSeqNo"`
	RoutingNumber          string `json:"routingNumber"`
	AccountNumber          string `json:"accountNumber"`
	CheckNumber            string `json:"checkNumber"`
	AuxiliaryOnUs          string `json:"auxiliaryOnUs"`
	PayorBankRoutingNumber string `json:"payorBankRoutingNumber"`
	PayorBankCheckDigit    string `json:"payorBankCheckDigit"`
	OnUs                   string `json:"onUs"`
}

/*
 * This processes one or more X9 files and extracts the image and JSON for each check.
 * Each image is written to a file named "<x9FilePrefix>-<num>.tiff" and
 * each JSON to a file named  "<x9FilePrefix>-<num>.json" where <num> is a monitonically increasing number
 * beginning at 1 for each X9 file.
 */

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: x9-extract <outputDir> <x9File1> [<x9File2> ...]")
		os.Exit(1)
	}
	dir := os.Args[1]
	// Create the output directory if it doesn't already exist
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Printf("Failed to create output directory %s: %s", dir, err)
		return
	}
	// Loop over each x9 file
        totalCheckCount := 0
	for _, filePath := range os.Args[2:] {
                fileName := filepath.Base(filePath)
		reader, err := os.Open(filePath)
		if err != nil {
			log.Println(err)
			return
		}
		bufSizeOpt := imagecashletter.BufferSizeOption(maxReaderBufSize)
		r := imagecashletter.NewReader(reader, imagecashletter.ReadVariableLineLengthOption(), imagecashletter.ReadEbcdicEncodingOption(), bufSizeOpt)
		f, err := r.Read()
		if err != nil {
			log.Printf("ERROR: Failed to read file %s: %v\n", filePath, err)
			return
		}
		// Loop over each cash letter in the x9 file
                fileSeqNo := 1
		for _, cl := range f.CashLetters {
			log.Printf("Processing %d checks from %s\n", cl.CashLetterControl.CashLetterItemsCount, filePath)
			// Loop over each bundle in the cash letter
			for _, b := range cl.Bundles {
				// Loop over each check in the bundle
				for _, c := range b.Checks {
					id := fmt.Sprintf("check-%d", totalCheckCount+1 )
					routingNumber := c.PayorBankRoutingNumber + c.PayorBankCheckDigit
					onUs := strings.Split(c.OnUs, "/")
					accountNumber := onUs[0]
					checkNumber := c.AuxiliaryOnUs
					if len(checkNumber) == 0 {
						checkNumber = onUs[1]
					}
					ci := CheckInfo{
						Id:                     id,
						FileName:               fileName,
						FileSeqNo:              fileSeqNo,
						RoutingNumber:          routingNumber,
						AccountNumber:          accountNumber,
						CheckNumber:            checkNumber,
						AuxiliaryOnUs:          c.AuxiliaryOnUs,
						PayorBankRoutingNumber: c.PayorBankRoutingNumber,
						PayorBankCheckDigit:    c.PayorBankCheckDigit,
						OnUs:                   c.OnUs}
					jsonCheck, err := json.MarshalIndent(ci, "", "   ")
					if err != nil {
						log.Printf("Failed to convert check to JSON: %s", err)
						return
					}
					err = ioutil.WriteFile(filepath.Join(dir, id+".tiff"), []byte(c.ImageViewData[0].ImageData), 0644)
					if err != nil {
						log.Printf("Failed to create file: %s", err)
						return
					}
					err = ioutil.WriteFile(filepath.Join(dir, id+".json"), []byte(jsonCheck), 0644)
					if err != nil {
						log.Printf("Failed to create file: %s", err)
						return
					}
					fileSeqNo++
					totalCheckCount++
				}
			}
		}
	}
	log.Printf("Finished extracting %d total checks into directory %s\n", totalCheckCount, dir)
}
