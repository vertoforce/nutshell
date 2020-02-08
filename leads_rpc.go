package nutshell

import "fmt"

// GetLead ...
func (c *Client) GetLead(id int) (*Lead, error) {
	l := &Lead{}
	err := c.rpc.Call("getLead", map[string]interface{}{"leadId": id}, l)
	if err != nil {
		return nil, err
	}
	return l, nil
}

// SearchLeads ...
func (c *Client) SearchLeads(query string, limit int) (LeadsList, error) {
	limit = clampLimit(limit)
	args := map[string]interface{}{
		"string": query,
		"limit":  limit,
	}
	out := LeadsList{}
	err := c.rpc.Call("searchLeads", args, &out)
	return out, err
}

// FindLeads ...
func (c *Client) FindLeads(query FindLeadsQuery, limit, page int) (LeadsList, error) {
	limit = clampLimit(limit)
	if page <= 0 {
		page = 1
	}

	args := map[string]interface{}{
		"query": query,
		"limit": limit,
		"page":  page,
	}

	out := LeadsList{}
	err := c.rpc.Call("findLeads", args, &out)
	return out, err
}

// NewLead ...
func (c *Client) NewLead(nl UpsertLead) (*Lead, error) {
	return nil, fmt.Errorf("nope")
}
