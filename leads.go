package nutshell

type LeadStatus int

const (
	LeadOpen     LeadStatus = 0
	LeadPending             = 1
	LeadWon                 = 10
	LeadLost                = 11
	LeadCanceled            = 12
)

type ActivityCount string

const (
	ActCountScheduled ActivityCount = "0"
	ActCountLogged                  = "1"
	ActCountCancelled               = "2"
	ActCountOverdue                 = "-1"
)

type LeadsFilter int

const (
	LeadsMine LeadsFilter = 0
	LeadsTeam             = 1
	LeadsAll              = 2
)

type leadPriority int

const (
	LeadPriorityHot    leadPriority = 1
	LeadPriorityNormal              = 0
)

type Lead struct {
	ID                int                   `json:"id,omitempty"`
	EntityType        string                `json:"entityType,omitempty"`
	Rev               string                `json:"rev,omitempty"`
	ModifiedTime      string                `json:"modifiedTime,omitempty"`
	CreatedTime       string                `json:"createdTime,omitempty"`
	Name              string                `json:"name,omitempty"`
	Description       string                `json:"description,omitempty"`
	HTMLURL           string                `json:"htmlUrl,omitempty"`
	HTMLURLPath       string                `json:"htmlUrlPath,omitempty"`
	Creator           Creator               `json:"creator,omitempty"`
	ActivitiesCount   map[ActivityCount]int `json:"activitiesCount,omitempty"`
	PrimaryAccount    interface{}           `json:"primaryAccount,omitempty"`
	Milestone         Milestone             `json:"milestone,omitempty"`
	Stageset          Stageset              `json:"stageset,omitempty"`
	Status            LeadStatus            `json:"status,omitempty"`
	Confidence        int                   `json:"confidence,omitempty"`
	Completion        int                   `json:"completion,omitempty"`
	Urgency           string                `json:"urgency,omitempty"`
	IsOverdue         bool                  `json:"isOverdue,omitempty"`
	LastContactedDate interface{}           `json:"lastContactedDate,omitempty"`
	Market            Market                `json:"market,omitempty"`
	Assignee          Assignee              `json:"assignee,omitempty"`
	Sources           SourcesList           `json:"sources,omitempty"`
	Competitors       CompetitorsList       `json:"competitors,omitempty"`
	Products          ProductsList          `json:"products,omitempty"`
	Contacts          ContactsList          `json:"contacts,omitempty"`
	Accounts          AccountsList          `json:"accounts,omitempty"`
	Tags              TagsList              `json:"tags,omitempty"`
	Priority          int                   `json:"priority,omitempty"`
	CustomFields      CustomFields          `json:"customFields,omitempty"`
	Processes         ProcessList           `json:"processes,omitempty"`
	Notes             NotesList             `json:"notes,omitempty"`
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

// FindLeadsQuery ...
type FindLeadsQuery struct {
	Status      *LeadStatus  `json:"status,omitempty"`
	Filter      *LeadsFilter `json:"filter,omitempty"`
	AccountId   *int         `json:"account_id,omitempty"`
	ContactId   *int         `json:"contact_id,omitempty"`
	MilestoneId *int         `json:"milestone_id,omitempty"`
	StagesetId  *int         `json:"stageset_id,omitempty"`
	StagesetIds []int        `json:"stageset_ids,omitempty"`
	DueTime     *string      `json:"due_time,omitempty"`
	Assignee    []EntityType `json:"assignee,omitempty"`
	Origin      []int        `json:"origin,omitempty"`
	Source      []int        `json:"source,omitempty"`
	Number      int          `json:"number,omitempty"`
	Tag         TagsList     `json:"tag,omitempty"`
	Priority    leadPriority `json:"priority,omitempty"`
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
