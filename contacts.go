package nutshell

// API uses 'Contact' to refer to 'People'

// UpsertContact is used in `CreateContact` or `EditContact`
type UpsertContact struct {
	Name         string                 `json:"name,omitempty"`
	Description  string                 `json:"description,omitempty"`
	Phone        []string               `json:"phone,omitempty"`
	Email        []string               `json:"email,omitempty"`
	Address      []Address              `json:"address,omitempty"`
	Leads        []interface{}          `json:"leads,omitempty"`
	Accounts     []interface{}          `json:"accounts,omitempty"`
	CustomFields map[string]interface{} `json:"customFields,omitempty"`

	TerritoryId string     `json:"territory_id,omitempty"`
	Owner       *OwnerLink `json:"owner,omitempty"`
}

// Contact represents a contact ( aka "Person" ) in Nutshell
type Contact struct {
	Stub       bool   `json:"stub"`
	ID         int    `json:"id"`
	Rev        string `json:"rev"`
	EntityType string `json:"entityType"`
	// Name is sometimes returned as a single string and sometimes
	// as a map[string]string
	Name           interface{} `json:"name,omitempty"`
	HtmlURL        string      `json:"htmlUrl,omitempty"`
	Description    string      `json:"description,omitempty"`
	Tags           []string    `json:"tags,omitempty"`
	Creator        Creator     `json:"creator,omitempty"`
	ContactedCount int         `json:"contactedCount,omitempty"`

	Leads             LeadsList          `json:"leads,omitempty"`
	Accounts          AccountsList       `json:"accounts,omitempty"`
	Address           map[string]Address `json:"address,omitempty"`
	Phone             map[string]Phone   `json:"phone,omitempty"`
	URL               map[string]string  `json:"url,omitempty"`
	Email             map[string]string  `json:"email,omitempty"`
	Notes             NotesList          `json:"notes,omitempty"`
	Avatar            map[string]string  `json:"avatar,omitempty"`
	Territory         Territory          `json:"territory,omitempty"`
	LastContactedDate interface{}        `json:"lastContactedDate,omitempty"`
	DeletedTime       string             `json:"deletedTime,omitempty"`
	ModifiedTime      string             `json:"modifiedTime,omitempty"`
	CreatedTime       string             `json:"createdTime,omitempty"`
}

type ContactsList []Contact

// HasEmail ...
func (c Contact) HasEmail(e string) bool {
	for _, ee := range c.Email {
		if ee == e {
			return true
		}
	}
	return false
}
