package nutshell

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

// NewContact ...
func (c *Client) NewContact(nc UpsertContact) (*Contact, error) {
	con := &Contact{}
	err := c.rpc.Call("newContact", map[string]interface{}{"contact": nc}, con)
	if err != nil {
		return nil, err
	}
	return con, nil
}

// SearchContacts ...
func (c *Client) SearchContacts(query string, limit int) (ContactsList, error) {
	out := ContactsList{}

	if limit > 100 {
		limit = 100
	}
	if limit < 1 {
		limit = 1
	}

	q := map[string]interface{}{
		"string": query,
		"limit":  limit,
	}
	err := c.rpc.Call("searchContacts", q, &out)
	return out, err
}
