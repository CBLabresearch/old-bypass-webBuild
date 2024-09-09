package config

//cookie设置，设置cookie的站点
const (
	HostName = "127.0.0.1"
)

//server设置
const (
	Host          = "0.0.0.0"
	Port          = "80"
	MysqlUsername = "bypass" //name    数据库配置
	MysqlPassword = "bypass" //password
	MysqlHost     = "127.0.0.1" //数据库地址
	MysqlPort     = "3306" //数据库端口
	MysqlDbname   = "bypass"  //库名
)

//管理员设置

const (
	AdminUserName = "try"
)
//Jwt的key

const (
	Jwt_key = "这是一个未知的私钥"
)
