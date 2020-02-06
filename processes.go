package nutshell

type Process struct {
	Stub
	Type               string      `json:"type"`
	StartedTime        string      `json:"startedTime"`
	ClosedTime         interface{} `json:"closedTime"`
	CurrentMilestoneID string      `json:"currentMilestoneId"`
	ProcessTemplateID  string      `json:"processTemplateId"`
}

type ProcessList []Process
