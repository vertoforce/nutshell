package nutshell

type Stub struct {
	Stub       bool   `json:"stub"`
	ID         int    `json:"id"`
	Rev        string `json:"rev"`
	EntityType string `json:"entityType"`
	Name       string `json:"name"`
}

type Phone struct {
	CountryCode string      `json:"countryCode,omitempty"`
	Number      string      `json:"number,omitempty"`
	Extension   interface{} `json:"extension,omitempty"`
}

type Territory struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Name struct {
	GivenName   string      `json:"givenName,omitempty"`
	FamilyName  string      `json:"familyName,omitempty"`
	Salutation  interface{} `json:"salutation,omitempty"`
	DisplayName string      `json:"displayName,omitempty"`
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
