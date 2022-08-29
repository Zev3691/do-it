package user

import (
	"context"
	"re_new/repository/mysql"
	"re_new/util/errorx"
)

type CreateRequest struct {
	Name           string `json:"name" validate:"require"`
	Pwd            string `json:"pwd" validate:"require"`
	EmployeeNumber string `json:"employee_number" validate:"require"`
	Persona        string `json:"persona"`
	Enable         int    `json:"enable"`
	Phone          string `json:"phone"`
}

func Create(ctx context.Context, req *CreateRequest) error {
	model := &mysql.User{
		Name:           req.Name,
		EmployeeNumber: req.EmployeeNumber,
		Persona:        req.Persona,
		Enable:         req.Enable,
		Password:       req.Pwd,
		Phone:          req.Phone,
		Token:          "",
		Salt:           "",
	}
	if err := model.Create(ctx); err != nil {
		return errorx.New(errorx.ErrDBOptFailed, errorx.NewMsg(err.Error()))
	}
	return nil
}
