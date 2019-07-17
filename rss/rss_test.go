package rss

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-05-18

import (
	"testing"
)

func TestGet(t *testing.T) {
	list, err := Get("https://morphs.ru/rss.xml")

	if err != nil {
		t.Fatal("Get return error")
	}

	if list == nil {
		t.Fatal("empty response")
	}

	if len(list) == 0 {
		t.Fatal("empty rss")
	}
}
