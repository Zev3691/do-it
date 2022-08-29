package user

import (
	"context"
	"re_new/repository/mysql"
)

type CreateRequest struct {
	Name           string `json:"name"`
	Pwd            string `json:"pwd"`
	EmployeeNumber string `json:"employee_number"`
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
	return model.Create(ctx)
}
