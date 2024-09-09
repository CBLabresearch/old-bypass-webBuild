package admin
//删除邀请码和用户函数
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"web/config"
	"web/testpackage"
)

type T struct {
	Id int `json:"id"`
}

func Delidkey(c *gin.Context) { //删除邀请码函数
	t := &T{}
	c.ShouldBindJSON(t)
	fmt.Println()
	id := t.Id
	username := c.MustGet("username").(string)
	if username != config.AdminUserName { //验证当前用户是否为管理员
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "对不起,当前用户没有权限！",
		})
		return
	}
	err1 := testpackage.InitDB() // 调用输出化数据库的函数
	if err1 != nil {
		fmt.Printf("init db failed,err:%v\n", err1)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "发生甚么事了？数据库连接失败！",
		})
		return
	}
	if id ==0{
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "发生甚么事了？",
		})
		return
	}
	if testpackage.SQLInject(string(id)) == true {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "别来注入了亲爱的。",
		})
		return
	}
	if testpackage.Delkey(id) == true {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "删除成功！",
			"durl":"/api/admin/getinfo?istype=all",
		})
		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "删除失败！",
		})
		return
	}
}
func Deliduser(c *gin.Context) { //删除用户函数
	t := &T{}
	c.ShouldBindJSON(t)
	fmt.Println()
	id := t.Id
	username := c.MustGet("username").(string)
	if username != config.AdminUserName { //验证当前用户是否为管理员
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "对不起,当前用户没有权限！",
		})
		return
	}
	err1 := testpackage.InitDB() // 调用输出化数据库的函数
	if err1 != nil {
		fmt.Printf("init db failed,err:%v\n", err1)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "发生甚么事了？数据库连接失败！",
		})
		return
	}
	if id ==0{
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "发生甚么事了？",
		})
		return
	}
	if testpackage.SQLInject(string(id)) == true {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "别来注入了亲爱的。",
		})
		return
	}
	if testpackage.Deluser(id) == true {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "删除成功！",
			"durl":"/api/admin/getinfo?istype=user",
		})
		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "删除失败！",
		})
		return
	}
}
func Delidfile(c *gin.Context){		//删除文件记录
	t := &T{}
	c.ShouldBindJSON(t)
	fmt.Println()
	id := t.Id
	username := c.MustGet("username").(string)
	if username != config.AdminUserName { //验证当前用户是否为管理员
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "对不起,当前用户没有权限！",
		})
		return
	}
	err1 := testpackage.InitDB() // 调用输出化数据库的函数
	if err1 != nil {
		fmt.Printf("init db failed,err:%v\n", err1)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "发生甚么事了？数据库连接失败！",
		})
		return
	}
	if id ==0{
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "发生甚么事了？",
		})
		return
	}
	if testpackage.SQLInject(string(id)) == true {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "别来注入了亲爱的。",
		})
		return
	}
	if testpackage.Delfile(id) == true {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "删除成功！",
			"durl":"/api/admin/getinfo?istype=file",
		})
		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "删除失败！",
		})
		return
	}

}