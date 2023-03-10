package app

import (
	"github.com/dgrijalva/jwt-go"
	//"github.com/gin-gonic/gin"
	"go-gorm/conf"
	"go-gorm/models"
	"strconv"
	"time"
	"net/http"
)

func DoLogin(w http.ResponseWriter, r *http.Request, user models.Users)  error{
	secure:=IsHttps(w, r)
	if conf.Cfg.OpenJwt { //返回jwt
		customClaims :=&CustomClaims{
			UserId:         user.Id,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Duration(MAXAGE)*time.Second).Unix(), // 过期时间，必须设置
			},
		}
		accessToken, err :=customClaims.MakeToken()
		if err != nil {
			return err
		}
		refreshClaims :=&CustomClaims{
			UserId:         user.Id,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Duration(MAXAGE+1800)*time.Second).Unix(), // 过期时间，必须设置
			},
		}
		refreshToken, err :=refreshClaims.MakeToken()
		if err != nil {
			return err
		}
		r.Header.Add(ACCESS_TOKEN,accessToken)
		r.Header.Add(REFRESH_TOKEN,refreshToken)
		cookie := http.Cookie{Name: ACCESS_TOKEN, Value: accessToken, MaxAge: MAXAGE, Path: "/", Secure:secure, HttpOnly:true}
		http.SetCookie(w, &cookie)
		//w.SetCookie(ACCESS_TOKEN,accessToken,MAXAGE,"/", "",	 secure,true)
		cookie = http.Cookie{Name: REFRESH_TOKEN, Value: refreshToken, MaxAge: MAXAGE, Path: "/", Secure:secure, HttpOnly:true}
		http.SetCookie(w, &cookie)
		//w.SetCookie(REFRESH_TOKEN,refreshToken,MAXAGE,"/", "",	 secure,true)
	}
	//claims,err:=ParseToken(accessToken)
	//if err!=nil {
	//	return err
	//}
	id := strconv.Itoa(int(user.Id))
	cookie := http.Cookie{Name: COOKIE_TOKEN, Value: id, MaxAge: MAXAGE, Path: "/", Secure:secure, HttpOnly:true}
	//w.SetCookie(COOKIE_TOKEN,id,MAXAGE,"/", "",	 secure,true)
	http.SetCookie(w, &cookie)

	return nil
}
//判断是否https
func IsHttps(w http.ResponseWriter, r *http.Request ) bool {
	//h := new(http.Request)
	if r.Header.Get(HEADER_FORWARDED_PROTO) =="https" || r.TLS!=nil{
		return true
	}
	return false
}