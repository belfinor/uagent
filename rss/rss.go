package rss

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-07-17

import (
	"bytes"
	"encoding/xml"
	"errors"
	"time"

	"github.com/belfinor/uagent"
	"golang.org/x/net/html/charset"
)

type Rss struct {
	Channel RssChannel `xml:"channel"`
}

type RssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

type RssChannel struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Items       []RssItem `xml:"item"`
}

func Get(url string) ([]RssItem, error) {
	ua := uagent.New()

	ua.UserAgent = "Mozilla/5.0 (Windows NT 6.1) Gecko/20100101 Thunderbird/52.7.0 Lightning/5.4"
	ua.Timeout = time.Second * 15

	resp, err := ua.Request("GET", url, nil, nil)

	if err != nil {
		return nil, err
	}

	if resp == nil || resp.Content == nil {
		return nil, errors.New("empty response")
	}

	buffer := bytes.NewBuffer([]byte(resp.Content))
	xml := xml.NewDecoder(buffer)

	xml.CharsetReader = charset.NewReaderLabel

	rss := new(Rss)

	if err = xml.Decode(rss); err != nil {
		return nil, errors.New("xml decode failed")
	}

	return rss.Channel.Items, nil
}
