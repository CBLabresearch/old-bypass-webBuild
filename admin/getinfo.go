package admin

//查询相关信息
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"web/config"
	"web/testpackage"
)

type Keyinfo struct { //邀请码信息
	Id        int    `json:"Id"`
	Code      string `json:"Code"`
	Isuse     int    `json:"Isuse"`
	Username  string `json:"Username`
	Usetime   string `json:"Usetime`
	Ip        string `json:"Ip"`
	Creattime string `json:"Creattime`
	Creatip   string `json:"Creatip"`
}
type Usersinfo struct { //用户信息
	Id        int    `json:"Id"`
	QQ        string `json:"QQ"`
	Username  string `json:"Username`
	Password  string `json:"Password`
	Vip       int    `json:"Vip"`
	Regtime   string `json:"Regtime`
	Lastlogin string `json:"Lastlogin"`
}
type Fileinfo struct { //邀请码信息
	Id       int    `json:"Id"`
	User     string `json:"User"`
	Filename string `json:"Filename`
	Time     string `json:"Time`
	Filetype string `json:"Filetype"`
	Ip       string `json:"Ip`
}
type Userlog struct {		//用户登录记录
	Id        int    `json:"Id"`
	Username  string `json:"Username"`
	Ip        string `json:"Ip"`
	Logintime string `json:"Logntime"`
}

func Createkey(c *gin.Context) { //生成邀请码
	ipadd := c.ClientIP() //获取ip地址
	username := c.MustGet("username").(string)
	if username != config.AdminUserName { //验证当前用户是否为管理员
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "对不起,当前用户没有权限！",
		})
		return
	}
	if testpackage.Addkey(ipadd) != true { //邀请码是否添加成功
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "对不起,发生了一些错误",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "生成成功！",
			"durl":    "/api/admin/getinfo?istype=allkey",
		})
		return
	}
}
func Select(c *gin.Context) {
	istype := c.DefaultQuery("istype", "")
	usname := c.DefaultQuery("name", "")
	fmt.Println(usname)
	username := c.MustGet("username").(string)
	if username != config.AdminUserName {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "对不起,当前用户没有权限哦！",
		})
		return
	} else {
		err1 := testpackage.InitDB() // 调用输出化数据库的函数
		if err1 != nil {
			fmt.Printf("init DB failed,err:%v\n", err1)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  500,
				"message": "发生甚么事了？数据库连接失败！",
			})
			return
		}
		if istype == "nokey" {
			c.HTML(http.StatusOK, "getkey.html", gin.H{"res": selectisuse()})
		} else if istype == "allkey" {
			c.HTML(http.StatusOK, "getkey.html", gin.H{"res": selectall()})
		} else if istype == "user" {
			c.HTML(http.StatusOK, "getusers.html", gin.H{"res": selectusers()})
		} else if istype == "file" {
			c.HTML(http.StatusOK, "getfile.html", gin.H{"res": selectfiles()}) //获取文件生成记录
		} else if istype == "userlog" {
			c.HTML(http.StatusOK, "userlog.html", gin.H{"res": selectuserlog(usname)}) //查询用户登录日志记录
		} else {
			c.HTML(http.StatusOK, "getkey.html", gin.H{})
		}

	}

}
func selectall() []Keyinfo {
	rows, errq := testpackage.DB.Query("select id,code,isuse,username,usetime,ip,creattime from codelist order by id desc")
	if errq != nil {
	}
	var keys []Keyinfo
	//遍历结果
	for rows.Next() {
		var u Keyinfo
		errn := rows.Scan(&u.Id, &u.Code, &u.Isuse, &u.Username, &u.Usetime, &u.Ip, &u.Creattime)
		if errn != nil {
			//fmt.Printf("%v", errn)
		}
		keys = append(keys, u)
	}
	return keys
}
func selectisuse() []Keyinfo { //未用的邀请码
	rows, errq := testpackage.DB.Query("select id,code,isuse,username,usetime,ip,creattime from codelist where isuse=0 order by id desc")
	if errq != nil {

	}
	var keys []Keyinfo

	//遍历结果
	for rows.Next() {
		var u Keyinfo
		errn := rows.Scan(&u.Id, &u.Code, &u.Isuse, &u.Username, &u.Usetime, &u.Ip, &u.Creattime)
		if errn != nil {
			//fmt.Printf("%v", errn)
		}
		keys = append(keys, u)
	}
	return keys
}
func selectusers() []Usersinfo { //查询用户
	rows, errq := testpackage.DB.Query("select id,qq,username,vip,lastlogin,regtime from users order by id desc")
	if errq != nil {

	}
	var users []Usersinfo

	//遍历结果
	for rows.Next() {
		var u Usersinfo
		errn := rows.Scan(&u.Id, &u.QQ, &u.Username, &u.Vip, &u.Lastlogin, &u.Regtime)
		if errn != nil {
			//fmt.Printf("%v", errn)
		}
		users = append(users, u)
	}
	return users
}
func selectfiles() []Fileinfo {
	rows, errq := testpackage.DB.Query("select id,user,filename,filetype,time,ip from file order by id desc")
	if errq != nil {

	}
	var files []Fileinfo

	for rows.Next() {
		var f Fileinfo
		errn := rows.Scan(&f.Id, &f.User, &f.Filename, &f.Filetype, &f.Time, &f.Ip)
		if errn != nil {

		}
		files = append(files, f)
	}
	return files
}
func selectuserlog(username string)[]Userlog {	//查询用户登录日志函数
	rows,err :=testpackage.DB.Query("select id,username,logintime,ip from loginlog where username=? order by id desc",username)
	if err !=nil{

	}
	var userlog []Userlog
	for rows.Next(){
		var u Userlog
		err :=rows.Scan(&u.Id,&u.Username,&u.Logintime,&u.Ip)
		if err !=nil{

		}
		userlog = append(userlog,u)

	}
	return userlog
}
