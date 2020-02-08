package nutshell

import (
	"fmt"
)

var ErrNoContact = fmt.Errorf("no contact found for given email")

// GetContact retrieves the specified Contact
func (c *Client) GetContact(id int) (*Contact, error) {
	con := &Contact{}
	err := c.rpc.Call("getContact", map[string]interface{}{"contactId": id}, con)
	if err != nil {
		return nil, err
	}
	return con, nil
}

// NewContact ...
func (c *Client) NewContact(nc UpsertContact) (*Contact, error) {
	// tags can't be set during creation
	nc.Tags = TagsList{}

	con := &Contact{}
	err := c.rpc.Call("newContact", map[string]interface{}{"contact": nc}, con)
	if err != nil {
		return nil, err
	}
	return con, nil
}

// EditContact edits a contact.
// Fields which allow multiples (phone, email, URL, etc.) will be completely replaced by whatever
// data you supply, if you supply any data for the field. (Eg. the new phone array replaces all
// previously known phone numbers.) If you are updating a multi-value field, you must include all
// values you wish to keep (not just the new values) for the field.

// If a note is specified, it will be added to the existing notes (preexisting notes are not affected).

// Tags can be edited as well – the specified array of tags will replace all existing tags, so you
// should retrieve and include existing tags if you only want to add a new one. Note that the tags
// must already exist ( see NewTag )
func (c *Client) EditContact(id int, rev string, uc UpsertContact) (*Contact, error) {
	args := map[string]interface{}{
		"contactId": id,
		"rev":       rev,
		"contact":   uc,
	}

	con := &Contact{}
	err := c.rpc.Call("editContact", args, con)
	if err != nil {
		return nil, err
	}
	return con, nil
}

// FindContacts ...
func (c *Client) FindContacts(query FindContactQuery, limit, page int) (ContactsList, error) {
	out := ContactsList{}
	limit = clampLimit(limit)
	if page <= 0 {
		page = 1
	}
	args := map[string]interface{}{
		"query": query,
		"limit": limit,
		"page":  page,
	}
	err := c.rpc.Call("findContacts", args, &out)
	return out, err
}

// SearchContactByEmail ...
func (c *Client) SearchContactByEmail(query string) (*Contact, error) {
	res, err := c.SearchContacts(query, 1)
	if err != nil {
		return nil, err
	}

	if len(res) <= 0 {
		return nil, ErrNoContact
	}

	for _, r := range res {
		con, err := c.GetContact(r.ID)
		if err != nil {
			return nil, err
		}
		if con.HasEmail(query) {
			return con, nil
		}
	}

	return nil, ErrNoContact
	// con, err := c.GetContact(res[0].ID)
	// if err != nil {
	// 	return nil, err
	// }
	// return con, nil
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
