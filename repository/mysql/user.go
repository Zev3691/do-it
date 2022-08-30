package mysql

import (
	"context"
	"errors"
	"re_new/util/log"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type User struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	EmployeeNumber string `json:"employee_number"`
	Persona        string `json:"persona"`
	Enable         int    `json:"enable"`
	Password       string `json:"password"`
	Phone          string `json:"phone"`
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
	db := NewMysqlDB(ctx)

	u1 := &User{}
	if err := db.Model(u1).Where("name = ?", u.Name).First(u1).Error; err != nil {
		log.Error(ctx, err.Error(), zap.String("login", "根据name查找用户失败"))
		return errors.New("找不到用户")
	}

	if err := db.Model(u).Where("name = ? AND password = ?", u.Name, u.Password).First(u).Error; err != nil {
		log.Error(ctx, err.Error(), zap.String("login", "匹配用户失败"))
		return errors.New("匹配用户失败")
	}
	return nil
}

func (u *User) Update(ctx context.Context) error {
	db := NewMysqlDB(ctx)
	return db.Model(u).Updates(u).Error
}

func (u *User) Save(ctx context.Context) error {
	db := NewMysqlDB(ctx)
	return db.Model(u).Save(u).Error
}

func (u *User) List(ctx context.Context, limit, offset int, scopes ...func(*gorm.DB) *gorm.DB) ([]User, int64, error) {
	db := NewMysqlDB(ctx)
	db = db.Model(u).Select("id,name,employee_number,persona,enable,phone")

	for _, v := range scopes {
		db = v(db)
	}
	var count int64
	db = db.Count(&count)
	var ret []User
	if err := db.Limit(limit).Offset(offset).Find(&ret).Error; err != nil {
		log.Error(ctx, err.Error())
		return nil, 0, err
	}

	return ret, count, nil
}

func FindById(ctx context.Context, id int) (*User, error) {
	db := NewMysqlDB(ctx)
	u := &User{ID: id}
	if err := db.Model(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}
