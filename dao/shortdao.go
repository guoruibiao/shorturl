package dao

type ShortDao struct{}

func New() (*ShortDao, error) {
	return &ShortDao{}, nil
}

func (dao *ShortDao) SinaURLShort(origin string) (string, error) {

	return "http://t.cn/RxnlTYR", nil
}
