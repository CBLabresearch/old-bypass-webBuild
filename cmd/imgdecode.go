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

func Encodeimg(c *gin.Context) {
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
	} else if path.Ext(file.Filename) != ".txt" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "第一个文件必须是txt文件",
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
	} else if path.Ext(file2.Filename) != ".jpg" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "第二个文件必须是jpg文件",
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
	log.Println(file.Filename)
	nfilename := Base64EncodeString(username + ":" + fmt.Sprintf("%d", time.Now().Unix()))
	filetxt := fmt.Sprintf("./upload/txt/%s", nfilename+".txt") //bin文件名称
	fileimg := fmt.Sprintf("./upload/img/%s", nfilename+".jpg") //exe文件名称
	c.SaveUploadedFile(file, filetxt)                                       //保存txt
	c.SaveUploadedFile(file2, fileimg)
	cmd := exec.Command("./payload/encodeimg/encodeimg", "-e", "-i", fileimg, "-mi", filetxt, "-o", "./getshell/img/"+nfilename+".jpg") //生成加密后的文件
	err1 := cmd.Run()
	if err1 != nil {
		// 命令执行失败
		fmt.Println(err1)
	}
	if Reshell("./getshell/img/"+nfilename+".jpg") == true { //验证文件是否成功
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"durl":    "/img/" + nfilename + ".jpg",
			"message": fmt.Sprintf("生成成功,点击跳转下载。"),
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": fmt.Sprintf("生成错误，请重新尝试。"),
		})
	}
}
