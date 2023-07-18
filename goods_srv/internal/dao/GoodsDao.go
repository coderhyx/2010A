package dao

import (
	"context"
	"goods_srv/internal/model"

	"gorm.io/gorm"
)

type GoodsDao struct {
	db *gorm.DB
}

func NewGoodsDao() *GoodsDao {
	return &GoodsDao{
		db: GetDB(),
	}
}

func (m *GoodsDao) FindById(ctx context.Context, id int64) (good *model.Goods, err error) {
	var goods model.Goods
	result := m.db.First(&goods, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &goods, nil
}
