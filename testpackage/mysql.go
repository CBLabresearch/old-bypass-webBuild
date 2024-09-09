package testpackage

//mysql操作模块
import (
	"crypto/md5"
	"crypto/rand"
	"database/sql"
	"fmt"
	"log"
	"time"
	"web/config"
)

type userinfo struct {
	id       int
	username string
	password string
}
type codeinfo struct {
	code string
	id   int
	is   int
}

var DB *sql.DB

// 定义一个初始化数据库的函数
func InitDB() (err error) {
	// DSN:Data Source Name
	dsn := config.MysqlUsername + ":" + config.MysqlPassword + "@tcp(" + config.MysqlHost + ":" + config.MysqlPort + ")/" + config.MysqlDbname
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量DB
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = DB.Ping()
	if err != nil {
		return err
	}
	return nil
}
func Authcode(icode string) bool { //验证数据库里邀请码是否存在，如果存在并且未使用返回true
	sqlStr := "select id,code from codelist where isuse=0 and code=?"
	var c codeinfo
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := DB.QueryRow(sqlStr, icode).Scan(&c.id, &c.code)
	if err != nil {
		//fmt.Printf("scan failed, err:%v\n", err)
		return false
	}
	fmt.Printf("code:%s\n", c.code)
	return true
}
func Usecode(icode string, username string, ipadd string) bool { //使用了邀请码
	timeStr := time.Now().Format("2006-01-02 15:04:05") //获取注册时间
	sqlStr := "update codelist set isuse=1,usetime=?,username=?,ip=? where code = ?"
	_, err := DB.Exec(sqlStr, timeStr, username, ipadd, icode)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return false
	}
	fmt.Println("邀请码已经使用。")
	return true
}
func Authuser(uname string) bool { //验证数据库里用户是否存在，如果存在返回true
	sqlStr := "select id,username from users where username=?"
	var u userinfo
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := DB.QueryRow(sqlStr, uname).Scan(&u.id, &u.username)
	if err != nil {
		//fmt.Printf("scan failed, err:%v\n", err)
		return false
	}
	fmt.Printf("username:%s\n", u.username)
	return true
}
func Useradd(qq string, uname string, passwd string) bool { //用户注册模块
	timeStr := time.Now().Format("2006-01-02 15:04:05") //获取注册时间
	md5passwd := md5.Sum([]byte(passwd))                //将密码进行md5加密
	encodepasswd := fmt.Sprintf("%x", md5passwd)        //将[]byte转成16进制
	sqlStr := "insert into users(qq,username, password,vip,regtime) values (?,?,?,0,?)"
	_, err := DB.Exec(sqlStr, qq, uname, encodepasswd, timeStr)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return false
	}
	return true
}
func UserLogin(username string, passwd string) bool { //用户登录模块
	var u userinfo
	md5passwd := md5.Sum([]byte(passwd))         //将密码进行md5加密
	encodepasswd := fmt.Sprintf("%x", md5passwd) //将[]byte转成16进制
	sqlstr := "select id,username from users where username=? and password=?"
	err := DB.QueryRow(sqlstr, username, encodepasswd).Scan(&u.username, &encodepasswd)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return false
	}
	return true
}
func FileAdd(username string, filename string, ipadd string, filetype string) bool { //用户生成木马记录模块
	timeStr := time.Now().Format("2006-01-02 15:04:05") //获取生成时间
	sqlStr := "insert into file(user, filename,filetype,time,ip) values (?,?,?,?,?)"
	_, err := DB.Exec(sqlStr, username, filename, filetype, timeStr, ipadd)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return false
	}
	return true
}
func VipFileAdd(username string, filename string, ipadd string, filetype string) bool { //vip用户生成木马记录模块
	//timeStr := time.Now().Format("2006-01-02 15:04:05") //获取生成时间
	sqlStr := "insert into file(user, filename,filetype,time,ip) values (?,?,?,?,?)"
	_, err := DB.Exec(sqlStr, username, filename, filetype, "1999-09-09 14:20:58", ipadd)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return false
	}
	return true
}
func LoginLog(username string, ipadd string) bool { //用户登录日志
	timeStr := time.Now().Format("2006-01-02 15:04:05") //获取生成时间
	sqlStr := "insert into loginlog(username, logintime,ip) values (?,?,?)"
	_, err := DB.Exec(sqlStr, username, timeStr, ipadd)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return false
	}
	return true
}
func Lastlogin(username string) bool { //用户登录日志
	timeStr := time.Now().Format("2006-01-02 15:04:05") //获取生成时间
	sqlStr := "UPDATE `users` SET `lastlogin`=? WHERE `username`=?;"
	_, err := DB.Exec(sqlStr, timeStr, username)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return false
	}
	return true
}
func Authfilenub(username string) string { //验证用户生成次数
	timeStr := time.Now().Format("2006-01-02") //获取生成时间
	sqlStr := "SELECT COUNT(*) FROM `file` WHERE `time` LIKE ? AND `user` =?"
	err := DB.QueryRow(sqlStr, timeStr+"%", username).Scan(&username)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return "err"
	}
	fmt.Println(string(username))
	return string(username)
}
func GetuserQQ(username string) string { //获取用户qq号
	//timeStr := time.Now().Format("2006-01-02") //获取生成时间
	var qq string
	sqlStr := "SELECT qq FROM `users` WHERE `username` = ? "
	err := DB.QueryRow(sqlStr, username).Scan(&qq)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return "err"
	}
	fmt.Println(qq)
	return qq
}
func Authvip(username string) bool { //验证当前用户是否vip，如果是：返回true
	sqlStr := "select id,username from users where vip=1 and username=?"
	var c userinfo
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := DB.QueryRow(sqlStr, username).Scan(&c.id, &c.username)
	if err != nil {
		//fmt.Printf("scan failed, err:%v\n", err)
		return false
	}
	fmt.Printf("code:%s\n", c.username)
	return true
}
func Getkey() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	fmt.Println(uuid)
	return uuid
}
func Addkey(ipadd string) bool { //生成邀请码
	timeStr := time.Now().Format("2006-01-02 15:04:05") //获取生成时间
	sqlStr := "INSERT INTO `codelist` (`code`,`isuse`,`creattime`,`creatip`) VALUES (?,0,?,?)"
	_, err := DB.Exec(sqlStr, Getkey(), timeStr, ipadd)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return false
	}
	return true
}
func Delkey(id int) bool { //删除邀请码
	sqlStr := "DELETE FROM codelist WHERE id=?"
	_, err := DB.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return false
	}
	return true
}
func Deluser(id int) bool { //删除用户
	sqlStr := "DELETE FROM users WHERE id=?"
	_, err := DB.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return false
	}
	return true
}
func Delfile(id int) bool { //删除文件记录
	sqlStr := "DELETE FROM file WHERE id=?"
	_, err := DB.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return false
	}
	return true
}
func Addvip(id int) bool { //升级为内测用户
	sqlStr := "UPDATE `users` SET `vip`=1 WHERE `id`=?"
	_, err := DB.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return false
	}
	return true
}
func Delvip(id int) bool { //降级为内测用户
	sqlStr := "UPDATE `users` SET `vip`=0 WHERE `id`=?"
	_, err := DB.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return false
	}
	return true
}
