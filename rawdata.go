package irdev

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Prefix int8

const (
	Milli Prefix = iota
	Micro
)

type Pulse struct {
	Prefix Prefix
	Width  int16
}

type RawData []Pulse

func (rd RawData) ConvertToRawData() (RawData, error) {
	return rd, nil
}

func (rd *RawData) UnmarshalJSON(data []byte) error {
	var err error = nil
	var width int16 = 0
	strRawData := struct {
		Pluses []string `json:"data"`
	}{}

	if err := json.Unmarshal(data, &strRawData); err != nil {
		return err
	}

	rawData := make([]Pulse, len(strRawData.Pluses))

	for i, pulse := range strRawData.Pluses {

		pulse = strings.ReplaceAll(pulse, " ", "")

		if _, err = fmt.Sscanf(pulse, "%dus", &width); err == nil {
			rawData[i].Prefix = Micro
			rawData[i].Width = width
			continue
		}

		if _, err = fmt.Sscanf(pulse, "%dms", &width); err == nil {
			rawData[i].Prefix = Milli
			rawData[i].Width = width

		}

		return fmt.Errorf("raw data's format is wrong: %s", err)

	}

	*rd = rawData
	return nil
}

func (rawData RawData) MarshalJSON() ([]byte, error) {
	strRawData := struct {
		Pluses []string `json:"data"`
	}{}

	strRawData.Pluses = make([]string, len(rawData))

	for i, pulse := range rawData {

		if pulse.Prefix == Micro {
			strRawData.Pluses[i] = fmt.Sprintf("%dus", pulse.Width)
			continue
		}

		if pulse.Prefix == Milli {
			strRawData.Pluses[i] = fmt.Sprintf("%dus", pulse.Width)
			continue
		}
	}

	return json.Marshal(strRawData)
}
