package user

//注册功能
import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"web/testpackage"
)

func UserRag(c *gin.Context) {
	ipadd := c.ClientIP()              //获取ip地址
	username := c.PostForm("username") //获取用户名
	password := c.PostForm("password") //获取密码
	captcha := c.PostForm("captcha")   //获取验证码
	qq := c.PostForm("qq")             //获取qq
	icode := c.PostForm("icode")       //获取邀请码
	err := testpackage.InitDB()                    // 调用输出化数据库的函数
	if username == "" || password == "" || qq == "" || icode == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "请填写各个输入框内容，谢谢~",
		})
		return
	}
	if testpackage.SQLInject(username) == true || testpackage.SQLInject(password) == true || testpackage.SQLInject(qq) == true || testpackage.SQLInject(icode) == true {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "别来注入了亲爱的，好好注册账号吧！",
		})
		return
	}
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "发生甚么事了？数据库连接失败！",
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
		if testpackage.Authcode(icode) != true { //验证邀请码是否存在并且未使用
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  500,
				"message": "邀请码不可用！",
			})
			return
		} else {
			if testpackage.Authuser(username) == true { //验证用户是否注册
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  500,
					"message": "该用户已经注册！",
				})
				return
			} else {

				if testpackage.Useradd(qq, username, password) == true { //如果注册成功，返回true
					testpackage.Usecode(icode, username, ipadd) //使用了邀请码
					c.JSON(http.StatusOK, gin.H{
						"status":  200,
						"url":     "/",
						"message": "注册成功！且用且珍惜,切勿乱传,争取多免杀一段时间，谢谢。",
					})
					return
				} else {
					c.JSON(http.StatusInternalServerError, gin.H{
						"status":  500,
						"message": "不知道出了什么问题，反正注册失败了",
					})
					return
				}
			}
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "验证码错误!",
		})
		return
	}
}
