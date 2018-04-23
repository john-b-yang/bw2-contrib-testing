package main

import (
	"fmt"

	bw2 "gopkg.in/immesys/bw2bind.v5"
)

/*
Parameters:
- clientName: Name of thermostat for creating new service client (i.e. s.pelican, s.national-weather-service)
- baseURI: Bosswave URI for subscribing to signals and publishing to slots
- ponum: Expected structure of the payload package object
- input: Input for creating payload package object
- assert: Checks whether output from test is correct / what is expected
*/

type assertFunc func(string) bool

// type testStruct struct {
// 	test int
// }

type StateMessage struct {
	HeatingValue float64
	CoolingValue float64
	Override     bool
	Mode         int
	Fan          bool
}

func main() {
	sampleStruct := StateMessage{70, 78, true, 2, false}
	test("s.pelican", "john/test", "i.xbos.thermostat", "2.1.1.0", sampleStruct, sampleAsseration)
}

// To be defined by the user
func sampleAsseration(input string) bool {
	return false
}

func test(clientName, baseURI, interfaceName, ponum string, class interface{}, assert assertFunc) bool {
	// Part 1: Establish Bosswave Connection
	fmt.Println("Hello World")
	bwClient := bw2.ConnectOrExit("")
	bwClient.OverrideAutoChainTo(true)
	bwClient.SetEntityFromEnvironOrExit()

	// Part 2: Create Service & Interface Clients
	serviceClient := bwClient.NewServiceClient(baseURI, clientName)

	// Part 3: Create Payload Packages
	// Part A: Create Message Pack

	// Part B: Use Discovery Function to get instance

	// Part 4: Send Payload Packages and Parse Outputs
	prefix := "410test" // To Replace
	interfaceClient := serviceClient.AddInterface(prefix, interfaceName)

	// Part 5: Use user function to check whether output is correct

	// Return "True" to indicate test was passed, "False" for visa versa
	return assert("To Replace With Struct / Output from Payload Object")
}
