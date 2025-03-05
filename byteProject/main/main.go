package main

import (
	"codeup.aliyun.com/codeup/go-micro/Goods"
	"codeup.aliyun.com/codeup/go-micro/Logger"
	"codeup.aliyun.com/codeup/go-micro/Order"
	"codeup.aliyun.com/codeup/go-micro/ShopList"
	"codeup.aliyun.com/codeup/go-micro/User"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	DB *gorm.DB
)

func main() {
	InitServer()
	defer DB.Close()
	r := gin.Default()
	//用户的增删查改
	//TODO casbin用户登录验证
	//用户的增删改查
	userG := r.Group("user")
	{
		userG.GET("/:id", GetUser)
		userG.POST("/", AddUser)
		userG.PUT("/:id", UpdateUser)
		userG.DELETE("/:id", DeleteUser)
	}
	//商品的增加，获取商品信息
	goodsG := r.Group("goods")
	{
		goodsG.GET("/:id", GetOneGoods)
		goodsG.GET("/", GetSomeGoods)
		goodsG.POST("/", AddGoods)
	}
	//用户购物清单增删改查
	shoplistG := r.Group("list")
	{
		shoplistG.GET("/:userID", GetShoplist)
		shoplistG.POST("/", AddGoodsToList)
		shoplistG.PUT("/:userID", RemoveGoodsfromList)
		shoplistG.DELETE("/:userID", ClearShoplist)
	}
	//创建订单,查看订单
	orderG := userG.Group("order")
	{
		orderG.POST("/", CreateOrder)
		//orderG.GET("/:userID", GetOrder)
	}
	r.Run(":9090")
}

func InitServer() {
	if err := Logger.InitLogger(); err != nil {
		log.Error("日志记录器初始化失败")
		return
	} else {
		log.Info("日志记录器初始化成功")
	}
	err := InitMysql()
	if err != nil {
		log.Error("数据库连接失败")
		return
	}
	log.Info("数据库连接成功")
}
func InitMysql() (e error) {
	dsn := "test:123456@tcp(127.0.0.1:3306)/ByteShop?charset=utf8mb4&parseTime=True&loc=Local"
	//DB是全局变量，所以此处不需要冒号
	//如果创建没问题就看能不能ping通
	DB, e = gorm.Open("mysql", dsn) //上面的返回已经声明了返回变量
	if e != nil {
		//创建失败就返回e
		return e
	}
	DB.AutoMigrate(&User.Userinfo{})
	DB.AutoMigrate(&Goods.Goodsinfo{})
	DB.AutoMigrate(&ShopList.Relation{})
	DB.AutoMigrate(&Order.Order{})
	return DB.DB().Ping()
}
func GetUser(c *gin.Context) {
	User.GetUser(c, DB)
}
func AddUser(c *gin.Context) {
	fmt.Println("add user")
	User.AddUser(c, DB)
}
func UpdateUser(c *gin.Context) {
	User.UpdateUser(c, DB)
}
func DeleteUser(c *gin.Context) {
	User.DeleteUser(c, DB)
}

func GetShoplist(c *gin.Context) {
	ShopList.GetList(c, DB)
}
func AddGoodsToList(c *gin.Context) {
	ShopList.AddGoodsToList(c, DB)
}
func RemoveGoodsfromList(c *gin.Context) {
	ShopList.RemoveGoods(c, DB)
}
func ClearShoplist(c *gin.Context) {
	ShopList.ClearShopList(c, DB)
}

func AddGoods(c *gin.Context) {
	log.Info("增加商品")
	Goods.AddGoods(c, DB)
}
func GetOneGoods(c *gin.Context) {
	log.Info("获得一个商品")
	Goods.GetOneGoods(c, DB)
}
func GetSomeGoods(c *gin.Context) {
	log.Info("批量获得商品")
	Goods.GetSomeGoods(c, DB)
}

func CreateOrder(c *gin.Context) {
	ShopList.CreateOrder(c, DB)
}
func GetOrder(c *gin.Context) {
	Order.GetOrder(c, DB)
}
