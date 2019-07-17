package uagent

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-07-17

import (
	"bytes"
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
)

type Client struct {
	Timeout            time.Duration
	UserAgent          string
	NoDecode           bool
	DisableKeepAlives  bool
	InsecureSkipVerify bool
}

func New() *Client {
	return &Client{
		UserAgent:          "uagent",
		DisableKeepAlives:  true,
		InsecureSkipVerify: true,
		Timeout:            time.Duration(5 * time.Second),
	}
}

type Response struct {
	StatusCode int
	Content    []byte
	Header     http.Header
}

func (c *Client) Request(method string, url string, headers map[string]string, content []byte) (*Response, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.InsecureSkipVerify},
	}

	ua := &http.Client{
		Timeout:   c.Timeout,
		Transport: tr,
	}

	tr.DisableKeepAlives = c.DisableKeepAlives

	var rd io.Reader

	if content != nil && len(content) > 0 && (method == "POST" || method == "PUT") {
		rd = bytes.NewReader(content)
	}

	req, err := http.NewRequest(method, url, rd)

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	resp, err := ua.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var text []byte

	if c.NoDecode {
		text, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
	} else {
		utf8, err1 := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
		if err1 != nil {
			return nil, err1
		}

		text, err = ioutil.ReadAll(utf8)
		if err != nil {
			return nil, err
		}
	}

	r := &Response{
		StatusCode: resp.StatusCode,
		Content:    text,
		Header:     resp.Header,
	}

	return r, nil
}
