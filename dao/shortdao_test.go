package dao

import (
	"testing"
)

func Test_SinaURLShort(t *testing.T) {
	originurl := "https://github.com"
	dao, _ := New()
	response, _ := dao.SinaURLShort(originurl)
	if response.URLS[0].ShortURL == "http://t.cn/RxnlTYR" {
		t.Log("SinaURLShort passed")
	} else {
		t.Error("SinaURLShort failed", originurl, " to ", response.URLS[0].ShortURL)
	}
}
