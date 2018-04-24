package main

import (
	"fmt"
	"time"

	"github.com/immesys/bw2bind"
	bw2 "gopkg.in/immesys/bw2bind.v5"
)

type assertFunc func(bw2bind.PayloadObject) bool

type StateMessage struct {
	HeatingValue float64
	CoolingValue float64
	Override     bool
	Mode         int
	Fan          bool
}

type PelicanState struct {
	Temperature     float64 `msgpack:"temperature"`
	RelHumidity     float64 `msgpack:"relative_humidity"`
	HeatingSetpoint float64 `msgpack:"heating_setpoint"`
	CoolingSetpoint float64 `msgpack:"cooling_setpoint"`
	Override        bool    `msgpack:"override"`
	Fan             bool    `msgpack:"fan"`
	Mode            int32   `msgpack:"mode"`
	State           int32   `msgpack:"state"`
	Time            string  `msgpack:"time"`
}

func main() {
	sampleStruct := StateMessage{70, 78, true, 2, false}
	test("s.pelican", "john/test", "i.xbos.thermostat", "2.1.1.0", sampleStruct, sampleAsseration)
}

/*
  To be defined by the user. The only things a user would need to modify
  is the struct type of the "getStatus" variable
*/
func sampleAsseration(input bw2bind.PayloadObject) bool {
	var getStatus PelicanState
	valueIntoErr := input.(bw2.MsgPackPayloadObject).ValueInto(&getStatus)
	if valueIntoErr != nil {
		// return false, fmt.Errorf("Payload Object could not be interpretted into Pelican State")
		return false
	}
	return true
}

/*
Parameters:
- clientName: Name of thermostat for creating new service client (i.e. s.pelican, s.national-weather-service)
- baseURI: Bosswave URI for subscribing to signals and publishing to slots
- ponum: Expected structure of the payload package object
- input: Input for creating payload package object
- assert: Checks whether output from test is correct / what is expected
*/
func test(clientName, baseURI, interfaceName, ponum string, packStruct interface{}, assert assertFunc) (bool, error) {
	// Part 1: Establish Bosswave Connection
	fmt.Println("Hello World")
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
	// Part B: Use Discovery Function to get instance

	// Part 4: Send Payload Packages and Parse Outputs
	prefix := "410test" // To Replace
	interfaceClient := serviceClient.AddInterface(prefix, interfaceName)

	// Subscribe to Slot (Testing Driver's Get Function)
	expected := false
	subscribeErr := interfaceClient.SubscribeSignal("info", func(message *bw2.SimpleMessage) {
		subscribePO := message.GetOnePODF(ponum)
		if subscribePO != nil {
			expected = assert(subscribePO)
		} else {
			expected = false
			// return false, fmt.Errorf("Payload Object from Signal Subscription does not fit the expected format")
		}
	})

	// Publish Payload (Testing Driver's Set Function)
	for i := 0; i < 5; i++ {
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
