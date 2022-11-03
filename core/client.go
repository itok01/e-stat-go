package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	querystring "github.com/google/go-querystring/query"
)

const (
	ApiBaseURL = "http://api.e-stat.go.jp/rest/3.0/app"
)

type HttpClient struct {
	httpClient http.Client
	debug      bool
}

type IHttpClient interface {
	Get(ctx context.Context, path string, query any) (int, []byte, error)
	Post(ctx context.Context, path string, data any) (int, []byte, error)
	PostJsonWithQuery(ctx context.Context, path string, query any, structuredData any) (int, []byte, error)
}

func NewClient(debug bool) IHttpClient {
	return &HttpClient{
		httpClient: *http.DefaultClient,
		debug:      debug,
	}
}

func (c *HttpClient) doRequest(req *http.Request) (int, []byte, error) {
	if c.debug {
		log.Printf("%s %s", req.Method, req.URL)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, nil, err
	}
	log.Print(resp.Status)

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode, b, nil
}

func (c *HttpClient) request(ctx context.Context, method string, path string, structuredData any) (int, []byte, error) {
	targetURL := urlWithXmlFormatFromPath(path)

	data, err := querystring.Values(structuredData)
	if err != nil {
		return 0, nil, err
	}

	var req *http.Request
	if method == http.MethodGet {
		req, err = http.NewRequestWithContext(ctx, method, targetURL, nil)
		req.URL.RawQuery = data.Encode()
		log.Print(req.URL.RawQuery)
	} else if method == http.MethodPost {
		req, err = http.NewRequestWithContext(ctx, method, targetURL, strings.NewReader(data.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		log.Print(data.Encode())
	}
	if err != nil {
		return 0, nil, err
	}

	return c.doRequest(req)
}

func (c *HttpClient) Get(ctx context.Context, path string, query any) (int, []byte, error) {
	return c.request(ctx, http.MethodGet, path, query)
}

func (c *HttpClient) Post(ctx context.Context, path string, data any) (int, []byte, error) {
	return c.request(ctx, http.MethodPost, path, data)
}

func (c *HttpClient) PostJsonWithQuery(ctx context.Context, path string, query any, structuredData any) (int, []byte, error) {
	targetURL := urlWithXmlFormatFromPath(path)

	queryData, err := querystring.Values(query)
	if err != nil {
		return 0, nil, err
	}

	data, err := json.Marshal(structuredData)
	if err != nil {
		return 0, nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, targetURL, bytes.NewBuffer(data))
	req.Header.Add("Content-Type", "application/json")
	req.URL.RawQuery = queryData.Encode()
	log.Print(string(data))
	if err != nil {
		return 0, nil, err
	}

	return c.doRequest(req)
}

func urlWithXmlFormatFromPath(path string) string {
	return fmt.Sprintf("%s%s", ApiBaseURL, path)
}
