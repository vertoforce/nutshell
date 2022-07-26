package nutshell

import "fmt"

type Lead struct {
	ID                int             `json:"id,omitempty"`
	EntityType        string          `json:"entityType,omitempty"`
	Rev               string          `json:"rev,omitempty"`
	ModifiedTime      string          `json:"modifiedTime,omitempty"`
	CreatedTime       string          `json:"createdTime,omitempty"`
	Name              string          `json:"name,omitempty"`
	Description       string          `json:"description,omitempty"`
	HTMLURL           string          `json:"htmlUrl,omitempty"`
	HTMLURLPath       string          `json:"htmlUrlPath,omitempty"`
	Creator           Creator         `json:"creator,omitempty"`
	ActivitiesCount   map[string]int  `json:"activitiesCount,omitempty"`
	PrimaryAccount    interface{}     `json:"primaryAccount,omitempty"`
	Milestone         Milestone       `json:"milestone,omitempty"`
	Stageset          Stageset        `json:"stageset,omitempty"`
	Status            int             `json:"status,omitempty"`
	Confidence        int             `json:"confidence,omitempty"`
	Completion        int             `json:"completion,omitempty"`
	Urgency           string          `json:"urgency,omitempty"`
	IsOverdue         bool            `json:"isOverdue,omitempty"`
	LastContactedDate interface{}     `json:"lastContactedDate,omitempty"`
	Market            Market          `json:"market,omitempty"`
	Assignee          Assignee        `json:"assignee,omitempty"`
	Sources           SourcesList     `json:"sources,omitempty"`
	Competitors       CompetitorsList `json:"competitors,omitempty"`
	Products          ProductsList    `json:"products,omitempty"`
	Contacts          ContactsList    `json:"contacts,omitempty"`
	Accounts          AccountsList    `json:"accounts,omitempty"`
	Tags              TagsList        `json:"tags,omitempty"`
	Priority          int             `json:"priority,omitempty"`
	CustomFields      CustomFields    `json:"customFields,omitempty"`
	Processes         ProcessList     `json:"processes,omitempty"`
	Notes             NotesList       `json:"notes,omitempty"`
}

type LeadsList []Lead

type UpsertLead struct {
	PrimaryAccount struct {
		ID string `json:"id,omitempty"`
	} `json:"primaryAccount,omitempty"`
	DueTime string `json:"dueTime,omitempty"`
	Market  struct {
		ID string `json:"id,omitempty"`
	} `json:"market,omitempty"`
	Tags         []string           `json:"tags,omitempty"`
	Description  string             `json:"description,omitempty"`
	Accounts     AccountsList       `json:"accounts,omitempty"`
	Contacts     ContactsList       `json:"contacts,omitempty"`
	Products     ProductsList       `json:"products,omitempty"`
	Competitors  CompetitorsList    `json:"competitors,omitempty"`
	Sources      SourcesList        `json:"sources,omitempty"`
	Confidence   int                `json:"confidence,omitempty"`
	Assignee     Stub               `json:"assignee,omitempty"`
	CustomFields UpsertCustomFields `json:"customFields,omitempty"`
	Notes        []string           `json:"notes,omitempty"`
	Priority     int                `json:"priority,omitempty"`
	IsPending    bool               `json:"is_pending,omitempty"`
}

type Milestone struct {
	Stub
}

type Stageset struct {
	Stub
}

type Market struct {
	Stub
}

type Assignee struct {
	Stub
	ModifiedTime    string   `json:"modifiedTime,omitempty"`
	CreatedTime     string   `json:"createdTime,omitempty"`
	Emails          []string `json:"emails,omitempty"`
	IsEnabled       bool     `json:"isEnabled,omitempty"`
	IsAdministrator bool     `json:"isAdministrator,omitempty"`
}

type Contacts struct {
	Stub
	ModifiedTime string `json:"modifiedTime,omitempty"`
	CreatedTime  string `json:"createdTime,omitempty"`
	JobTitle     string `json:"jobTitle,omitempty"`
	Description  string `json:"description,omitempty"`
	Relationship string `json:"relationship,omitempty"`
}

// GetLead ...
func (c *Client) GetLead(id int) (*Lead, error) {
	l := &Lead{}
	err := c.rpc.Call("getLead", map[string]interface{}{"leadId": id}, l)
	if err != nil {
		return nil, err
	}
	return l, nil
}

// GetLead ...
func (c *Client) GetLeads() ([]Lead, error) {
	l := []Lead{}
	err := c.rpc.Call("searchLeads", map[string]interface{}{"string": ""}, &l)
	if err != nil {
		return nil, err
	}
	return l, nil
}

// NewLead ...
func (c *Client) NewLead(nl UpsertLead) (*Lead, error) {
	return nil, fmt.Errorf("nope")
}

func (c *Client) FindLeads(number int64) ([]Lead, error) {
	l := []Lead{}
	err := c.rpc.Call("findLeads", map[string]interface{}{"query": map[string]interface{}{"number": number}}, &l)
	if err != nil {
		return nil, err
	}
	return l, nil
}
