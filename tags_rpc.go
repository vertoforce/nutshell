package nutshell

// FindTags ...
func (c *Client) FindTags() (FindTagResult, error) {
	var out FindTagResult
	err := c.rpc.Call("findTags", nil, &out)
	return out, err
}

// NewTag ...
func (c *Client) NewTag(name Tag, tagType TagType) (*Tag, error) {
	var t EntityType
	switch tagType {
	case TagContact:
		t = EntityContact
	case TagCompany:
		t = EntityCompany
	default:
		t = EntityLead
	}

	args := map[string]interface{}{
		"tag": map[string]interface{}{
			"name":       name,
			"entityType": t,
		},
	}

	var out struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}
	err := c.rpc.Call("newTag", args, &out)
	if err != nil {
		return nil, err
	}
	var tg Tag = Tag(out.Name)
	return &tg, nil
}
