package dao

import (
	"testing"
)

func Test_SinaURLShort(t *testing.T) {
	originurl := "http://github.com"
	dao, _ := New()
	shorturl, _ := dao.SinaURLShort(originurl)
	if shorturl == "http://t.cn/RxnlTYR" {
		t.Log("SinaURLShort passed")
	} else {
		t.Error("SinaURLShort failed", originurl, " to ", shorturl)
	}
}
