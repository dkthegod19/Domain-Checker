package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const spfPrefix = "v=spf1"
const sprPrefix = "v=DMARC1"
const dmarcDomain = "_dmarc."

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Domain Checker \n Domain \n HasMX \n HasSPF \n SPRRecord \n HasDemarc \n DmarcRecord \n")
	for scanner.Scan() {
		checkdomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error: Could not read from input: %v", err)
	}
}

func checkdomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, sprRecord string
	maxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(maxRecords) > 0 {
		hasMX = true
	}
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, spfPrefix) {
			hasSPF = true
			spfRecord = record
			break
		}
	}
	dmarcRecords, err := net.LookupTXT(dmarcDomain + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, sprPrefix) {
			hasDMARC = true
			sprRecord = record
			break
		}
	}
	fmt.Printf("Domain Checker \n Domain: %v \n HasMX: %v \n HasSPF: %v \n SPRRecord: %v \n HasDemarc: %v \n DmarcRecord: %v \n", domain, hasMX, hasSPF, spfRecord, hasDMARC, sprRecord)
}
