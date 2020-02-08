package nutshell

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type rt struct {
	user string
	key  string
	orig http.RoundTripper
}

// RoundTrip ...
func (rt rt) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(rt.user, rt.key)
	req.Header.Set("Content-Type", "application/json")

	res, err := rt.orig.RoundTrip(req)
	if err != nil {
		return res, err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	res.Body.Close()

	var out interface{}
	err = json.Unmarshal(b, &out)
	if err != nil {
		return nil, err
	}

	// got here without error, so the body is probably fine, set the status code to 200
	res.StatusCode = http.StatusOK
	res.Body = ioutil.NopCloser(bytes.NewBuffer(b))

	return res, nil
}
