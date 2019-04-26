package service

import (
	"testing"
)

func Test_ShortURL(t *testing.T) {
	service, err := New()
	if err != nil {
		t.Error("service New failed")
	}
	origin := "https://github.com"
	resp, err := service.ShortURL(origin)
	if err != nil {
		t.Error("service ShortURL failed")
	}
	if resp.URLS[0].ShortURL != "http://t.cn/RxnlTYR" {
		t.Error("service shorturl result failed")
	}
}
