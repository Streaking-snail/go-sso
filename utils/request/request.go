package request

import (
	"encoding/json"
	//"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net"
)



func GetClientIp(r *http.Request) string {
	remoteIp := r.RemoteAddr

	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		remoteIp = ip
	} else if ip = r.Header.Get("X-forwarded-For"); ip != "" {
		remoteIp = ip
	} else {
		remoteIp, _, _ = net.SplitHostPort(remoteIp)
	}
	if remoteIp == "::1" {
		remoteIp="127.0.0.1"
	}
	return remoteIp
}
func GetJson(r *http.Request) (map[string]interface{},error ){
	jsonstr, _ := ioutil.ReadAll(r.Body)
	var data map[string]interface{}
	err := json.Unmarshal(jsonstr, &data)
	return data,err
}

// func GetParam(r *http.Request, key string)(string,bool){
// 	val:=r.Header.Get(key)
// 	if val!=""{
// 		return val,true
// 	}
// 	val, err :=r.Cookie(key)
// 	if err!=nil{
// 		return "",false
// 	}
// 	return val,true
// }

