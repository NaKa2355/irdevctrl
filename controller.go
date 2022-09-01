package irdevctrl

import (
	"encoding/json"
)

type Feature int
type Features []Feature

const (
	Sending Feature = iota
	Receiving
)

type Controller interface {
	ReceiveIRData() (RawData, error)
	SendIRData(RawData) error
	GetBufferSize() uint16
	GetSupportingFeatures() Features
	Drop() error
}

func (features Features) MarshalJSON() ([]byte, error) {
	jsonSf := []string{}
	for _, feature := range features {

		if feature == Sending {
			jsonSf = append(jsonSf, "sending")
		}

		if feature == Receiving {
			jsonSf = append(jsonSf, "receiving")
		}
	}

	return json.Marshal(jsonSf)
}
