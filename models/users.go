package models

import (
	"fmt"
	"time"
	"reflect"
)

//定义User模型，绑定users表，ORM库操作数据库，需要定义一个struct类型和MYSQL表进行绑定或者叫映射，struct字段和MYSQL表字段一一对应
//在这里Users类型可以代表mysql users表
type Users struct {
	//通过在字段后面的标签说明，定义golang字段和表字段的关系
    //例如 `gorm:"column:name"` 标签说明含义是: Mysql表的列名（字段名)为name
    //这里golang定义的Name变量和MYSQL表字段name一样，他们的名字可以不一样。
	Ctime  int       `json:"ctime" gorm:"column:ctime"`
	Email  string    `json:"email" gorm:"column:email"`
	Ext    string    `json:"ext" gorm:"column:ext"`
	Id     int64     `json:"id"`
	Mtime  time.Time `json:"mtime" gorm:"column:mtime"`
	Name   string    `json:"name" gorm:"column:name"`
	Passwd string    `json:"passwd" gorm:"passwd"`
	Mobile  string   `gorm:"column:mobile;unique;not null" json:"mobile"`
	Salt   string    `json:"salt" gorm:"salt"`
	Status int       `json:"status" gorm:"status"`
}

//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (u Users) TableName() string {
    //绑定MYSQL表名为users
    return "users"
}

var UsersStatusOk = 1
//手机号是否已存在
func (u *Users) GetRow() bool{
	err := db.Where(u).First(u)
	if err.Error == nil && !u.IsEmpty() {
		return true
	}
	return false
}

//判断结构体是否为空
func (u Users) IsEmpty() bool {
    return reflect.DeepEqual(u, Users{})
}

//新增用户
func (u *Users) Add(trace *Trace, device *Device) (int64, error) {
	tx := db.Begin()
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return 0, err
	}

	result := tx.Create(u)
	if err := result.Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	trace.Uid = u.Id
	result = tx.Create(trace)
	if err := result.Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	device.Uid = u.Id
	result = tx.Create(device)
	if err := result.Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	return u.Id, tx.Commit().Error
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
