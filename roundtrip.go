package nutshell

import "net/http"

type rt struct {
	user string
	key  string
	orig http.RoundTripper
}

// RoundTrip ...
func (rt rt) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(rt.user, rt.key)
	req.Header.Set("Content-Type", "application/json")
	return rt.orig.RoundTrip(req)
}
