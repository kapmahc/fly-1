package root

import "github.com/astaxie/beego/toolbox"

// DatabaseCheck check database
type DatabaseCheck struct {
}

// Check check
func (p *DatabaseCheck) Check() error {
	// TODO
	return nil
}

// CacheCheck check cache
type CacheCheck struct {
}

// Check check
func (p *CacheCheck) Check() error {
	// TODO
	return nil
}

func init() {
	toolbox.AddHealthCheck("database", &DatabaseCheck{})
	toolbox.AddHealthCheck("cache", &CacheCheck{})
}
