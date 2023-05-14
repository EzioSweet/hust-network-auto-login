package main

import (
	"fmt"
	"github.com/go-zoox/fetch"
	"github.com/gookit/ini/v2"
	"github.com/gookit/slog"
	"net"
	"regexp"
	"time"
)

func NetWorkStatus() bool {
	timeout := time.Second
	t1 := time.Now()
	_, err := net.DialTimeout("tcp", "www.baidu.com:443", timeout)
	slog.Info("waist time :", time.Now().Sub(t1))
	if err != nil {
		slog.Error("未连接网络，尝试重连")
		return false
	}
	slog.Info("网络已连接")
	return true
}

func main() {
	err := ini.LoadFiles("./config.ini")
	if err != nil {
		slog.Panic("未找到配置文件，请配置config.ini")
		panic(err)
	}
	username := ini.String("username")
	password := ini.String("password")
	echo := ini.Int("echo")
	for {
		time.Sleep(time.Duration(echo) * time.Second)
		if !NetWorkStatus() {
			response, err := fetch.Get("http://123.123.123.123")
			regQuerySring := regexp.MustCompile("index.jsp\\?(.*?)'</script>")
			regIp := regexp.MustCompile("http://(.*?)/eportal")
			eportalIp := regIp.FindStringSubmatch(string(response.Body))[1]
			queryString := regQuerySring.FindStringSubmatch(string(response.Body))[1]
			fetch.Post(fmt.Sprintf("http://%s/eportal/InterFace.do?method=login", eportalIp), &fetch.Config{
				Query: map[string]string{
					"userId":          username,
					"password":        password,
					"service":         "",
					"queryString":     queryString,
					"passwordEncrypt": "false",
				},
				Headers: map[string]string{
					"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
					"Accept":       "*/*",
					"User-Agent":   "hust-connector",
				},
			})
			if err != nil {
				slog.Error("登录失败")
			}
			slog.Info("登录成功")
		}
	}

}
