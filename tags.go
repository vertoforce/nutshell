package nutshell

import "strings"

type Tag string

type TagsList []Tag

func CreateTagsList(in []string) TagsList {
	out := TagsList{}
	for _, i := range in {
		out = append(out, Tag(i))
	}
	return out
}

// Len ...
func (tl TagsList) Len() int {
	return len(tl)
}

// Swap ...
func (tl TagsList) Swap(i, j int) {
	tl[i], tl[j] = tl[j], tl[i]
}

// Less ...
func (tl TagsList) Less(i, j int) bool {
	a := string(tl[i])
	b := string(tl[j])
	return strings.Compare(a, b) == -1
}

type FindTagResult map[EntityType][]Tag

// HasTag ...
func (ftr FindTagResult) HasTag(name Tag, tagType TagType) bool {
	var t EntityType
	switch tagType {
	case TagContact:
		t = EntityContact
	case TagCompany:
		t = EntityCompany
	case TagLead:
		fallthrough
	default:
		t = EntityLead
	}

	tags, ok := ftr[t]
	if !ok {
		return false
	}

	for _, t := range tags {
		if t == name {
			return true
		}
	}

	return false
}
