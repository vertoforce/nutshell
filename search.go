package nutshell

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

type SearchByEmailResult struct {
	Contacts ContactsList `json:"contacts,omitempty"`
	Accounts AccountsList `json:"accounts,omitempty"`
}

// SearchByEmail ...
func (c *Client) SearchByEmail(email string) (*SearchByEmailResult, error) {
	var res *SearchByEmailResult
	args := map[string]string{
		"emailAddressString": email,
	}
	var out interface{}
	err := c.rpc.Call("searchByEmail", args, out)
	spew.Dump(out, err)
	if out == nil {
		return nil, fmt.Errorf("no results")
	}
	spew.Dump(out)
	return res, err
}
