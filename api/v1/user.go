package v1

import (
	"gin-project/model"
	"gin-project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Testtype struct {
	Name string      `json:"name"`
	Tags []model.Tag `json:"tags"`
}

// Test doc
// @Description  Test
// @Accept       json
// @Produce      json
// @Tags         User
// @Param        data  body      Testtype  true  "用户名，密码"
// @Success      200   {string}  string    "{"status": true, "message": "注册成功"}"
// @Router       /user/test [post]
func Test(c *gin.Context) {
	var data Testtype
	if err := c.ShouldBindJSON(&data); err != nil {
		panic(err)
	}
	println(data.Name)
	println(data.Tags)
	println(data.Tags[0].Name)
	println(data.Tags[1].Name)
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": data.Name,
	})
}

// Register doc
// @Description  Register
// @Tags         User
// @Param        username   formData  string  true  "username"
// @Param        password1  formData  string  true  "password1"
// @Param        password2  formData  string  true  "password2"
// @Param        email      formData  string  true  "email"
// @Success      200        {string}  string  "{"status": true, "message": "注册成功"}"
// @Router       /user/register [post]
func Register(c *gin.Context) {
	username := c.Request.FormValue("username")
	password1 := c.Request.FormValue("password1")
	password2 := c.Request.FormValue("password2")
	email := c.Request.FormValue("email")
	if password1 != password2 {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "两次输入的密码不一致",
		})
		return
	}
	_, notFoundUsername := service.QueryUserByUsername(username)
	_, notFoundEmail := service.QueryUserByEmail(email)
	if notFoundUsername && notFoundEmail {
		user := model.User{
			Username: username,
			Password: password1,
			Email:    email,
		}
		service.CreateUser(&user)
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "注册成功",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  false,
		"message": "用户名/邮箱已存在",
	})
}

// Login doc
// @Description  Login
// @Tags         User
// @Param        username  formData  string  true  "username"
// @Param        password  formData  string  true  "password"
// @Success      200       {string}  string  "{"status": true, "message": "登录成功","data": user}"
// @Router       /user/login [post]
func Login(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	user, notFound := service.QueryUserByUsername(username)
	if notFound {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "用户名不存在",
		})
		return
	}
	if user.Password != password {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "密码错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "登录成功",
		"data":    user,
	})
}

// ShowUserInfo doc
// @Description  ShowUserInfo
// @Tags         User
// @Param        user_id  formData  string  true  "user_id"
// @Success      200      {string}  string  "{"status": true, "message": "查询成功", "data": user}"
// @Router       /user/info [post]
func ShowUserInfo(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Request.FormValue("user_id"), 0, 64)
	user, notFound := service.QueryUserByUserID(userID)
	if notFound {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": "用户名不存在",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "查询成功",
		"data":    user,
	})
}