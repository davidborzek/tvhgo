package tvheadend

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type (
	// ClientOpts defines the configuration options for the client.
	ClientOpts struct {
		URL      string
		Username string
		Password string
	}

	// Query defines a tvheadend request query.
	Query struct {
		m map[string]string
	}

	// Response defines a response from tvheadend.
	Response struct {
		*http.Response
	}

	// Client defines a tvheadend client.
	Client interface {
		// Exec performs a request to the given path
		// and decodes the json-encoded response into the provided interface.
		Exec(ctx context.Context, path string, dst interface{}, query ...Query) (*Response, error)
	}

	client struct {
		opts ClientOpts
		http *http.Client
	}
)

// NewQuery creates a new tvheadend request query.
func NewQuery() Query {
	return Query{
		m: make(map[string]string),
	}
}

// NewWithClient creates a new tvheadend client with a custom http client.
func NewWithClient(opts ClientOpts, httpc *http.Client) Client {
	return &client{
		opts: opts,
		http: httpc,
	}
}

// NewStreamingClient creates a new tvheadend client for streaming
// without http client timeout.
func NewStreamingClient(opts ClientOpts) Client {
	httpc := &http.Client{
		Timeout: 0,
	}
	return NewWithClient(opts, httpc)
}

// New creates a new tvheadend client with a default http client.
func New(opts ClientOpts) Client {
	httpc := &http.Client{
		Timeout: 10 * time.Second,
	}
	return NewWithClient(opts, httpc)
}

func (c *client) Exec(
	ctx context.Context,
	path string,
	dst interface{},
	query ...Query,
) (*Response, error) {
	u, err := url.JoinPath(c.opts.URL, path)
	if err != nil {
		return nil, err
	}

	q := url.Values{}

	if len(query) > 0 {
		for key, value := range query[0].m {
			q.Set(key, value)
		}
	}

	encoded := strings.NewReader(q.Encode())
	req, err := http.NewRequest(http.MethodPost, u, encoded)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if c.opts.Username != "" && c.opts.Password != "" {
		req.SetBasicAuth(c.opts.Username, c.opts.Password)
	}

	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if dst == nil {
		return &Response{res}, nil
	}

	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(dst); err != nil {
		return nil, err
	}

	return &Response{res}, nil
}

// Limit sets the limit parameter.
func (p *Query) Limit(limit int64) {
	p.m["limit"] = strconv.FormatInt(limit, 10)
}

// Start sets the start (offset) parameter.
func (p *Query) Start(offset int64) {
	p.m["start"] = strconv.FormatInt(offset, 10)
}

// SortKey sets the sort key parameter.
func (p *Query) SortKey(key string) {
	p.m["sort"] = key
}

// SortDir sets the sort direction parameter.
func (p *Query) SortDir(dir string) {
	p.m["dir"] = dir
}

// Set sets a parameter with a specific key.
func (p *Query) SetInt(key string, value int64) {
	p.m[key] = strconv.FormatInt(value, 10)
}

// Set sets a parameter with a specific key.
func (p *Query) Set(key string, value string) {
	p.m[key] = value
}

// SetJSON sets a parameter with s specific key
// and encodes the value as json.
func (p *Query) SetJSON(key string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	p.m[key] = string(data)
	return nil
}

// Conf sets the conf parameter and encodes the value as json.
func (p *Query) Conf(v interface{}) error {
	return p.SetJSON("conf", v)
}

// Node sets the node parameter and encodes the value as json.
func (p *Query) Node(v interface{}) error {
	return p.SetJSON("node", v)
}

// Get returns the value of a parameter with a specific key.
func (p *Query) Get(key string) string {
	return p.m[key]
}

func (p *Query) Filter(v []FilterQuery) error {
	return p.SetJSON("filter", v)
}
