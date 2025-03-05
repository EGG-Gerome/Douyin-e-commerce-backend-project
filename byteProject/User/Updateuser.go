package User

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func UpdateUser(c *gin.Context, DB *gorm.DB) {
	log.WithField("func", "Updateuser")
	var newInfo Userinfo
	var origin Userinfo
	c.ShouldBindJSON(&newInfo)
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "获取id失败"})
		log.Error("获取id失败")
		return
	}
	//看是否能找到原数据
	if err := DB.Where("id=?", id).Find(&origin).Error; err != nil {
		//找不到id
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "用户未注册登录",
		})
		log.Errorf("用户未注册登录:%v", err)
		return
	} else {
		//去修改
		if err := DB.Debug().Model(&origin).Updates(&newInfo).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "更新失败",
			})
			log.Errorf("更新失败:%v", err)
			return
		}
		//打印修改结果
		if err := DB.Where("id=?", id).Find(&newInfo).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "查找数据错误",
			})
			log.Errorf("查找数据错误:%v", err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message":  "success",
			"id":       newInfo.ID,
			"userName": newInfo.UserName,
			"userSex":  newInfo.UserSex,
		})
		log.Infof("用户%v更新数据成功", newInfo.ID)
	}

}
