package user

import (
	"fmt"
	"net/http"
	"go-gorm/utils/response"
	"go-gorm/models"
	"go-gorm/modules/app"
	"encoding/json"
	"io/ioutil"
	"time"
	"go-gorm/utils/common"
	"go-gorm/utils/request"
)

type UserMobile struct {
	Mobile string `json:"mobile" gorm:"not null"`
	Passwd string `json:"passwd" gorm:"max=20;min=6"`
	Code   string `json:"code" gorm:"len=6"`
}

type UserMobilePasswd struct {
	Mobile string `json:"mobile" gorm:"not null"`
	Passwd string `json:"passwd" gorm:"not null;max=20;min=6"`
}

var MobileTrans = map[string]string{"mobile": "手机号"}

var UserMobileTrans = map[string]string{"Mobile": "手机号", "Passwd": "密码", "Code": "验证码"}

//注册
func Logon(w http.ResponseWriter, r *http.Request) {
	var userMobile UserMobile
	//接收post提交的json数据
	body, _ := ioutil.ReadAll(r.Body)
	//把json数据转化为结构体
	err := json.Unmarshal(body, &userMobile)
	if err != nil {
		response.ShowValidatorError(w, err)
		return
	}
	// if err != nil {
	// 	panic(err)
	// }
	//手机号是否存在
	model := models.Users{Mobile: userMobile.Mobile}
	if has := model.GetRow(); has {
		response.ShowError(w, "手机号已存在")
		return
	}
	
	model.Salt = common.GetRandomBoth(4)
	model.Passwd = common.Sha1En(userMobile.Passwd + model.Salt)
	model.Ctime = int(time.Now().Unix())
	model.Status = models.UsersStatusOk
	model.Mtime = time.Now()

	traceModel := models.Trace{Ctime: model.Ctime}
	traceModel.Ip = common.IpStringToInt(request.GetClientIp(r))
	traceModel.Type = models.TraceTypeReg

	deviceModel := models.Device{Ctime: model.Ctime, Ip: traceModel.Ip, Client: r.Header.Get("User-Agent")}
	_, err = model.Add(&traceModel, &deviceModel)
	if err != nil {
		fmt.Println(err)
		response.ShowError(w, "fail")
		return
	}
	response.ShowSuccess(w, "注册成功")

	return
}

//登录
func Login(w http.ResponseWriter, r *http.Request) {
	var userMobile UserMobilePasswd
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body,&userMobile)
	if err != nil {
		response.ShowValidatorError(w, err)
		return
	}
	
	model := models.Users{Mobile: userMobile.Mobile}
	//fmt.Printf("%v",model)
	if has := model.GetRow(); !has {
	    response.ShowError(w, "手机号不存在")
	 	return
	}
	if common.Sha1En(userMobile.Passwd+model.Salt) != model.Passwd {
	 	response.ShowError(w, "登录失败")
		return
	}
	err = app.DoLogin(w, r, model)
	if err != nil {
	 	response.ShowError(w, "登录失败")
	    return
    }
	response.ShowSuccess(w, "success")
	return
}

func checkErr(err error) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}
	}()
	if err != nil {
		panic(err)
	}
}
