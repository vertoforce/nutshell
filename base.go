package nutshell

import (
	"fmt"
	"net/http"

	"github.com/seanhage/rpc-codec/jsonrpc2"
)

type Client struct {
	rpc *jsonrpc2.Client
}

type rt struct {
	user string
	key  string
	orig http.RoundTripper
}

// RoundTrip ...
func (rt rt) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(rt.user, rt.key)
	return rt.orig.RoundTrip(req)
}

const initialEndpoint = "https://api.nutshell.com/v1/json"

// getApiEndpoint ...
func getApiEndpoint(user string) (string, error) {
	rpc := jsonrpc2.NewCustomHTTPClient(initialEndpoint)
	out := map[string]string{}
	err := rpc.Call("getApiForUsername", map[string]string{"username": user}, &out)
	if err != nil {
		return "", err
	}
	end, ok := out["api"]
	if ok {
		return end, nil
	}
	return "", fmt.Errorf("no api endpoint found for user '%v'", user)
}

// New ...
func New(user, key string) (*Client, error) {
	ep, err := getApiEndpoint(user)
	if err != nil {
		return nil, err
	}
	cl := &http.Client{
		Transport: rt{
			user: user,
			key:  pass,
			orig: http.DefaultTransport,
		},
	}

	rpc := jsonrpc2.NewCustomHTTPClient(ep, cl)
	return &Client{rpc}, nil
}
