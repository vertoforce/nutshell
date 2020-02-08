package nutshell

// NewCompany ...
func (c *Client) NewCompany(nc UpsertCompany) (*Company, error) {
	// can't add note while creating a company
	nc.Note = nil
	out := &Company{}
	err := c.rpc.Call("newAccount", map[string]interface{}{"account": nc}, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchCompanies ...
func (c *Client) SearchCompanies(query string, limit int) (CompaniesList, error) {
	out := CompaniesList{}
	args := map[string]interface{}{
		"string": query,
		"limit":  limit,
	}
	err := c.rpc.Call("searchAccounts", args, out)
	return out, err
}

// GetCompany returns the specified Company
func (c *Client) GetCompany(id int) (*Company, error) {
	out := &Company{}
	args := map[string]interface{}{
		"accountId": id,
	}
	err := c.rpc.Call("getAccount", args, out)
	return out, err
}

// FindAccountTypes ...
func (c *Client) FindAccountTypes(limit, page int) (StubList, error) {
	out := StubList{}
	limit = clampLimit(limit)
	if page <= 0 {
		page = 1
	}

	args := map[string]interface{}{
		"limit": limit,
		"page":  page,
	}

	err := c.rpc.Call("findAccountTypes", args, &out)
	return out, err
}

// FindCompanies ...
func (c *Client) FindCompanies(query FindCompanyQuery, limit, page int) (CompaniesList, error) {
	out := CompaniesList{}
	limit = clampLimit(limit)

	if page <= 0 {
		page = 1
	}

	args := map[string]interface{}{
		"query": query,
		"limit": limit,
		"page":  page,
	}

	err := c.rpc.Call("findAccounts", args, &out)
	return out, err
}

// EditCompany edits an account.
//
// Fields which allow multiples (phone, email, URL, etc.) will be completely replaced by whatever data you
// supply, if you supply any data for the field. (Eg. the new phone array replaces all previously known phone
// numbers.) If you are updating a multi-value field, you must include all values you wish to keep (not just
// the new values) for the field.
//
// If a note is specified, it will be added to the existing notes (preexisting notes are not affected).
//
// Tags can be edited as well – the specified array of tags will replace all existing tags, so you should
// retrieve and include existing tags if you only want to add a new one. Note that the tags must already exist.
func (c *Client) EditCompany(id int, uc UpsertCompany) (*Company, error) {
	out := &Company{}
	args := map[string]interface{}{
		"accountId": id,
		"account":   uc,
	}
	err := c.rpc.Call("editAccount", args, out)
	return out, err
}
