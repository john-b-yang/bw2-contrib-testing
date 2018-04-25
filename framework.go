package main

import (
	"fmt"
	"time"

	"github.com/immesys/bw2bind"
	bw2 "gopkg.in/immesys/bw2bind.v5"
)

type assertFunc func(bw2bind.PayloadObject) bool

type TestStruct struct {
	ClientName    string
	BaseURI       string
	InterfaceName string
	Ponum         string
	Prefix        string
}

/*
Parameters:
- clientName: Name of thermostat for creating new service client (i.e. s.pelican, s.national-weather-service)
- baseURI: Bosswave URI for subscribing to signals and publishing to slots
- ponum: Expected structure of the payload package object
- input: Input for creating payload package object
- assert: Checks whether output from test is correct / what is expected
*/
func TestDriver(params TestStruct, packStruct interface{}, validate assertFunc) (bool, error) {
	clientName := params.ClientName
	baseURI := params.BaseURI
	interfaceName := params.InterfaceName
	ponum := params.Ponum
	prefix := params.Prefix

	// Part 1: Establish Bosswave Connection
	bwClient := bw2.ConnectOrExit("")
	bwClient.OverrideAutoChainTo(true)
	bwClient.SetEntityFromEnvironOrExit()

	// Part 2: Create Service & Interface Clients
	serviceClient := bwClient.NewServiceClient(baseURI, clientName)

	// Part 3: Create Payload Packages
	// Part A: Create Message Pack
	payload, poErr := bw2.CreateMsgPackPayloadObject(bw2.FromDotForm(ponum), packStruct)
	if poErr != nil {
		return false, fmt.Errorf("Payload Object Creation Error: %v", poErr)
	}

	// Part 4: Send Payload Packages and Parse Outputs
	interfaceClient := serviceClient.AddInterface(prefix, interfaceName)

	// Subscribe to Slot (Testing Driver's Get Function)
	// TODO: Change expected to channel
	expected := false
	subscribeErr := interfaceClient.SubscribeSignal("info", func(message *bw2.SimpleMessage) {
		subscribePO := message.GetOnePODF(ponum)
		if subscribePO != nil {
			expected = validate(subscribePO)
		} else {
			expected = false
			// return false, fmt.Errorf("Payload Object from Signal Subscription does not fit the expected format")
		}
	})
	if subscribeErr != nil {
		return false, fmt.Errorf("Signal Subscription Error: %v", subscribeErr)
	}

	// Publish Payload (Testing Driver's Set Function)
	// TODO: Channel + Select (time.after, return channel after a certain delay)
	// i := 0; i < 5; i++
	for {
		publishErr := interfaceClient.PublishSlot("publish", payload)
		if publishErr != nil {
			return false, fmt.Errorf("Payload Object Publishing Error: %v", publishErr)
		}
		// TODO: How to get output from publish?
		time.Sleep(5 * time.Second)
	}

	// Return "True" to indicate test was passed, "False" for visa versa
	return expected, nil
}
