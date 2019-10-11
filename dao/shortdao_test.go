package dao

import (
	"testing"
	"net/url"
)

func TestShortDao_Api985URLShort(t *testing.T) {
	longurl := url.QueryEscape("https://github.com")
	dao, _ := New()
	resp, err := dao.Api985URLShort(longurl)
	if err!=nil {
		t.Error(err)
	}
	t.Log(resp.Result)
}