package User

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func DeleteUser(c *gin.Context, DB *gorm.DB) {
	log.WithField("func", "Deleteuser")
	id, _ := c.Params.Get("id")
	var origin Userinfo
	//看是否能找到原数据
	if err := DB.Where("id=?", id).Find(&origin).Error; err != nil {
		//找不到id
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "用户未注册登录",
		})
		log.Errorf("用户未注册登录:%v", err)
	} else {
		//执行删除
		if err := DB.Where("id=?", id).Delete(&Userinfo{}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "删除失败",
			})
			log.Errorf("删除失败:%v，用户id:%v", err, id)

			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
		log.Infof("删除用户%v成功", id)

	}
}
