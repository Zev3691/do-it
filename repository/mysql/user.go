package mysql

import (
	"context"
	"re_new/util/log"
)

type User struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	EmployeeNumber string `json:"employee_number"`
	Persona        string `json:"persona"`
	Enable         int    `json:"enable"`
	Password       string `json:"password"`
	Phone          string `json:"phone"`
	Token          string `json:"token"`
	Salt           string `json:"salt"`
}

func (u *User) Table() string {
	return "users"
}

func (u *User) Create(ctx context.Context) error {
	db := NewMysqlDB(ctx)
	return db.Create(u).Error
}

func (u *User) LoginMath(ctx context.Context) error {
	request_id := ctx.Value("request_id")
	log.Debugf("request_idrequest_idrequest_idrequest_id %v", request_id)
	db := NewMysqlDB(ctx)
	db = db.Model(u).Where("name = ? AND password = ?", u.Name, u.Password)
	return db.First(u).Error
}

func (u *User) Update(ctx context.Context) error {
	db := NewMysqlDB(ctx)
	return db.Model(u).Updates(u).Error
}

func (u *User) Save(ctx context.Context) error {
	db := NewMysqlDB(ctx)
	return db.Model(u).Save(u).Error
}
