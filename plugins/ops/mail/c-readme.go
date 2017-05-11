package mail

import (
	"net/http"

	"github.com/kapmahc/h2o"
)

func (p *Plugin) getReadme(c *h2o.Context) error {
	return c.TEXT(http.StatusOK, "ops/vpn/readme.md", h2o.H{})
}
