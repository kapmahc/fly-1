package widgets

// Dropdown dropdown
type Dropdown struct {
	Label string
	Items []*Link
}

// NewDropdown new dropdown
func NewDropdown(label string, links ...*Link) *Dropdown {
	if links == nil {
		links = make([]*Link, 0)
	}
	return &Dropdown{Label: label, Items: links}
}

// Link link
type Link struct {
	Label string
	Href  string
}

// NewLink new link
func NewLink(label, href string) *Link {
	return &Link{Label: label, Href: href}
}
