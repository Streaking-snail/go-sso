package response

import (
	"net/http"
	"encoding/json"
)

type JsonResult  struct{
    Code int `json:"code"`
    Msg  string `json:"msg"`
}

type JsonData  struct{
    Code int `json:"code"`
    Msg  interface{} `json:"msg"`
}

func ShowValidatorError(w http.ResponseWriter, data interface{}){
	msg, _ := json.Marshal(JsonData{Code: 400, Msg: data})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(msg)
}

func ShowError(w http.ResponseWriter, data string){
	msg, _ := json.Marshal(JsonResult{Code: 400, Msg: data})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(msg)
}

func ShowSuccess(w http.ResponseWriter, data string){
	msg, _ := json.Marshal(JsonResult{Code: 200, Msg: data})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(msg)
}

func ShowData(w http.ResponseWriter, data interface{}){
	msg, _ := json.Marshal(JsonData{Code: 200, Msg: data})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(msg)
}
