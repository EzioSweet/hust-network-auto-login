# HUST校园网自动重连

本程序旨在为校内一些使用校园网的嵌入式设备或者服务器提供网络重连。

HUST的校园网连接机制是先从123.123.123.123获取queryString和epotalIP，然后访问http://eportalIp/eportal/InterFace.do?method=login通过application/x-www-form-urlencoded提交userId，password和queryString达成登录效果。

## Config说明

+ username:用户名
+ passowrd:密码
+ echo:每隔多少时间进行一次重连(单位：秒)

## 使用方式
从release里下载然后直接使用就行了(Linux下可能需要chmod 777)

请保证可执行文件和config.ini在同一目录