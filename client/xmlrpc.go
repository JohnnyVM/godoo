package godoo

import (
	"errors"
	"net/url"

	"github.com/kolo/xmlrpc"
)

// ClientConfig is the configuration to create a new *Client by givin connection infomations.
type ClientConfig struct {
	Database string
	User     string
	Password string
	Host     string // host or host:port
}

// Client provides high and low level functions to interact with odoo
type Client struct {
	common   *xmlrpc.Client
	object   *xmlrpc.Client
	uid      int64
	password string
	database string
}

func NewClient(conf *ClientConfig) (*Client, error) {
	client := new(Client)
	common := url.URL{
		Scheme: "https",
		Host:   conf.Host,
		Path:   "/xmlrpc/2/common",
	}
	var err error
	client.common, err = xmlrpc.NewClient(common.String(), nil)
	if err != nil {
		return nil, err
	}

	p := make([]any, 4)
	p[0] = conf.Database
	p[1] = conf.User
	p[2] = conf.Password
	p[3] = ""

	var reply any
	err = client.common.Call("authenticate", p, &reply)
	if err != nil {
		return nil, err
	}
	if _, ok := reply.(bool); ok {
		return nil, errors.New("Couldn't authenticate into server: Invalid user/password")
	}
	client.uid = reply.(int64)

	object := url.URL{
		Scheme: "https",
		Host:   conf.Host,
		Path:   "/xmlrpc/2/object",
	}
	client.object, err = xmlrpc.NewClient(object.String(), nil)
	if err != nil {
		return nil, err
	}
	client.password = conf.Password
	client.database = conf.Database
	return client, nil
}

func Close(client *Client) {
	client.common.Close()
	client.object.Close()
	client.uid = 0
	client.password = ""
}

func (client *Client) Close() {
	Close(client)
}

// ExecuteKw base operation
func (c *Client) ExecuteKw(method string, model string, args []any, opt map[string]any) (any, error) {
	var reply any
	err := c.object.Call("execute_kw", []any{c.database, c.uid, c.password, model, method, args, opt}, &reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// Return the list of model ids that match the criteria
func (c *Client) Search(model string, args []any, opt map[string]any) ([]int64, error) {
	reply, err := c.ExecuteKw("search", model, args, opt)
	if err != nil {
		return nil, err
	}
	iarr, ok := reply.([]any)
	if !ok {
		return nil, errors.New("Invalid cast (expected []any)")
	}
	var out []int64 = make([]int64, len(iarr))
	for idx, val := range iarr {
		num, ok := val.(int64)
		if !ok {
			return nil, errors.New("Invalid cast (expected int64)")
		}
		out[idx] = num
	}
	return out, nil
}

func (c *Client) SearchRead(model string, args []any, opt map[string]any) ([]map[string]any, error) {
	reply, err := c.ExecuteKw("search_read", model, args, opt)
	if err != nil {
		return nil, err
	}
	out, ok := reply.([]map[string]any)
	if !ok {
		return nil, errors.New("Invalid cast (expected []map[string]any)")
	}
	return out, nil
}
