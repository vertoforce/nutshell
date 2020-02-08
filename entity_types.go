package nutshell

type EntityType string

const (
	EntityActivity                EntityType = "Activities"
	EntityActivityType                       = "Activity_Types"
	EntityAssignment                         = "Assignments"
	EntityBackup                             = "Backups"
	EntityContact                            = "Contacts"
	EntityCompetitor                         = "Competitors"
	EntityCompany                            = "Accounts"
	EntityCompanyType                        = "Account_Type"
	EntityDelay                              = "Delays"
	EntityDelayApplication                   = "Delays"
	EntityEmail                              = "Emails"
	EntityIndustry                           = "Industries"
	EntityLead                               = "Leads"
	EntityLeadCompanyMap                     = EntityCompany
	EntityLeadCompetitorMap                  = EntityCompetitor
	EntityLeadContactMap                     = EntityContact
	EntityLeadOutcome                        = "Lead_Outcomes"
	EntityLeadProductMap                     = EntityProduct
	EntitySourceMap                          = EntitySource
	EntityMarket                             = "Markets"
	EntityMilestone                          = "Milestones"
	EntityNote                               = "Notes"
	EntityOrigin                             = "Origins"
	EntityProcess                            = "Processes"
	EntityProduct                            = "Products"
	EntitySource                             = "Sources"
	EntityStageset                           = "Stagesets"
	EntityStep                               = "Steps"
	EntityStepChosenEntity                   = EntityContact
	EntityStepTemplateRequirement            = EntityContact
	EntityTask                               = "Tasks"
	EntityTeam                               = "Teams"
	EntityUser                               = "Users"
)

type TagType int

const (
	TagLead    TagType = 1
	TagContact         = 2
	TagCompany         = 3
)

type EntityAttribute string

const (
	AttrCurrency EntityAttribute = "Currency"
	AttrDate                     = "Date"
	AttrEmail                    = "Email"
	AttrEnum                     = "Enum"
	AttrLocation                 = "Location"
	AttrPhone                    = "Phone"
	AttrString                   = "String"
	AttrText                     = "Text"
	AttrURI                      = "Url"
)

type EntityLink struct {
	EntityType EntityType `json:"entity_type"`
	Id         int        `json:"id"`
}
