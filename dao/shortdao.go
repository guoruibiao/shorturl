package dao

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"shorturl/model"
	// model "github.com/guoruibiao/shorturl/model"
)

type ShortDao struct{}

func New() (*ShortDao, error) {
	return &ShortDao{}, nil
}

func (dao *ShortDao) SinaURLShort(origin string) (*model.Response, error) {
	// TODO origin可能为非法字符串，需要考虑下校验
	resp, err := http.Get("https://api.weibo.com/2/short_url/shorten.json?source=2257828842&url_long=" + origin)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil
	}
	var response model.Response
	json.Unmarshal(body, &response)
	return &response, nil
}
