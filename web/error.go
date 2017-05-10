package web

// HTTPError http error
type HTTPError struct {
	Message string
	Code    int
}

func (p *HTTPError) Error() string {
	return p.Message
}
