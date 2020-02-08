package nutshell

// API uses 'Account' to refer to 'Company'

type UpsertLink struct {
	Relationship string `json:"relationship,omitempty"`
	ID           int    `json:"id,omitempty"`
}

type UpsertCompany struct {
	Name          string     `json:"name,omitempty"`
	Owner         *OwnerLink `json:"owner,omitempty"`
	IndustryId    string     `json:"industryId,omitempty"`
	AccountTypeId string     `json:"accountTypeId,omitempty"`
	TerritoryId   string     `json:"territoryId,omitempty"`
	URL           []string   `json:"url,omitempty"`
	Phone         []string   `json:"phone,omitempty"`

	Address      *Address     `json:"address,omitempty"`
	CustomFields CustomFields `json:"customFields,omitempty"`

	Note *string `json:"note,omitempty"`

	Leads    []UpsertLink `json:"leads,omitempty"`
	Contacts []UpsertLink `json:"contacts,omitempty"`
}

type Company struct {
	ID                int               `json:"id,omitempty"`
	EntityType        EntityType        `json:"entityType,omitempty"`
	Rev               string            `json:"rev,omitempty"`
	Name              string            `json:"name,omitempty"`
	Stub              bool              `json:"stub,omitempty"`
	Description       string            `json:"description,omitempty"`
	HtmlURL           string            `json:"htmlUrl,omitempty"`
	Tags              TagsList          `json:"tags,omitempty"`
	Industry          *Stub             `json:"industry,omitempty"`
	AccountType       *Stub             `json:"accountType,omitempty"`
	Leads             LeadsList         `json:"leads,omitempty"`
	Contacts          ContactsList      `json:"contacts,omitempty"`
	Phone             PhoneList         `json:"phone,omitempty"`
	URLs              map[string]string `json:"url,omitempty"`
	Emails            map[string]string `json:"email,omitempty"`
	Avatar            map[string]string `json:"avatar,omitempty"`
	Territory         *Stub             `json:"territory,omitempty"`
	LastContactedDate string            `json:"lastContactedDate,omitempty"`
	DeletedTime       string            `json:"deletedTime,omitempty"`
	ModifiedTime      string            `json:"modifiedTime,omitempty"`
	CreatedTime       string            `json:"createdTime,omitempty"`
}

type CompaniesList []Company

type FindCompanyQuery struct {
	AccountTypeId *int     `json:"account_type_id,omitempty"`
	Industry      []int    `json:"industry,omitempty"`
	Territory     []int    `json:"territory,omitempty"`
	Origin        []int    `json:"origin,omitempty"`
	HasOpenLeads  bool     `json:"has_open_leads,omitempty"`
	Tag           []string `json:"tag,omitempty"`
}
