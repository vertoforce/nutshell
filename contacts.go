package nutshell

import "fmt"

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
	Stub

	HtmlURL     string   `json:"htmlUrl,omitempty"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Creator     Creator  `json:"creator,omitempty"`

	Leads          LeadsList          `json:"leads,omitempty"`
	Accounts       AccountsList       `json:"accounts,omitempty"`
	ContactedCount int                `json:"contactedCount,omitempty"`
	Address        map[string]Address `json:"address,omitempty"`
	Phone          map[string]Phone   `json:"phone,omitempty"`
	URL            map[string]string  `json:"url,omitempty"`
	Email          map[string]string  `json:"email,omitempty"`
	Notes          NotesList          `json:"notes,omitempty"`
	Avatar         map[string]string  `json:"avatar,omitempty"`
	// Territory         Territory          `json:"territory,omitempty"`
	LastContactedDate interface{}            `json:"lastContactedDate,omitempty"`
	DeletedTime       string                 `json:"deletedTime,omitempty"`
	ModifiedTime      string                 `json:"modifiedTime,omitempty"`
	CreatedTime       string                 `json:"createdTime,omitempty"`
	CustomFields      map[string]interface{} `json:"customFields"`
}

type ContactsList []Contact

// GetContact ...
func (c *Client) GetContact(id int) (*Contact, error) {
	con := &Contact{}

	// out := map[string]interface{}{}
	// err := c.rpc.Call("getContact", map[string]interface{}{"contactId": id}, &out)
	err := c.rpc.Call("getContact", map[string]interface{}{"contactId": id}, con)
	if err != nil {
		return nil, err
	}
	// spew.Dump(out)

	return con, nil
}

// GetContact ...
func (c *Client) FindContacts(query map[string]interface{}) (*Contact, error) {
	con := &Contact{}

	// out := map[string]interface{}{}
	// err := c.rpc.Call("getContact", map[string]interface{}{"contactId": id}, &out)
	err := c.rpc.Call("findContacts", map[string]interface{}{"query": query}, con)
	if err != nil {
		return nil, err
	}
	// spew.Dump(out)

	return con, nil
}

// NewContact ...
func (c *Client) NewContact(nc UpsertContact) (*Contact, error) {
	return nil, fmt.Errorf("not yet")
}
