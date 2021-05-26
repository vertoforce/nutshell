package nutshell

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type isCustomField interface {
	_isCustom()
}

type CustomFields map[string]interface{}
type UpsertCustomFields map[string]isCustomField

func (cf CustomFields) GetStringSlice(name string) []string {
	iSlice := cf.GetSlice(name)

	ret := []string{}
	for _, i := range iSlice {
		ret = append(ret, fmt.Sprintf("%s", i))
	}

	return ret
}

func (cf CustomFields) GetSlice(name string) []interface{} {
	t, ok := cf[name]
	if !ok {
		return nil
	}

	ret, ok := t.([]interface{})
	if !ok {
		return nil
	}

	return ret
}

func (cf CustomFields) GetString(name string) string {
	t, ok := cf[name]
	if !ok {
		return ""
	}

	ret, ok := t.(string)
	if !ok {
		return ""
	}

	return ret
}

func (cf CustomFields) GetFloat64(name string) float64 {
	strVal := cf.GetString(name)
	if strVal == "" {
		return 0
	}

	ret, err := strconv.ParseFloat(strVal, 64)
	if err != nil {
		return 0
	}

	return ret
}

// get ...
func (cf CustomFields) get(name string, out isCustomField) error {
	t, ok := cf[name]
	if !ok {
		return fmt.Errorf("no custom field named '%v'", name)
	}

	b, err := json.Marshal(t)
	if err != nil {
		return fmt.Errorf("unable to marshal")
	}

	if err = json.Unmarshal(b, out); err != nil {
		return fmt.Errorf("wrong type in custom field '%v'", name)
	}
	return nil
}

// GetTimestamp ...
func (cf CustomFields) GetTimestamp(name string) (*Timestamp, error) {
	ts := &Timestamp{}
	err := cf.get(name, ts)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

// GetAddress ...
func (cf CustomFields) GetAddress(name string) (*Address, error) {
	addr := &Address{}
	err := cf.get(name, addr)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

type Timestamp struct {
	Timestamp  string `json:"timestamp"`
	HumanShort string `json:"humanShort"`
	HumanDate  string `json:"humanDate"`
}

// _isCustom ...
func (t Timestamp) _isCustom() {}

type Address struct {
	Name             interface{} `json:"name,omitempty"`
	Location         Location    `json:"location,omitempty"`
	LocationAccuracy string      `json:"locationAccuracy,omitempty"`
	Address1         string      `json:"address_1,omitempty"`
	Address2         string      `json:"address_2,omitempty"`
	Address3         string      `json:"address_3,omitempty"`
	City             string      `json:"city,omitempty"`
	State            string      `json:"state,omitempty"`
	PostalCode       string      `json:"postalCode,omitempty"`
	Country          string      `json:"country,omitempty"`
	Timezone         string      `json:"timezone,omitempty"`
}

// _isCustom ...
func (a Address) _isCustom() {}

type Location struct {
	Longitude float64 `json:"longitude,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
}
