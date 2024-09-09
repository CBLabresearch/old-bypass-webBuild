GoWebShellCode-Bypass
===
一个练手的基于gin框架搞的在线免杀平台，实现功能如下:

```azure
1.邀请码注册登录
2.邀请码生成
3.简单的后台管理
4.用户生成记录查询
5.普通用户每日可生成一次，内测用户可无限生成
    ....
```
需要golang环境和mysql   

config.go里配置连接地址和管理员用户名 以及设置cookie的站点地址

# 截图
后台
![](img/后台.png)
用户前台
![](img/用户前台.png)
# 更新日志
===
````
学习中
===
2020/11/11 增加延时上线执行功能，修复一堆bug
2020/12/14 增加hex加密，恢复免杀，去除一些特征。
2020/12/24 后端改用go重写，恢复免杀，去除延时和反虚拟机功能
2020/12/26 增加注册功能。
2020/12/27 增加延时上线设置功能。
2021/01/25 增加后台管理。
2021/01/26 增加捆绑马，自启马。
2021/03/10 增加图片隐写功能
2020/03/18 增加溯源模式
===
````
# Link

[Myblog](https://www.nctry.com)