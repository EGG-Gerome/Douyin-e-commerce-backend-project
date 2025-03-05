package User

import (
	"codeup.aliyun.com/codeup/go-micro/Order"
	"github.com/jinzhu/gorm"
)

type Userinfo struct {
	gorm.Model
	UserName string `json:"userName"`
	UserSex  string `json:"userSex"`
	Password string `json:"password"`
	//UserShopList ShopList.Shoplist `json:"userShopList"`
	UserOrder Order.Order `json:"userOrder"`
	Salt      string      `json:"salt"`
}
