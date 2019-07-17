package uagent

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-07-17

import (
	"testing"
)

func TestUserAgent(t *testing.T) {

	ua := New()

	resp, err := ua.Request("GET", "https://morphs.ru", nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatal("wrong status code")
	}

	if string(resp.Content) == "" {
		t.Fatal("wrong respose")
	}
}
