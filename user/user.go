package user

//用户中心函数
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/testpackage"
)

func Usertype(c *gin.Context) {
	istype := c.DefaultQuery("type", "")
	username := c.MustGet("username").(string)
	number := testpackage.Authfilenub(username)
	if istype == "bypassav" {
		c.HTML(http.StatusOK, "bypassav.html", gin.H{
			"user":   username,
			"number": number,
		})
	} else if istype == "binding" {
		c.HTML(http.StatusOK, "binding.html", gin.H{
			"user":   username,
			"number": number,
		})
	} else if istype == "addself" {
		c.HTML(http.StatusOK, "addself.html", gin.H{
			"user":   username,
			"number": number,
		})
	} else if istype == "encodeimg" {
		c.HTML(http.StatusOK, "encodeimg.html", gin.H{
			"user":   username,
			"number": number,
		})
	} else if istype=="inject" {
		c.HTML(http.StatusOK, "Inject.html", gin.H{
			"user":   username,
			"number": number,
		})
	}else {
		c.HTML(http.StatusOK, "user.html", gin.H{
			"user":   username,
			"number": number,
		})
	}

}
