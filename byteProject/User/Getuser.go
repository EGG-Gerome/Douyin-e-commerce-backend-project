package User

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetUser(c *gin.Context, DB *gorm.DB) {
	log.WithField("func", "GetUser")
	var userInfo Userinfo
	id, _ := c.Params.Get("id")
	if err := DB.Where("id=?", id).Find(&userInfo).Error; err != nil {
		//找不到id
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "未注册登录",
		})
		log.Errorf("id:%v未注册登录:%v", id, err)
	} else {
		//只返回用户名,性别
		c.JSON(http.StatusOK, gin.H{
			"message":  "success",
			"id":       userInfo.ID,
			"userName": userInfo.UserName,
			"userSex":  userInfo.UserSex,
		})
		log.Infof("获取用户%v信息成功", id)
	}
}
