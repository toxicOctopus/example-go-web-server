package page

// Content content structure
type Content struct {
	content string
}

// Add adds content
func (c *Content) Add(s string) {
	c.content += s
}

// Get gets content
func (c *Content) Get() string {
	return c.content
}
