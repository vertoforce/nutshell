package nutshell

import (
	"fmt"
	"net/http"

	"github.com/seanhagen/rpc-codec/jsonrpc2"
)

type Client struct {
	rpc *jsonrpc2.Client
}

const initialEndpoint = "https://api.nutshell.com/v1/json"

// getApiEndpoint ...
func getApiEndpoint(user string) (string, error) {
	rpc := jsonrpc2.NewCustomHTTPClient(initialEndpoint, nil)
	out := map[string]string{}
	err := rpc.Call("getApiForUsername", map[string]string{"username": user}, &out)
	if err != nil {
		return "", err
	}
	end, ok := out["api"]
	if !ok {
		return "", fmt.Errorf("no api endpoint found for user '%v'", user)
	}

	end = fmt.Sprintf("https://%v/api/v1/json", end)
	return end, nil
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
			key:  key,
			orig: http.DefaultTransport,
		},
	}

	rpc := jsonrpc2.NewCustomHTTPClient(ep, cl)
	return &Client{rpc}, nil
}
