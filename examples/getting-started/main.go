// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/moov-io/paygate/examples/common"
)

func main() {
	fmt.Printf("using requestID %s\n\n", common.RequestID)

	// Create source customer
	sourceCustomer1, err := common.CreateCustomer("John", "Doe", "john.doe@moov.io")
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	fmt.Printf("Created source customer %s\n", sourceCustomer1.CustomerID)

	// Approve customer
	sourceCustomer1, err = common.ApproveCustomer(sourceCustomer1)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	fmt.Printf("Customer status is %s\n", sourceCustomer1.Status)

	// Create account
	sourceAccount1, err := common.CreateAccount(sourceCustomer1, "123456", common.TeachersFCU, "Checking")
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	fmt.Printf("Created source customer account %s\n", sourceAccount1.AccountID)

	// Approve account
	_, err = common.ApproveAccount(sourceCustomer1, sourceAccount1)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	fmt.Println("Approved source account")
	fmt.Println("===========")

	// Create source customer 2
	sourceCustomer2, err := common.CreateCustomer("Mary", "Jane", "mary.jane@moov.io")
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	fmt.Printf("Created source customer2 %s\n", sourceCustomer2.CustomerID)

	// Approve customer 2
	sourceCustomer2, err = common.ApproveCustomer(sourceCustomer2)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	fmt.Printf("Customer status is %s\n", sourceCustomer2.Status)

	// Create account 2
	sourceAccount2, err := common.CreateAccount(sourceCustomer2, "654321", common.ChaseCO, "Checking")
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	fmt.Printf("Created source customer account %s\n", sourceAccount2.AccountID)

	// Approve account 2
	_, err = common.ApproveAccount(sourceCustomer2, sourceAccount2)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	fmt.Println("Approved source account")
	fmt.Println("===========")

	// Create destination customer
	destinationCustomer, err := common.CreateCustomer("My Odfi", " ", "blah@moov.io")
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	fmt.Printf("Created destination customer %s\n", destinationCustomer.CustomerID)

	// Approve customer
	destinationCustomer, err = common.ApproveCustomer(destinationCustomer)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	fmt.Printf("Customer status is %s\n", destinationCustomer.Status)

	// Create account
	destinationCustomerAccount, err := common.CreateAccount(destinationCustomer, "322070381", common.EastWestBank, "Checking")
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	fmt.Printf("Created destination customer account %s\n", destinationCustomerAccount.AccountID)

	// Approve account
	_, err = common.ApproveAccount(destinationCustomer, destinationCustomerAccount)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	fmt.Println("Approved destination account")
	fmt.Println("===========")

	// Initiate a transfer
	transfer1, err := common.MakeTransfer(125, sourceCustomer1, sourceAccount1, destinationCustomer, destinationCustomerAccount)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	fmt.Printf("Transfer id is %s\n", transfer1.TransferID)

	transfer2, err := common.MakeTransfer(220, sourceCustomer2, sourceAccount2, destinationCustomer, destinationCustomerAccount)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	fmt.Printf("Transfer id is %s\n", transfer2.TransferID)

	time.Sleep(1 * time.Second)

	_, err = common.TriggerCutOff()
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	fmt.Println("===========")

	// Get transfer and display
	// transfer, err = common.GetTransfer(transfer.TransferID)
	// if err != nil {
	// 	log.Fatalf("ERROR: %v", err)
	// }
	// var indentedJson, _ = json.MarshalIndent(transfer, "", "  ")
	// fmt.Println(string(indentedJson))
	fmt.Println("")
	fmt.Println("Success! A Transfer was created.")
	fmt.Println("An ACH file was uploaded to a test FTP server at ./testdata/ftp-server/outbound/")
	fmt.Println("")
	fmt.Println("Uploaded files:")
	common.PrintServerFiles(filepath.Join("testdata", "ftp-server", "outbound"))
}
