package main

import (
	"fmt"

	bw2 "gopkg.in/immesys/bw2bind.v5"
)

/*
Parameters:
- clientName: Name of thermostat for creating new service client (i.e. s.pelican, s.national-weather-service)
*/

func main() {

}

func test(clientName, baseURI, interfaceName, ponum string) bool {
	// Part 1: Establish Bosswave Connection
	fmt.Println("Hello World")
	bwClient := bw2.ConnectOrExit("")
	bwClient.OverrideAutoChainTo(true)
	bwClient.SetEntityFromEnvironOrExit()

	// Part 2: Create Service & Interface Clients
	serviceClient := bwClient.NewServiceClient(baseURI, clientName)

	// Part 3: Create Payload Packages

	// Part 4: Send Payload Packages and Parse Outputs

	// Part 5: Use user function to check whether output is correct
	// Return "True" to indicate test was passed, "False" for visa versa
	return true
}
