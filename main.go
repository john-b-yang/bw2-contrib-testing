package main

import (
	"github.com/immesys/bw2bind"
	bw2 "gopkg.in/immesys/bw2bind.v5"
)

type StateMessage struct {
	HeatingValue float64
	CoolingValue float64
	Override     bool
	Mode         int
	Fan          bool
}

type PelicanState struct {
	Temperature     float64
	RelHumidity     float64
	HeatingSetpoint float64
	CoolingSetpoint float64
	Override        bool
	Fan             bool
	Mode            int32
	State           int32
	Time            string
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

func main() {
	sampleStruct := StateMessage{
		HeatingValue: 70,
		CoolingValue: 78,
		Override:     true,
		Mode:         2,
		Fan:          false,
	}

	sampleTestStruct := TestStruct{
		ClientName:    "s.pelican",
		BaseURI:       "john/test",
		InterfaceName: "i.xbos.thermostat",
		Ponum:         "2.1.1.0",
	}

	TestDriver(sampleTestStruct, sampleStruct, sampleAsseration)
}
