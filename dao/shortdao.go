package dao

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	// model "github.com/guoruibiao/shorturl/model"
	"github.com/guoruibiao/shorturl/config"
	"github.com/guoruibiao/shorturl/model"
	"github.com/pkg/errors"
	)

type ShortDao struct{}

func New() (*ShortDao, error) {
	return &ShortDao{}, nil
}

func (dao *ShortDao) Api985URLShort(encodedurl string) (*model.Response, error) {
    // encodedurl 应该是已经被urlencode之后的数据
    apiLink := config.API_985_SO + encodedurl
    resp, err := http.Get(apiLink)
    if err != nil {
    	return nil, err
	}
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
    	return nil, err
	}
    type TmpResponse struct {
    	Url string `json:"url"`
    	Error string `json:"error"`
	}
    var tmpResponse = &TmpResponse{}
    json.Unmarshal(body, tmpResponse)
    if tmpResponse.Error == "" {
    	return &model.Response{
    		Result: tmpResponse.Url,
    		Error: nil,
		}, nil
	}
    return nil, errors.New("985.so cannot short this url.")
}

func (dao *ShortDao) ChkajaURLShort(encodedurl string) (*model.Response, error) {
	apiLink := config.API_CHKAJA_COM + encodedurl
	resp, err := http.Get(apiLink)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &model.Response{
		Result: string(body),
		Error: nil,
	},nil
}


func (dao *ShortDao) SouGouURLShort(encodedurl string) (*model.Response, error) {
	apiLink := config.API_SOGOU_COM + encodedurl
	resp, err := http.Get(apiLink)
	if err != nil {
		return nil ,err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &model.Response{
		Result: string(body),
		Error: nil,
	},nil
}