package main

import (
	"net/http"
	"go-gorm/api/user"
	"go-gorm/conf"
	"fmt"
	"log"
)


func main() {
	//初始化数据
	Load()

	http.HandleFunc("/user/logon", user.Logon)  //注册
	http.HandleFunc("/user/login", user.Login)  //登录
	//ping
	http.HandleFunc("/pong", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello astaxie!")
	})

	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Load() {
	c := conf.Config{}
	c.Routes=[]string{"/ping","/renewal","/login","/login/mobile","/sendsms","/signup/mobile","/signup/mobile/exist"}
	c.OpenJwt=true//开启jwt
	conf.Set(c)
	//初始化数据验证
	//handle.InitValidate()
}
