package dao

import (
	"context"
	"goods_srv/internal/model"

	"gorm.io/gorm"
)

type GoodsDao struct {
	db *gorm.DB
}

func NewGoodsDao(db *gorm.DB) *GoodsDao {
	return &GoodsDao{
		db: db,
	}
}

func (m *GoodsDao) FindById(ctx context.Context, id int64) (good *model.Goods, err error) {
	var goods model.Goods
	//方式一:
	d1 := GetDB()
	d1.Exec("select 1")
	//方式二:
	m.db.Exec("select 1")
	//result := m.db.First(&goods, id)
	//if result.Error != nil {
	//	return nil, result.Error
	//}
	return &goods, nil
}
