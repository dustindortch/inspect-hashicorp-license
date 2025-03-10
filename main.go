package main

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type License struct {
	ID              string              `json:"license_id"`
	CustomerID      string              `json:"customer_id"`
	InstallationID  string              `json:"installation_id"`
	IssueTime       string              `json:"issue_time"`
	StartTime       string              `json:"start_time"`
	ExpirationTime  string              `json:"end_time"`
	TerminationTime string              `json:"termination_time"`
	Product         string              `json:"product"`
	Flags           map[string][]string `json:"flags"`
}

func main() {
	var file string
	flag.StringVar(&file, "file", "", "License file to inspect")
	flag.Parse()

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	stat, _ := f.Stat()
	data := make([]byte, stat.Size())
	_, err = f.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	base32string, _ := base32.StdEncoding.DecodeString(string(data[2:]))
	firstField := strings.Split(string(base32string), ".")[0]
	base64string, _ := base64.StdEncoding.DecodeString(firstField)
	var jsonData License
	err = json.Unmarshal(base64string, &jsonData)
	if err != nil {
		log.Fatal(err)
	}

	output, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", string(output))
}
