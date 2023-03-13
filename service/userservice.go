package service

import (
	"fmt"
	"github/IM_gin+vue/models"
	"github/IM_gin+vue/utils"
	"math/rand"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary 所有用户
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := models.GetUserList()
	c.JSON(200, gin.H{
		"code":    0,
		"message": "查询成功",
		"data":    data,
	})

}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")

	//加入随机数
	r := rand.New(rand.NewSource(time.Now().UnixNano())) //以纳秒级时间戳作为种子
	random := r.Int31()
	fmt.Println("随机数random=", random)
	salt := fmt.Sprintf("%06d", random)

	data := models.FindUserByName(user.Name)
	if user.Name == "" || password == "" || repassword == "" {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": "用户名或密码不能为空！！！",
			"data":    user,
		})
		return
	}
	if data.Name != "" {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": "用户已注册",
			"data":    user,
		})
		return
	}
	if password != repassword {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": "两次密码输入不一致！！！",
			"data":    user,
		})
		return
	}
	user.Salt = salt
	user.PassWord = utils.MakePassword(password, salt)
	// user.PassWord = password
	fmt.Println("新建用户的用户名为:", user.Name, "密码为：", user.PassWord)
	user.LoginTime = time.Now()
	user.LogOutTime = time.Now()
	user.HeartbeatTime = time.Now()
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "新增用户成功！",
		"data":    user,
	})
}

// FindUserByNameAndPwd
// @Summary 使用用户名密码查询用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/findUserByNameAndPwd [post]
func FindUserByNameAndPwd(c *gin.Context) {

}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "id"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"message": "删除用户成功！！！",
	})
}

// UpdateUser
// @Summary 更新用户
// @Tags 用户模块
// @param id query string false "id"
// @param name query string false "name"
// @param password query string false "password"
// @param phone query string false "phone"
// @param email query string false "email"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	fmt.Println("update :", user)

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"message": "修改参数不匹配！！",
		})
	} else {
		models.UpdateUser(user)
		c.JSON(200, gin.H{
			"message": "修改用户成功！",
		})
	}

	models.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "新增用户成功！",
	})
}
