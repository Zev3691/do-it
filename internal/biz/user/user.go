package user

import (
	"context"
	"re_new/repository/mysql"
	"re_new/util/cryptox"
	"re_new/util/errorx"
	"re_new/util/log"
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
		Salt:           cryptox.GenSalt(16),
	}
	if err := model.Create(ctx); err != nil {
		log.Error(ctx, err.Error())
		return errorx.New(errorx.ErrDBOptFailed, errorx.NewMsg(err.Error()))
	}
	return nil
}

type UpdateRequest struct {
	Id             int    `json:"id" validate:"require"`
	Name           string `json:"name"`
	Pwd            string `json:"pwd"`
	EmployeeNumber string `json:"employee_number"`
	Persona        string `json:"persona"`
	Enable         int    `json:"enable"`
	Phone          string `json:"phone"`
}

func Update(ctx context.Context, req *UpdateRequest) error {
	model := &mysql.User{
		ID:             req.Id,
		Name:           req.Name,
		EmployeeNumber: req.EmployeeNumber,
		Persona:        req.Persona,
		Enable:         req.Enable,
		Password:       req.Pwd,
		Phone:          req.Phone,
	}
	if err := model.Update(ctx); err != nil {
		log.Error(ctx, err.Error())
		return errorx.New(errorx.ErrDBOptFailed, errorx.NewMsg(err.Error()))
	}
	return nil
}

type ListRequest struct {
	PageSize int `json:"page_size"`
	Page     int `json:"page"`
}

type ListItem struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	EmployeeNumber string `json:"employee_number"`
	Persona        string `json:"persona"`
	Enable         int    `json:"enable"`
	Password       string `json:"password"`
	Phone          string `json:"phone"`
}

type ListResponse struct {
	Total int        `json:"total"`
	List  []ListItem `json:"list"`
}

func List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	ret := new(ListResponse)
	model := new(mysql.User)
	users, count, err := model.List(ctx, req.PageSize, (req.Page-1)*req.PageSize)
	if err != nil {
		return nil, err
	}

	ret.Total = int(count)
	for _, v := range users {
		ret.List = append(ret.List, ListItem{
			ID:             v.ID,
			Name:           v.Name,
			EmployeeNumber: v.EmployeeNumber,
			Persona:        v.Persona,
			Enable:         v.Enable,
			Password:       "********",
			Phone:          v.Phone,
		})
	}

	return ret, nil
}

type DeleteRequset struct {
	Id int `json:"id" validata:"require"`
}

func Delete(ctx context.Context, req *DeleteRequset) error {
	model := mysql.User{ID: req.Id}
	if err := model.Delete(ctx); err != nil {
		return err
	}

	return nil
}

type GetRequest struct {
	Id int `json:"id" validata:"require"`
}

func Get(ctx context.Context, req *GetRequest) (*ListItem, error) {
	user, err := mysql.FindUserById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	ret := &ListItem{
		ID:             user.ID,
		Name:           user.Name,
		EmployeeNumber: user.EmployeeNumber,
		Persona:        user.Persona,
		Enable:         user.Enable,
		Password:       "********",
		Phone:          user.Phone,
	}
	return ret, nil
}

func GetUserPwd(ctx context.Context, req *GetRequest) (string, error) {
	user, err := mysql.FindUserById(ctx, req.Id)
	if err != nil {
		return "", err
	}

	return cryptox.Base64(user.Password), nil
}
