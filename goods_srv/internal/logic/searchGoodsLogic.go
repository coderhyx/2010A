package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"goods_srv/internal/svc"
	"goods_srv/pb/goods"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type SearchGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type SearchGood struct {
	Shards struct {
		Total      int `json:"total" gorm:"column:total"`
		Failed     int `json:"failed" gorm:"column:failed"`
		Successful int `json:"successful" gorm:"column:successful"`
		Skipped    int `json:"skipped" gorm:"column:skipped"`
	} `json:"_shards" gorm:"column:_shards"`
	Hits struct {
		Hits []struct {
			Index  string `json:"_index" gorm:"column:_index"`
			Type   string `json:"_type" gorm:"column:_type"`
			Source struct {
				AlipayNickname string      `json:"alipayNickname" gorm:"column:alipayNickname"`
				Birthday       interface{} `json:"birthday" gorm:"column:birthday"`
				Note           string      `json:"note" gorm:"column:note"`
				VaildPeroid    string      `json:"vaildPeroid" gorm:"column:vaildPeroid"`
				Sex            int         `json:"sex" gorm:"column:sex"`
				Batch          interface{} `json:"batch" gorm:"column:batch"`
				Mobile         string      `json:"mobile" gorm:"column:mobile"`
				Avatar         string      `json:"avatar" gorm:"column:avatar"`
				Privilege      int         `json:"privilege" gorm:"column:privilege"`
				CreatedAt      string      `json:"createdAt" gorm:"column:createdAt"`
				DeletedAt      interface{} `json:"deletedAt" gorm:"column:deletedAt"`
				Password       string      `json:"password" gorm:"column:password"`
				Constellation  string      `json:"constellation" gorm:"column:constellation"`
				Nickname       string      `json:"nickname" gorm:"column:nickname"`
				Location       string      `json:"location" gorm:"column:location"`
				ID             int         `json:"id" gorm:"column:id"`
				Email          interface{} `json:"email" gorm:"column:email"`
				UpdatedAt      string      `json:"updatedAt" gorm:"column:updatedAt"`
			} `json:"_source" gorm:"column:_source"`
			ID    string  `json:"_id" gorm:"column:_id"`
			Score float64 `json:"_score" gorm:"column:_score"`
		} `json:"hits" gorm:"column:hits"`
		Total struct {
			Value    int    `json:"value" gorm:"column:value"`
			Relation string `json:"relation" gorm:"column:relation"`
		} `json:"total" gorm:"column:total"`
		MaxScore float64 `json:"max_score" gorm:"column:max_score"`
	} `json:"hits" gorm:"column:hits"`
	Took     int  `json:"took" gorm:"column:took"`
	TimedOut bool `json:"timed_out" gorm:"column:timed_out"`
}

func NewSearchGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchGoodsLogic {
	return &SearchGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type Gen struct {
	Hits struct {
		Hits []struct {
			Source struct {
				AlipayNickname string      `json:"alipayNickname" gorm:"column:alipayNickname"`
				Birthday       interface{} `json:"birthday" gorm:"column:birthday"`
				Note           string      `json:"note" gorm:"column:note"`
				VaildPeroid    string      `json:"vaildPeroid" gorm:"column:vaildPeroid"`
				Sex            int         `json:"sex" gorm:"column:sex"`
				Batch          interface{} `json:"batch" gorm:"column:batch"`
				Mobile         string      `json:"mobile" gorm:"column:mobile"`
				Avatar         string      `json:"avatar" gorm:"column:avatar"`
				Privilege      int         `json:"privilege" gorm:"column:privilege"`
				CreatedAt      string      `json:"createdAt" gorm:"column:createdAt"`
				DeletedAt      interface{} `json:"deletedAt" gorm:"column:deletedAt"`
				Password       string      `json:"password" gorm:"column:password"`
				Constellation  string      `json:"constellation" gorm:"column:constellation"`
				Nickname       string      `json:"nickname" gorm:"column:nickname"`
				Location       string      `json:"location" gorm:"column:location"`
				ID             int         `json:"id" gorm:"column:id"`
				Email          interface{} `json:"email" gorm:"column:email"`
				UpdatedAt      string      `json:"updatedAt" gorm:"column:updatedAt"`
			} `json:"_source" gorm:"column:_source"`
		} `json:"hits" gorm:"column:hits"`
	} `json:"hits" gorm:"column:hits"`
}

type HitsResp struct {
	Hits Hits `json:"hits"`
}
type Hits struct {
	Hits []Hit `json:"hits"`
}
type Hit struct {
	Source User `json:"_source"`
}
type User struct {
	AlipayNickname string `json:"alipayNickname"`
	Avatar         string `json:"avatar"`
	Batch          string `json:"batch"`
	Birthday       string `json:"birthday"`
	Constellation  string `json:"constellation"`
	CreatedAt      string `json:"createdAt"`
	DeletedAt      string `json:"deletedAt"`
	Email          string `json:"email"`
	ID             int    `json:"id"`
	Location       string `json:"location"`
	Mobile         string `json:"mobile"`
	Nickname       string `json:"nickname"`
	Note           string `json:"note"`
	Password       string `json:"password"`
	Privilege      int    `json:"privilege"`
	Sex            int    `json:"sex"`
	UpdatedAt      string `json:"updatedAt"`
	VaildPeroid    string `json:"vaildPeroid"`
}

func (l *SearchGoodsLogic) SearchGoods(in *goods.SearchGoodsReq) (*goods.SearchGoodsResp, error) {
	//从kibana 可视化界面copy
	query := `{
	   "query": {
	       "bool": {
	           "must": [
	               {
	                   "match_phrase": {
	                       "nickname": {
	                           "query": "prprpr"
	                       }
	                   }
	               }
	           ]
	       }
	   }
	}`
	result, err := l.svcCtx.Es.SearchDocuments("user", query)
	var hitsResp HitsResp
	//通过结构体解析
	err = json.Unmarshal(result, &hitsResp)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}
	//装载数据
	if len(hitsResp.Hits.Hits) > 0 {
		var user User
		var users []User
		for i, _ := range hitsResp.Hits.Hits {
			err = copier.Copy(&user, &hitsResp.Hits.Hits[i].Source)
			if err != nil {
				fmt.Println("Error copying values:", err)
				return nil, err
			}
			users = append(users, user)
		}
		fmt.Println("------------------>User:", users)
	} else {
		fmt.Println("No hits found.")
	}
	return &goods.SearchGoodsResp{}, nil
}
