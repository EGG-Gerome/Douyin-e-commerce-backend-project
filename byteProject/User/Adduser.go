package User

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"math/big"
	"net/http"
)

func generateSalt(length int) (string, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[num.Int64()]
	}
	return string(result), nil
}

// hashPassword 使用 SHA-256 进行加盐哈希
func hashPassword(password, salt string) string {
	hasher := sha256.New()
	hasher.Write([]byte(salt))
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}
func AddUser(c *gin.Context, DB *gorm.DB) {
	log.WithField("func", "AddUser")
	var user Userinfo
	var password string
	//获取用户名和密码
	c.ShouldBindJSON(&user)
	if user.UserName == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "没有输入用户名或密码",
		})
		log.Error("没有输入用户名或密码")
		return
	}
	//使用哈希加盐散列来加密存储用户密码
	salt, _ := generateSalt(10)
	user.Password = hashPassword(password, salt)
	user.Salt = salt
	//获取性别
	if user.UserSex == "" {
		user.UserSex = "Unknown"
	}
	//存储
	result := DB.Create(&user)
	if result.Error != nil {
		// 处理错误
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "wrong",
		})
		fmt.Println("Error:", result.Error)
		log.Errorf("error:%v", result.Error)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "success",
		"id":       user.ID,
		"userName": user.UserName,
		"userSex":  user.UserSex,
	})
	log.Infof("成功注册用户,ID:%d", user.ID)
}
