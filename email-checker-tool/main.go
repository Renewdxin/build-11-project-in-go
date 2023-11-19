package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// main function reads input domains from stdin and calls checkDomain function for each domain
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord\n")

	// Loop through each line of input from stdin
	for scanner.Scan() {
		// Call checkDomain function for each domain
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input:%v \n", err)
	}
}

// checkDomain function performs DNS lookups to check for MX, SPF, and DMARC records for a given domain
func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// Perform MX record lookup for the domain
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Println("Error:%v\n", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	// Perform TXT record lookup for the domain
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Println("Error:%v\n", err)
	}

	// Check for SPF record in the TXT records
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	// Perform TXT record lookup for DMARC record (_dmarc subdomain)
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Println("Error:%v\n", err)
	}

	// Check for DMARC record in the TXT records
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
