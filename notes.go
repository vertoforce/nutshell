package nutshell

type Note struct {
	CreatedTime string `json:"createdTime"`
	Date        string `json:"date"`
	EntityType  string `json:"entityType"`
	ID          int64  `json:"id"`
	Note        string `json:"note"`
	NoteHTML    string `json:"noteHtml"`
	NoteMarkup  string `json:"noteMarkup"`
	OriginID    int64  `json:"originId"`
	Rev         string `json:"rev"`
	Stub        bool   `json:"stub"`
	User        User   `json:"user"`
}

type User struct {
	CreatedTime     string   `json:"createdTime"`
	Emails          []string `json:"emails"`
	EntityType      string   `json:"entityType"`
	ID              int64    `json:"id"`
	IsAdministrator bool     `json:"isAdministrator"`
	IsEnabled       bool     `json:"isEnabled"`
	ModifiedTime    string   `json:"modifiedTime"`
	Name            string   `json:"name"`
	Rev             string   `json:"rev"`
	Stub            bool     `json:"stub"`
}

type NotesList []Note
