package service

import (
	// "shorturl/dao"
	// "shorturl/model"
	dao "github.com/guoruibiao/shorturl/dao"
	model "github.com/guoruibiao/shorturl/model"
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

func (service ShortURLService) ShortURL(origin string) (response *model.Response, err error) {
	response, err = service.shortUrlDao.SinaURLShort(origin)
	return
}
