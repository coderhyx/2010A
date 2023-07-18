package model

import "time"

type Goods struct {
	Id              int32     `gorm:"column:id;type:INT(11);AUTO_INCREMENT;NOT NULL"`
	AddTime         time.Time `gorm:"column:add_time;type:DATETIME(3);"`
	UpdateTime      time.Time `gorm:"column:update_time;type:DATETIME(3);"`
	DeletedAt       time.Time `gorm:"column:deleted_at;type:DATETIME(3);"`
	IsDeleted       int8      `gorm:"column:is_deleted;type:TINYINT(1);"`
	CategoryId      int32     `gorm:"column:category_id;type:INT(11);NOT NULL"`
	BrandsId        int32     `gorm:"column:brands_id;type:INT(11);NOT NULL"`
	OnSale          int8      `gorm:"column:on_sale;type:TINYINT(1);NOT NULL"`
	ShipFree        int8      `gorm:"column:ship_free;type:TINYINT(1);NOT NULL"`
	IsNew           int8      `gorm:"column:is_new;type:TINYINT(1);NOT NULL"`
	IsHot           int8      `gorm:"column:is_hot;type:TINYINT(1);NOT NULL"`
	Name            string    `gorm:"column:name;type:VARCHAR(50);NOT NULL"`
	GoodsSn         string    `gorm:"column:goods_sn;type:VARCHAR(50);NOT NULL"`
	ClickNum        int32     `gorm:"column:click_num;type:INT(11);NOT NULL"`
	SoldNum         int32     `gorm:"column:sold_num;type:INT(11);NOT NULL"`
	FavNum          int32     `gorm:"column:fav_num;type:INT(11);NOT NULL"`
	MarketPrice     float64   `gorm:"column:market_price;type:FLOAT;NOT NULL"`
	ShopPrice       float64   `gorm:"column:shop_price;type:FLOAT;NOT NULL"`
	GoodsBrief      string    `gorm:"column:goods_brief;type:VARCHAR(100);NOT NULL"`
	GoodsFrontImage string    `gorm:"column:goods_front_image;type:VARCHAR(200);NOT NULL"`
}

func (*Goods) TableName() string {
	return "goods"
}
