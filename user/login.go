package user

//登录模块
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"web/config"
	"web/testpackage"
)

type uinfo struct {
	username string
	passwd   string
}

func Login(c *gin.Context) { //登录模块
	ipadd := c.ClientIP()              //获取ip地址
	username := c.PostForm("username") //获取用户名
	password := c.PostForm("password") //获取密码
	captcha := c.PostForm("captcha")   //获取验证码
	err1 := testpackage.InitDB()       // 调用输出化数据库的函数
	if err1 != nil {
		fmt.Printf("init db failed,err:%v\n", err1)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "发生甚么事了？数据库连接失败！",
		})
		return
	}
	var u uinfo
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "无效的参数",
		})
		return
	}
	if username == "" || password == "" || captcha == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "请填写各个输入框内容，谢谢~",
		})
		return
	}
	if testpackage.SQLInject(username) == true || testpackage.SQLInject(password) == true {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "别来注入了亲爱的。",
		})
		return
	}
	if testpackage.SQLInject(ipadd) == true {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "你x个锤子呢,你觉得能不能x嘛。"})
		return
	}
	if testpackage.CaptchaVerify(c, captcha) { //验证验证码是否正确
		if testpackage.UserLogin(username, password) == true {
			if testpackage.Authfilenub(username) <= "1" {
				testpackage.Lastlogin(username) //记录上次登录时间
				c.SetCookie("FileAuth", "", 3600, "/", config.HostName, false, true)
				tokenString, _ := testpackage.GenToken(username)
				testpackage.LoginLog(username, ipadd)
				fmt.Println(ipadd)
				c.SetCookie("Authorization", tokenString, 3600, "/", config.HostName, false, true)
				if username == config.AdminUserName {
					c.JSON(http.StatusOK, gin.H{
						"static":  200,
						"message": "登录成功！亲爱的管理员",
						"url":     "/api/admin/getinfo",
						"data":    gin.H{"token": tokenString},
					})
					return
				} else {
					c.JSON(http.StatusOK, gin.H{
						"static":  200,
						"message": "登录成功！,欢迎你:" + username,
						"url":     "/center/user",
						"data":    gin.H{"token": tokenString},
					})
					return
				}
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  500,
				"message": "登录失败！",
			})
			return
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "验证码错误!",
		})
		return
	}
}
