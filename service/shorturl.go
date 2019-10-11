package service

import (
	// dao "github.com/guoruibiao/shorturl/dao"
	// model "github.com/guoruibiao/shorturl/model"
	"github.com/guoruibiao/shorturl/model"
	"github.com/guoruibiao/shorturl/dao"
)

type ShortURLService struct {
	shortUrlDao dao.ShortDao
}

func New() (ShortURLService, error) {
	dao, err := dao.New()
	var service ShortURLService = ShortURLService{
		shortUrlDao: *dao,
	}
	if err != nil {
		return service, err
	}
	return service, nil
}

func (service ShortURLService) ShortURL(encodedurl string) (response *model.Response, err error) {
	response, err = service.shortUrlDao.SouGouURLShort(encodedurl)
	if err != nil {
		response, err = service.shortUrlDao.Api985URLShort(encodedurl)
		if err != nil {
			response, err = service.shortUrlDao.ChkajaURLShort(encodedurl)
		}
	}
	return
}
