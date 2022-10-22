package nutshell

import (
	"encoding/json"
)

type Stub struct {
	Stub       bool   `json:"stub"`
	ID         int    `json:"id"`
	Rev        string `json:"rev"`
	EntityType string `json:"entityType"`
	Name       Name   `json:"name"`
}

type Phone struct {
	CountryCode string      `json:"countryCode,omitempty"`
	Number      string      `json:"number,omitempty"`
	Extension   interface{} `json:"extension,omitempty"`
}

type Territory struct {
	ID   interface{} `json:"id,omitempty"`
	Name string      `json:"name,omitempty"`
}

type Name struct {
	GivenName   string `json:"givenName,omitempty"`
	FamilyName  string `json:"familyName,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

// Sometimes name comes in as single string, sometimes full Name struct
func (s *Name) UnmarshalJSON(b []byte) error {
	var v interface{}
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	if v == nil {
		return nil
	}

	// Parsing according to if it is a struct or string
	switch o := v.(type) {
	case map[string]interface{}:
		s.GivenName, _ = o["givenName"].(string)
		s.FamilyName, _ = o["familyName"].(string)
		s.DisplayName, _ = o["displayName"].(string)
	case string:
		s.DisplayName = o
	}

	return nil
}

type OwnerLink struct {
	EntityType string `json:"entityType,omitempty"`
	ID         int    `json:"id,omitempty"`
}

type Creator struct {
	Stub            bool     `json:"stub"`
	ID              int      `json:"id"`
	Rev             string   `json:"rev"`
	EntityType      string   `json:"entityType"`
	ModifiedTime    string   `json:"modifiedTime"`
	CreatedTime     string   `json:"createdTime"`
	Name            string   `json:"name"`
	Emails          []string `json:"emails"`
	IsEnabled       bool     `json:"isEnabled"`
	IsAdministrator bool     `json:"isAdministrator"`
}
